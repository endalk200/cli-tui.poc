package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v6"

	"github.com/go-git/go-git/v6/plumbing/object"
	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/spf13/cobra"
)

// commitCmd creates a commit for the currently staged changes. It supports:
//   - Providing a message via -m / --message
//   - If no message provided: auto-generating one via OpenAI (requires OPENAI_API_KEY)
//   - --author flag override ("Name <email>")
//   - --allow-empty to permit empty tree commits
//   - --dry-run to show what would be committed without writing
//   - --no-ai to force interactive/manual message requirement instead of AI fallback
//
// The auto-generated commit message aims to summarize staged paths with a concise
// imperative subject line and optional bullet points.
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit from staged changes (AI message fallback)",
	Long: `Create a commit from staged changes. If -m/--message is omitted and --no-ai
is not set, an AI generated message will be requested using OpenAI. This requires
OPENAI_API_KEY to be present in the environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			exitWithError(fmt.Errorf("cannot determine working directory: %w", err))
		}
		cli, err := NewGitCLI(cwd)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				exitWithError(fmt.Errorf("no git repository found at %s", cwd))
			}
			exitWithError(err)
		}

		message, _ := cmd.Flags().GetString("message")
		allowEmpty, _ := cmd.Flags().GetBool("allow-empty")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		noAI, _ := cmd.Flags().GetBool("no-ai")
		authorStr, _ := cmd.Flags().GetString("author")

		wt, err := cli.repo.Worktree()
		if err != nil {
			exitWithError(fmt.Errorf("failed to access worktree: %w", err))
		}
		status, err := wt.Status()
		if err != nil {
			exitWithError(fmt.Errorf("failed to compute status: %w", err))
		}

		// Collect staged entries (index changes) similar to status command logic.
		var stagedPaths []string
		for path, s := range status {
			if s.Staging != git.Unmodified { // something staged
				stagedPaths = append(stagedPaths, path)
			}
		}
		if len(stagedPaths) == 0 && !allowEmpty {
			exitWithError(errors.New("no staged changes to commit (use 'bgit add' or --allow-empty)"))
		}

		// Auto-generate commit message if needed.
		if message == "" && !noAI {
			apiKey := os.Getenv("OPENAI_API_KEY")
			if apiKey == "" {
				fmt.Println("OPENAI_API_KEY not set; falling back to basic generated message.")
				message = basicHeuristicMessage(stagedPaths)
			} else {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				defer cancel()
				generated, genErr := generateAICommitMessage(ctx, apiKey, stagedPaths)
				if genErr != nil {
					fmt.Printf("AI generation failed: %v\n", genErr)
					message = basicHeuristicMessage(stagedPaths)
				} else {
					message = generated
				}
			}
		}

		if message == "" { // message still empty and AI disabled
			exitWithError(errors.New("commit message required (provide -m or enable AI)"))
		}

		// Show plan in dry-run mode.
		if dryRun {
			fmt.Println("Dry run commit preview:")
			fmt.Println("Message:")
			fmt.Printf("  %s\n", firstLine(message))
			fmt.Println("Files:")
			for _, p := range stagedPaths {
				fmt.Printf("  • %s\n", p)
			}
			if len(stagedPaths) == 0 {
				fmt.Println("  (none; empty commit would be created)")
			}
			return
		}

		// Prepare author signature.
		var authorName, authorEmail string
		if authorStr != "" {
			// Expect format Name <email>
			parts := strings.Split(authorStr, "<")
			if len(parts) == 2 && strings.HasSuffix(parts[1], ">") {
				authorName = strings.TrimSpace(parts[0])
				authorEmail = strings.TrimSuffix(strings.TrimSpace(parts[1]), ">")
			}
		}
		if authorName == "" || authorEmail == "" {
			// Fallback to environment (like conventional git) or placeholders.
			authorName = getenvDefault("GIT_AUTHOR_NAME", "bgit user")
			authorEmail = getenvDefault("GIT_AUTHOR_EMAIL", "user@example.com")
		}

		// Construct commit object using worktree.Commit.
		commitHash2, err := wt.Commit(message, &git.CommitOptions{
			Author: &object.Signature{
				Name:  authorName,
				Email: authorEmail,
				When:  time.Now(),
			},
		})
		if err != nil {
			exitWithError(fmt.Errorf("failed to create commit: %w", err))
		}

		commitObj, err := cli.repo.CommitObject(commitHash2)
		if err != nil {
			exitWithError(fmt.Errorf("commit created but retrieval failed: %w", err))
		}

		fmt.Println("Commit created:")
		fmt.Printf("  Hash: %s\n", commitObj.Hash.String())
		fmt.Printf("  Author: %s <%s>\n", commitObj.Author.Name, commitObj.Author.Email)
		fmt.Printf("  Date: %s\n", commitObj.Author.When.Format(time.RFC3339))
		fmt.Printf("  Subject: %s\n", firstLine(commitObj.Message))
		if len(stagedPaths) > 0 {
			fmt.Println("\nFiles included:")
			for _, p := range stagedPaths {
				fmt.Printf("  • %s\n", p)
			}
		} else {
			fmt.Println("(Empty commit)")
		}
		fmt.Println("\nNext: push with 'git push' or continue working.")
	},
}

// basicHeuristicMessage builds a simple commit message when AI is unavailable.
func basicHeuristicMessage(paths []string) string {
	if len(paths) == 0 {
		return "chore: empty commit"
	}
	// Summarize primary file types.
	var exts = map[string]int{}
	for _, p := range paths {
		dot := strings.LastIndex(p, ".")
		if dot != -1 && dot != len(p)-1 {
			exts[p[dot+1:]]++
		}
	}
	var top []string
	for ext := range exts {
		if exts[ext] > 0 {
			plural := ext
			if exts[ext] > 1 {
				plural += " files"
			} else {
				plural += " file"
			}
			top = append(top, fmt.Sprintf("%d %s", exts[ext], plural))
		}
	}
	if len(top) > 3 {
		top = top[:3]
	}
	return fmt.Sprintf("update %d paths (%s)", len(paths), strings.Join(top, ", "))
}

// generateAICommitMessage calls OpenAI (v3 SDK) to synthesize a commit message.
func generateAICommitMessage(ctx context.Context, apiKey string, paths []string) (string, error) {
	client := openai.NewClient(option.WithAPIKey(apiKey))
	prompt := "Generate a concise conventional commit style message summarizing changes to these paths: " + strings.Join(paths, ", ") + ". " +
		"Use an imperative subject (max ~70 chars). If helpful, add 1-3 bullet lines after a blank line."

	resp, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4oMini,
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
		return "", errors.New("no AI response content")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

// getenvDefault returns environment variable value or fallback.
func getenvDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

// firstLine returns the first line of a multi-line string.
func firstLine(s string) string {
	if i := strings.IndexByte(s, '\n'); i != -1 {
		return s[:i]
	}
	return s
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringP("message", "m", "", "Commit message (if omitted uses AI or heuristic)")
	commitCmd.Flags().Bool("allow-empty", false, "Permit an empty commit")
	commitCmd.Flags().Bool("dry-run", false, "Preview commit without creating it")
	commitCmd.Flags().Bool("no-ai", false, "Require explicit -m message; no AI fallback")
	commitCmd.Flags().String("author", "", "Override author signature (\"Name <email>\")")
}
