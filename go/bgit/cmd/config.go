package cmd

import (
	"fmt"
	"os"

	"github.com/endalk200/bgit/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage bgit configuration",
	Long: `View and manage bgit configuration settings.

Configuration is stored in ~/.bgit.yaml by default.
You can specify a custom config file with --config flag.`,
}

var configViewCmd = &cobra.Command{
	Use:   "view",
	Short: "View current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		fmt.Println("Current Configuration:")
		fmt.Println("======================")
		fmt.Printf("AI Provider: %s\n", cfg.AIProvider.Name)
		fmt.Printf("Environment Variable: %s\n", cfg.AIProvider.EnvName)
	},
}

var configSetProviderCmd = &cobra.Command{
	Use:   "set-provider [provider-name]",
	Short: "Set the AI provider",
	Long: `Set the AI provider for commit message generation.

Available providers:
  - OpenAI (uses OPENAI_API_KEY)
  - OpenRouter (uses OPENROUTER_API_KEY)
  - Anthropic (uses ANTHROPIC_API_KEY)

Example:
  bgit config set-provider OpenRouter`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		providerName := args[0]

		// Find the provider in available providers
		var found bool
		var provider config.Provider
		for _, p := range config.AvailableProviders {
			if p.Name == providerName {
				found = true
				provider = p
				break
			}
		}

		if !found {
			fmt.Fprintf(os.Stderr, "error: unknown provider '%s'\n\n", providerName)
			fmt.Println("Available providers:")
			for _, p := range config.AvailableProviders {
				fmt.Printf("  - %s (env: %s)\n", p.Name, p.EnvName)
			}
			os.Exit(1)
		}

		err := config.SetProvider(provider.Name, provider.EnvName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: failed to update config: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓ Successfully set AI provider to: %s\n", provider.Name)
		fmt.Printf("  Environment variable: %s\n", provider.EnvName)
	},
}

var configListProvidersCmd = &cobra.Command{
	Use:   "list-providers",
	Short: "List available AI providers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available AI Providers:")
		fmt.Println("=======================")
		currentProvider := config.GetProvider()
		for _, p := range config.AvailableProviders {
			current := ""
			if p.Name == currentProvider.Name {
				current = " (current)"
			}
			fmt.Printf("  • %s%s\n", p.Name, current)
			fmt.Printf("    Environment Variable: %s\n", p.EnvName)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configViewCmd)
	configCmd.AddCommand(configSetProviderCmd)
	configCmd.AddCommand(configListProvidersCmd)
}
