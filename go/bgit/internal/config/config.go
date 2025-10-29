package config

type Provider struct {
	Name          string
	EnvName       string
	APIKeyEnvName string
}

var Providers = []Provider{
	{
		Name:    "OpenAI",
		EnvName: "OPENAI_API_KEY",
	},
	{
		Name:    "OpenRouter",
		EnvName: "OPENROUTER_API_KEY",
	},
}
