package cmd

import (
	"fmt"
	"os"

	"github.com/endalk200/bgit/internal/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// bgit is a learning / experimental Git wrapper built with go-git and Cobra.
// It aims to provide modern, readable output while exposing internal concepts
// clearly for educational purposes. The goal is to be production-grade in
// structure (error handling, separation of concerns, testability) while also
// being approachable for someone studying how Git works under the hood.
var rootCmd = &cobra.Command{
	Use:   "bgit",
	Short: "Modern, educational Git wrapper CLI",
	Long: `bgit is a modern, educational Git wrapper built on top of the pure Go
implementation of Git (go-git). It focuses on:

  • Clean, readable, colorized output
  • Explaining what is happening internally (comments / structure)
  • Production-grade patterns (clear errors, future extensibility)

Currently implemented subcommands:

  status  – Show repository status (staged / unstaged / untracked) with color
  add     – Stage file(s) or all changes with --all
  commit  – Create a commit; auto-generates a message when -m not supplied
  config  – View and manage configuration (AI provider settings)

Examples:
  bgit status
  bgit add --all
  bgit add path/to/file.go another/file.txt
  bgit config view
  bgit config set-provider OpenRouter

Configuration:
  bgit uses Viper for configuration management. Settings are stored in
  ~/.bgit.yaml by default. Use 'bgit config' to manage settings.

More commands will be added incrementally as learning exercises.`,
}

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialize config before running any commands
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bgit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.InitConfig(cfgFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
}
