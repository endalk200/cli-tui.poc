package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6"
	"github.com/spf13/cobra"
)

// addCmd stages changes similar to `git add`. It supports either specifying
// individual file paths or using the --all flag to stage everything reported
// by worktree status.
var addCmd = &cobra.Command{
	Use:   "add [files...]",
	Short: "Stage file contents into the index",
	Long: `Stage file contents into the index (staging area) similar to 'git add'.
You can provide explicit file paths or use --all to stage all tracked modifications
and new untracked files. Patterns (globs) within shell expansion also work.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("cannot determine working directory: %w", err))
		}

		cli, err := NewGitCLI(cwd)
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				panic(fmt.Errorf("no git repository found at %s", cwd))
			}
			panic(err)
		}

		worktree, err := cli.repo.Worktree()
		if err != nil {
			panic(fmt.Errorf("failed to access worktree: %w", err))
		}

		all, _ := cmd.Flags().GetBool("all")
		var targets []string

		if all {
			// Derive list from status; include modified, added(untracked), deleted (for remove) but ignore clean.
			status, err := worktree.Status()
			if err != nil {
				panic(fmt.Errorf("failed to compute status: %w", err))
			}
			for path, s := range status {
				if s.Worktree != git.Unmodified || s.Staging != git.Unmodified {
					// We attempt to stage all types; deletions are handled implicitly when path missing.
					targets = append(targets, path)
				}
			}
		} else {
			if len(args) == 0 {
				panic(errors.New("no paths provided (specify files or use --all)"))
			}
			for _, raw := range args {
				// Normalize path for consistency; relative -> absolute relative to repo root
				if !filepath.IsAbs(raw) {
					norm := filepath.Clean(raw)
					targets = append(targets, norm)
				} else {
					targets = append(targets, raw)
				}
			}
		}

		if len(targets) == 0 {
			fmt.Println("Nothing to stage")
			return
		}

		var staged []string
		var failed []string
		for _, path := range targets {
			// Attempt Add; for deletions, Add will error so we fallback to Remove if file absent
			if _, statErr := os.Stat(path); statErr != nil && os.IsNotExist(statErr) {
				// Try removing from index (deleted file)
				if _, remErr := worktree.Remove(path); remErr != nil {
					failed = append(failed, fmt.Sprintf("%s (remove failed: %v)", path, remErr))
					continue
				}
				staged = append(staged, path)
				continue
			}
			if _, addErr := worktree.Add(path); addErr != nil {
				failed = append(failed, fmt.Sprintf("%s (add failed: %v)", path, addErr))
				continue
			}
			staged = append(staged, path)
		}

		// Output summary
		fmt.Printf("Staged %d paths\n", len(staged))
		for _, p := range staged {
			fmt.Printf("  • %s\n", p)
		}
		if len(failed) > 0 {
			fmt.Printf("\nFailed (%d):\n", len(failed))
			for _, f := range failed {
				fmt.Printf("  • %s\n", f)
			}
		}

		// Provide next guidance similar to git status hints.
		if len(staged) > 0 {
			fmt.Println("\nNext: run 'bgit status' to review, then a future 'bgit commit' (not yet implemented).")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("all", "A", false, "Stage all tracked and untracked changes")
}
