# bgit Configuration Guide

## Overview

bgit uses [Viper](https://github.com/spf13/viper) for configuration management. This allows you to customize the AI provider and other settings used for commit message generation.

## Configuration File Location

By default, bgit looks for a configuration file at:

- `~/.bgit.yaml` (in your home directory)
- `./.bgit.yaml` (in the current directory)

You can also specify a custom config file using the `--config` flag:

```bash
bgit --config /path/to/config.yaml commit
```

## Configuration Format

The configuration file uses YAML format. Here's an example:

```yaml
ai_provider:
  name: OpenAI
  env_name: OPENAI_API_KEY
```

## Configuration Options

### AI Provider Settings

Configure which AI provider to use for commit message generation:

| Field                  | Description                          | Default Value    |
| ---------------------- | ------------------------------------ | ---------------- |
| `ai_provider.name`     | The name of the AI provider          | `OpenAI`         |
| `ai_provider.env_name` | Environment variable for the API key | `OPENAI_API_KEY` |

### Supported AI Providers

1. **OpenAI** (default)

   - Name: `OpenAI`
   - Environment Variable: `OPENAI_API_KEY`
   - Get your API key: https://platform.openai.com/api-keys

2. **OpenRouter**

   - Name: `OpenRouter`
   - Environment Variable: `OPENROUTER_API_KEY`
   - Get your API key: https://openrouter.ai/keys

3. **Anthropic**
   - Name: `Anthropic`
   - Environment Variable: `ANTHROPIC_API_KEY`
   - Get your API key: https://console.anthropic.com/

## Managing Configuration

### View Current Configuration

```bash
bgit config view
```

Output:

```
Current Configuration:
======================
AI Provider: OpenAI
Environment Variable: OPENAI_API_KEY
```

### List Available Providers

```bash
bgit config list-providers
```

Output:

```
Available AI Providers:
=======================
  • OpenAI (current)
    Environment Variable: OPENAI_API_KEY
  • OpenRouter
    Environment Variable: OPENROUTER_API_KEY
  • Anthropic
    Environment Variable: ANTHROPIC_API_KEY
```

### Change AI Provider

```bash
bgit config set-provider OpenRouter
```

This will update your `~/.bgit.yaml` file with the new provider settings.

## First-Time Setup

When you run bgit for the first time, it will automatically create a default configuration file at `~/.bgit.yaml` with these settings:

```yaml
ai_provider:
  name: OpenAI
  env_name: OPENAI_API_KEY
```

## Environment Variables

Make sure to set the appropriate environment variable for your chosen AI provider:

**For OpenAI:**

```bash
export OPENAI_API_KEY="sk-..."
```

**For OpenRouter:**

```bash
export OPENROUTER_API_KEY="sk-or-v1-..."
```

**For Anthropic:**

```bash
export ANTHROPIC_API_KEY="sk-ant-..."
```

You can add these to your shell profile (`~/.bashrc`, `~/.zshrc`, etc.) to make them permanent.

## Example Workflows

### Switching to OpenRouter

1. Set your OpenRouter API key:

   ```bash
   export OPENROUTER_API_KEY="sk-or-v1-..."
   ```

2. Update bgit configuration:

   ```bash
   bgit config set-provider OpenRouter
   ```

3. Verify the change:

   ```bash
   bgit config view
   ```

4. Use bgit as normal:
   ```bash
   bgit add --all
   bgit commit  # Will use OpenRouter for AI-generated message
   ```

### Using a Custom Config File

1. Create a custom config file:

   ```bash
   cat > ~/my-bgit-config.yaml << EOF
   ai_provider:
     name: Anthropic
     env_name: ANTHROPIC_API_KEY
   EOF
   ```

2. Use it with bgit:
   ```bash
   bgit --config ~/my-bgit-config.yaml commit
   ```

## Troubleshooting

### "API key not found" Error

If you see an error like:

```
error: OpenAI provider failed: API key not found: 401 OPENAI_API_KEY not set
```

Make sure:

1. The environment variable is set: `echo $OPENAI_API_KEY`
2. The variable name matches your config: `bgit config view`
3. You've exported the variable in your current shell session

### Config File Not Found

If bgit can't find your config file, it will create a default one automatically. You can also create it manually:

```bash
cp .bgit.yaml.example ~/.bgit.yaml
```

Then edit `~/.bgit.yaml` with your preferred settings.

## Advanced Usage

### Multiple Configurations

You can maintain different config files for different projects or environments:

```bash
# Work project using OpenAI
bgit --config ~/.bgit-work.yaml commit

# Personal project using OpenRouter
bgit --config ~/.bgit-personal.yaml commit
```

### Priority Order

Viper loads configuration in the following priority order (highest to lowest):

1. Explicit flags (e.g., `--config`)
2. Environment variables
3. Configuration file
4. Default values

This means you can override config file settings with environment variables if needed.
