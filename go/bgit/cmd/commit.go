package cmd

import (
	"fmt"
	"os"

	"github.com/endalk200/bgit/internal/config"
	commitgenService "github.com/endalk200/bgit/internal/services/commitgen"
	gitService "github.com/endalk200/bgit/internal/services/git"

	"github.com/spf13/cobra"
)

type ErrCanNotDetermineWorkingDirectory struct {
	Message string
}

func (e ErrCanNotDetermineWorkingDirectory) Error() string {
	return fmt.Sprintf("cannot determine working directory: %s", e.Message)
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Create a commit from staged changes (AI message fallback)",
	Long: `Create a commit from staged changes. If -m/--message is omitted and --no-ai
is not set, an AI generated message will be requested using OpenAI. This requires
OPENAI_API_KEY to be present in the environment.`,
	Run: func(cmd *cobra.Command, args []string) {
		message, _ := cmd.Flags().GetString("message")
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		noAI, _ := cmd.Flags().GetBool("no-ai")

		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot determine working directory: %v\n", err)
			os.Exit(1)
		}

		gitClient, err := gitService.NewGitClient(cwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		stagedFiles, err := gitClient.StagedFiles()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to get staged files: %v\n", err)
			os.Exit(1)
		}

		if len(stagedFiles) == 0 {
			fmt.Println("No staged files to commit. Use 'bgit add' to stage files first.")
			return
		}

		fmt.Printf("Found %d staged files:\n", len(stagedFiles))
		for _, file := range stagedFiles {
			fmt.Printf("  • %s\n", file)
		}
		fmt.Println()

		// If no message provided, generate one using AI
		if message == "" && !noAI {
			fmt.Println("Generating commit message using AI...")
			stagedDiff, err := gitClient.GetStagedFilesDiff(stagedFiles)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: failed to get staged diff: %v\n", err)
				os.Exit(1)
			}

			// Try OpenAI first, fallback to OpenRouter
			var generatedMessage string
			for _, provider := range config.Providers {
				generatedMessage, err = commitgenService.GenerateCommitMessage(stagedDiff, provider)
				if err == nil {
					break
				}
				fmt.Fprintf(os.Stderr, "warning: %s provider failed: %v\n", provider.Name, err)
			}

			if generatedMessage == "" {
				fmt.Fprintf(os.Stderr, "error: failed to generate commit message with all providers\n")
				fmt.Println("Hint: Provide a message with -m flag or set OPENAI_API_KEY environment variable")
				os.Exit(1)
			}

			message = generatedMessage
			fmt.Printf("Generated message: %s\n\n", message)
		} else if message == "" {
			fmt.Fprintf(os.Stderr, "error: commit message is required. Use -m flag or enable AI generation\n")
			os.Exit(1)
		}

		if dryRun {
			fmt.Println("=== DRY RUN ===")
			fmt.Printf("Would commit with message: %s\n", message)
			return
		}

		// Perform the actual commit
		err = gitClient.Commit(message)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to create commit: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringP("message", "m", "", "Commit message (if omitted uses AI or heuristic)")
	commitCmd.Flags().Bool("dry-run", false, "Preview commit without creating it")
	commitCmd.Flags().Bool("no-ai", false, "Disable AI commit message generation")
}
