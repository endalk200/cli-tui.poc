# Charm Package Examples

Different code examples for experimenting with different Charm packages.

- **Lipgloss**: Style definitions for nice terminal layouts
- **Log**: Structured, colorful logging
- **Huh**: Interactive terminal forms and prompts

## 📋 Prerequisites

- Go 1.24 or higher
- Terminal that supports ANSI colors

## 🚀 Quick Start

### Installation

1. Navigate to the charm directory:

```bash
cd go/charm
```

2. Install dependencies:

```bash
go mod download
```

### Running the Examples

Run the interactive menu:

```bash
go run main.go
```

You'll be presented with a menu to choose which package examples to run:

1. Lipgloss examples
2. Log examples
3. Huh examples
4. All examples
5. Exit

## 📚 What's Inside

### Lipgloss Examples (`examples/lipgloss.go`)

#### Easy Examples

- **SimpleLipglossExample**: Basic text styling with colors
- **BasicColorsExample**: Different color formats (ANSI, Hex, Named)
- **SimpleBordersExample**: Adding borders to text

#### Medium Examples

- **PaddingAndMarginsExample**: Spacing control inside and outside borders
- **AlignmentExample**: Text alignment (left, center, right)
- **JoinExample**: Combining styled elements horizontally and vertically
- **StyleInheritanceExample**: Reusing and extending base styles

#### Hard Examples

- **ComplexLayoutExample**: Multi-panel dashboard layout
- **ProgressBarExample**: Creating visual progress indicators
- **TableExample**: Formatted tables with styling
- **AdaptiveLayoutExample**: Responsive-like notification cards

### Log Examples (`examples/log.go`)

#### Easy Examples

- **SimpleLogExample**: Basic log levels (Debug, Info, Warn, Error)
- **LogWithFieldsExample**: Structured logging with key-value pairs
- **LogFormattingExample**: Customizing log output format

#### Medium Examples

- **LogLevelsExample**: Controlling log level filtering
- **SubLoggerExample**: Creating contextual child loggers
- **StructuredDataExample**: Logging complex data structures
- **LoggerOptionsExample**: Various logger configuration options

#### Hard Examples

- **ApplicationLoggerExample**: Production-ready logging setup
- **PerformanceLoggingExample**: Measuring and logging execution time
- **ErrorTrackingExample**: Comprehensive error tracking with context
- **AuditLogExample**: Creating audit trails for compliance
- **DistributedTracingExample**: Logging with trace IDs for distributed systems

### Huh Examples (`examples/huh.go`)

#### Easy Examples

- **SimpleInputExample**: Basic text input field
- **SimpleConfirmExample**: Yes/no confirmation dialog
- **SimpleSelectExample**: Single-choice selection menu

#### Medium Examples

- **MultiFieldFormExample**: Forms with multiple input fields
- **ValidationExample**: Input validation with custom rules
- **MultiSelectExample**: Selecting multiple options from a list
- **TextAreaExample**: Multi-line text input

#### Hard Examples

- **MultiPageFormExample**: Multi-step wizard-like forms
- **DynamicFormExample**: Conditional fields based on user input
- **ComplexWorkflowExample**: Complete application workflow with authentication
- **FormWithInlineHelpExample**: Forms with contextual help text
