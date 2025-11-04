package cmd

import (
	"errors"
	"fmt"

	"os"
	"strings"

	gitService "github.com/endalk200/bgit/internal/services/git"
	"github.com/go-git/go-git/v6"
	"github.com/spf13/cobra"
)

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

		gitClient, err := gitService.NewGitClient(cwd)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				fmt.Fprintf(os.Stderr, "error: no git repository found at %s\n", cwd)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		branch, _ := gitClient.CurrentBranch() // non-critical

		staged, err := gitClient.StagedFiles()
		if err != nil {
			staged = []string{}
		}

		modified, err := gitClient.ModifiedFiles()
		if err != nil {
			modified = []string{}
		}

		added, err := gitClient.AddedFiles()
		if err != nil {
			added = []string{}
		}

		deleted, err := gitClient.DeletedFiles()
		if err != nil {
			deleted = []string{}
		}

		renamed, err := gitClient.RenamedFiles()
		if err != nil {
			renamed = []string{}
		}

		untracked, err := gitClient.UntrackedFiles()
		if err != nil {
			untracked = []string{}
		}

		var out strings.Builder
		out.WriteString(fmt.Sprintf("On branch %s\n\n", branch))

		// Sections
		out.WriteString(formatSection("Staged (index)", staged))
		out.WriteString(formatSection("Added (staged new files)", added))
		out.WriteString(formatSection("Modified (worktree)", modified))
		out.WriteString(formatSection("Deleted", deleted))
		out.WriteString(formatSection("Renamed", renamed))
		out.WriteString(formatSection("Untracked", untracked))

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
