package cmd

import (
	"fmt"
	"os"

	internal "github.com/endalk200/bgit/internal/services/git"
	"github.com/spf13/cobra"
)

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

		client, err := internal.NewGitClient(cwd)
		if err != nil {
			panic(err.Error())
		}

		all, _ := cmd.Flags().GetBool("all")

		if all {
			_, err := client.AddAllFiles()
			if err != nil {
				panic(err.Error())
			}
		} else {
			var targets []string
			targets = append(targets, args...)
			stagedFiles, err := client.AddFiles(targets)
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("Staged %d files\n", len(stagedFiles))
			for _, file := range stagedFiles {
				fmt.Printf("  • %s\n", file)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("all", "A", false, "Stage all tracked and untracked changes")
}
