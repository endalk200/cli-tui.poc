package cmd

import (
	"fmt"
	"os"

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
		// message, _ := cmd.Flags().GetString("message")
		// dryRun, _ := cmd.Flags().GetBool("dry-run")

		cwd, err := os.Getwd()
		if err != nil {
			panic(ErrCanNotDetermineWorkingDirectory{Message: err.Error()})
		}

		gitClient, err := gitService.NewGitClient(cwd)
		if err != nil {
			panic(err)
		}

		stagedFiles, err := gitClient.StagedFiles()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Found %d staged files\n", len(stagedFiles))

		stagedDiff, err := gitClient.GetStagedFilesDiff(stagedFiles)
		if err != nil {
			panic(err)
		}

		fmt.Println(stagedDiff)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringP("message", "m", "", "Commit message (if omitted uses AI or heuristic)")
	commitCmd.Flags().Bool("dry-run", false, "Preview commit without creating it")
}
