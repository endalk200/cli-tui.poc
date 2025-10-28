package cmd

import (
	"errors"
	"fmt"

	"os"
	"sort"
	"strings"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/spf13/cobra"
)

type ErrNotAGitRepo struct {
	Message string
	Path    string
}

func (e *ErrNotAGitRepo) Error() string {
	message := fmt.Sprintf("not a git repository: %s", e.Path)
	if e.Message != "" {
		message = fmt.Sprintf("%s: %s", message, e.Message)
	}
	return message
}

func (e *ErrNotAGitRepo) Unwrap() error {
	return errors.New(e.Message)
}

type GitCLI struct {
	repo *git.Repository
}

func NewGitCLI(repoPath string) (*GitCLI, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			return nil, &ErrNotAGitRepo{Path: repoPath}
		}
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}
	return &GitCLI{repo: r}, nil
}

func (g *GitCLI) collectStatus() (staged, modified, added, deleted, renamed, untracked []string, err error) {
	wt, err := g.repo.Worktree()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("failed to get worktree: %w", err)
	}
	status, err := wt.Status()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("failed to compute status: %w", err)
	}

	for path, s := range status {
		// Staged changes are indicated by Staging property (Index status)
		if s.Staging != git.Unmodified {
			// Determine specific staged type
			switch s.Staging {
			case git.Modified:
				staged = append(staged, path)
			case git.Added:
				added = append(added, path)
			case git.Deleted:
				deleted = append(deleted, path)
			case git.Renamed:
				renamed = append(renamed, path)
			case git.Untracked: // rarely appears as staging state
				untracked = append(untracked, path)
			}
		}
		// Worktree modifications (not yet staged)
		if s.Worktree != git.Unmodified {
			switch s.Worktree {
			case git.Modified:
				modified = append(modified, path)
			case git.Added:
				untracked = append(untracked, path) // treat added in worktree as untracked
			case git.Deleted:
				deleted = append(deleted, path)
			case git.Renamed:
				renamed = append(renamed, path)
			case git.Untracked:
				untracked = append(untracked, path)
			}
		}
	}

	// Keep output stable
	for _, slice := range [][]string{staged, modified, added, deleted, renamed, untracked} {
		sort.Strings(slice)
	}
	return
}

func (g *GitCLI) currentBranch() (string, error) {
	headRef, err := g.repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to resolve HEAD: %w", err)
	}

	// If symbolic ref contains refs/heads/ prefix we strip it; otherwise show short SHA
	name := headRef.Name().String()
	if strings.HasPrefix(name, "refs/heads/") {
		return strings.TrimPrefix(name, "refs/heads/"), nil
	}
	return headRef.Hash().String()[:12], nil
}

// latestCommit retrieves the latest commit for HEAD for additional context.
func (g *GitCLI) latestCommit() (*object.Commit, error) {
	headRef, err := g.repo.Head()
	if err != nil {
		return nil, err
	}
	commit, err := g.repo.CommitObject(headRef.Hash())
	if err != nil {
		return nil, err
	}
	return commit, nil
}

// formatSection renders a titled list with bullet points.
func formatSection(title string, items []string) string {
	if len(items) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(title)
	b.WriteString(" (" + fmt.Sprintf("%d", len(items)) + ")\n")
	for _, it := range items {
		b.WriteString("  • ")
		b.WriteString(it)
		b.WriteString("\n")
	}
	return b.String()
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show repository status with modern formatting",
	Long: `Displays tracked, staged, modified, and untracked files with concise
categorization. Mirrors 'git status' conceptually but focuses on clarity.`,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot determine working directory: %v\n", err)
			os.Exit(1)
		}
		cli, err := NewGitCLI(cwd)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				fmt.Fprintf(os.Stderr, "error: no git repository found at %s\n", cwd)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		branch, _ := cli.currentBranch() // non-critical
		commit, _ := cli.latestCommit()  // non-critical

		staged, modified, added, deleted, renamed, untracked, err := cli.collectStatus()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		var out strings.Builder
		out.WriteString(fmt.Sprintf("On branch %s\n", branch))
		if commit != nil {
			out.WriteString(fmt.Sprintf("Latest commit: %s %s\n\n", commit.Hash.String()[:12], strings.Split(commit.Message, "\n")[0]))
		}

		// Sections
		out.WriteString(formatSection("Staged (index)", staged))
		out.WriteString(formatSection("Added (staged new files)", added))
		out.WriteString(formatSection("Modified (worktree)", modified))
		out.WriteString(formatSection("Deleted", deleted))
		out.WriteString(formatSection("Renamed", renamed))
		out.WriteString(formatSection("Untracked", untracked))

		if out.Len() == 0 { // improbable, but guard
			out.WriteString("Clean working tree\n")
		}

		// If there are no changes at all show a single line.
		if len(staged)+len(modified)+len(added)+len(deleted)+len(renamed)+len(untracked) == 0 {
			fmt.Println("Working tree clean")
			return
		}

		fmt.Print(out.String())
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output (future use)")
}
