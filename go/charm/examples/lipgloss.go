package examples

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// ==============================================================================
// EASY EXAMPLES - Basic styling concepts
// ==============================================================================

// SimpleLipglossExample demonstrates the most basic lipgloss styling
// Concept: Creating a style and applying foreground color
func SimpleLipglossExample() {
	fmt.Println("\n=== EASY: Simple Text Styling ===")

	// Create a style with a single property - foreground color
	// Colors can be specified using ANSI codes (0-255) or hex values
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("205")) // Pink color

	// Render applies the style to the text
	fmt.Println(style.Render("Hello, Lipgloss! This text is pink."))

	// You can also chain multiple properties
	boldStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("46")) // Green

	fmt.Println(boldStyle.Render("This is bold and green!"))
}

// BasicColorsExample shows how to work with different color formats
// Concept: Understanding color specifications in lipgloss
func BasicColorsExample() {
	fmt.Println("\n=== EASY: Color Formats ===")

	// Method 1: ANSI 256 colors (0-255)
	ansiStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")) // Bright red
	fmt.Println(ansiStyle.Render("ANSI Color (196)"))

	// Method 2: Hex colors
	hexStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FF00")) // Green
	fmt.Println(hexStyle.Render("Hex Color (#00FF00)"))

	// Method 3: Named colors (via ANSI)
	backgroundStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).  // Black text
		Background(lipgloss.Color("226")) // Yellow background
	fmt.Println(backgroundStyle.Render("Black text on yellow background"))
}

// SimpleBordersExample introduces basic border styling
// Concept: Adding borders to styled text
func SimpleBordersExample() {
	fmt.Println("\n=== EASY: Basic Borders ===")

	// NormalBorder is a simple single-line border
	normalBorder := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")) // Purple
	fmt.Println(normalBorder.Render("Normal Border"))

	// RoundedBorder has rounded corners
	roundedBorder := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("86")) // Cyan
	fmt.Println(roundedBorder.Render("Rounded Border"))

	// DoubleBorder uses double-line characters
	doubleBorder := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("214")) // Orange
	fmt.Println(doubleBorder.Render("Double Border"))
}

// ==============================================================================
// MEDIUM EXAMPLES - Combining styles and layout
// ==============================================================================

// PaddingAndMarginsExample demonstrates spacing control
// Concept: Understanding the difference between padding and margins
func PaddingAndMarginsExample() {
	fmt.Println("\n=== MEDIUM: Padding and Margins ===")

	// Padding adds space INSIDE the border
	// Syntax: Padding(top, right, bottom, left) or Padding(vertical, horizontal)
	withPadding := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("205")).
		Padding(1, 3). // 1 line vertical, 3 spaces horizontal
		Foreground(lipgloss.Color("219"))
	fmt.Println(withPadding.Render("Text with padding"))

	// Margin adds space OUTSIDE the border
	withMargin := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("86")).
		Margin(1, 0). // 1 line vertical margin, 0 horizontal
		Foreground(lipgloss.Color("123"))
	fmt.Println(withMargin.Render("Text with margin"))

	// Combining both padding and margin
	combined := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("214")).
		Padding(0, 2).
		Margin(1, 4).
		Background(lipgloss.Color("235")).
		Foreground(lipgloss.Color("228"))
	fmt.Println(combined.Render("Padding + Margin + Background"))
}

// AlignmentExample shows text alignment options
// Concept: Controlling text alignment within a fixed width
func AlignmentExample() {
	fmt.Println("\n=== MEDIUM: Text Alignment ===")

	// Width sets the fixed width for the styled element
	baseStyle := lipgloss.NewStyle().
		Width(50).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("69"))

	// Left alignment (default)
	leftAlign := baseStyle.Copy().Align(lipgloss.Left)
	fmt.Println(leftAlign.Render("Left aligned text"))

	// Center alignment
	centerAlign := baseStyle.Copy().Align(lipgloss.Center)
	fmt.Println(centerAlign.Render("Center aligned text"))

	// Right alignment
	rightAlign := baseStyle.Copy().Align(lipgloss.Right)
	fmt.Println(rightAlign.Render("Right aligned text"))
}

// JoinExample demonstrates combining multiple styled elements
// Concept: Horizontal and vertical joining of styled blocks
func JoinExample() {
	fmt.Println("\n=== MEDIUM: Joining Elements ===")

	// Create individual styled blocks
	block1 := lipgloss.NewStyle().
		Background(lipgloss.Color("205")).
		Foreground(lipgloss.Color("0")).
		Padding(1, 2).
		Render("Block 1")

	block2 := lipgloss.NewStyle().
		Background(lipgloss.Color("86")).
		Foreground(lipgloss.Color("0")).
		Padding(1, 2).
		Render("Block 2")

	block3 := lipgloss.NewStyle().
		Background(lipgloss.Color("214")).
		Foreground(lipgloss.Color("0")).
		Padding(1, 2).
		Render("Block 3")

	// JoinHorizontal combines elements side by side
	// Parameters: vertical position, elements...
	horizontal := lipgloss.JoinHorizontal(lipgloss.Top, block1, block2, block3)
	fmt.Println(horizontal)

	// JoinVertical stacks elements
	vertical := lipgloss.JoinVertical(lipgloss.Left, block1, block2, block3)
	fmt.Println(vertical)
}

// StyleInheritanceExample shows how to copy and extend styles
// Concept: Reusing base styles with modifications
func StyleInheritanceExample() {
	fmt.Println("\n=== MEDIUM: Style Inheritance ===")

	// Create a base style
	baseStyle := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63"))

	// Copy() creates a new style with all properties of the base
	// Then you can modify specific properties
	successStyle := baseStyle.Copy().
		Foreground(lipgloss.Color("46")). // Green
		BorderForeground(lipgloss.Color("46"))

	warningStyle := baseStyle.Copy().
		Foreground(lipgloss.Color("226")). // Yellow
		BorderForeground(lipgloss.Color("226"))

	errorStyle := baseStyle.Copy().
		Foreground(lipgloss.Color("196")). // Red
		BorderForeground(lipgloss.Color("196"))

	fmt.Println(successStyle.Render("✓ Success message"))
	fmt.Println(warningStyle.Render("⚠ Warning message"))
	fmt.Println(errorStyle.Render("✗ Error message"))
}

// ==============================================================================
// HARD EXAMPLES - Advanced layouts and complex compositions
// ==============================================================================

// ComplexLayoutExample creates a dashboard-like layout
// Concept: Building complex UIs by combining multiple styling techniques
func ComplexLayoutExample() {
	fmt.Println("\n=== HARD: Complex Dashboard Layout ===")

	// Header style
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("230")). // Light yellow
		Background(lipgloss.Color("63")).  // Purple
		Padding(0, 1).
		Width(70).
		Align(lipgloss.Center)

	// Create individual panels
	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("241")).
		Padding(1).
		Width(32).
		Height(8)

	// Stats panel
	statsContent := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Render("📊 Statistics"),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render("Users: 1,234"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("86")).Render("Active: 456"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render("Revenue: $12.3k"),
	)
	statsPanel := panelStyle.Copy().Render(statsContent)

	// Activity panel
	activityContent := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("86")).Render("🔔 Recent Activity"),
		"",
		"• User logged in",
		"• File uploaded",
		"• Task completed",
	)
	activityPanel := panelStyle.Copy().Render(activityContent)

	// Status panel
	statusContent := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("226")).Render("⚡ System Status"),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render("✓ API: Online"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Render("✓ DB: Connected"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render("⚠ Cache: Slow"),
	)
	statusPanel := panelStyle.Copy().Render(statusContent)

	// Alerts panel
	alertsContent := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214")).Render("⚠️  Alerts"),
		"",
		lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("• 3 Failed logins"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render("• Disk 75% full"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render("• Update available"),
	)
	alertsPanel := panelStyle.Copy().Render(alertsContent)

	// Combine panels in rows
	topRow := lipgloss.JoinHorizontal(lipgloss.Top, statsPanel, " ", activityPanel)
	bottomRow := lipgloss.JoinHorizontal(lipgloss.Top, statusPanel, " ", alertsPanel)

	// Combine everything
	dashboard := lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render("🚀 Application Dashboard"),
		"",
		topRow,
		"",
		bottomRow,
	)

	fmt.Println(dashboard)
}

// ProgressBarExample creates animated-looking progress bars
// Concept: Using width and background to create visual indicators
func ProgressBarExample() {
	fmt.Println("\n=== HARD: Progress Bars ===")

	// Create a progress bar function
	renderProgressBar := func(label string, percent int, color lipgloss.Color) string {
		// Calculate filled and empty portions
		totalWidth := 40
		filledWidth := totalWidth * percent / 100
		emptyWidth := totalWidth - filledWidth

		// Create filled portion
		filled := lipgloss.NewStyle().
			Background(color).
			Render(strings.Repeat(" ", filledWidth))

		// Create empty portion
		empty := lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Render(strings.Repeat(" ", emptyWidth))

		// Create percentage label
		percentLabel := lipgloss.NewStyle().
			Foreground(color).
			Bold(true).
			Width(6).
			Align(lipgloss.Right).
			Render(fmt.Sprintf("%d%%", percent))

		// Create label
		labelStyle := lipgloss.NewStyle().
			Width(12).
			Foreground(lipgloss.Color("250")).
			Render(label)

		// Combine everything
		bar := lipgloss.JoinHorizontal(lipgloss.Left, filled, empty)
		return lipgloss.JoinHorizontal(lipgloss.Left, labelStyle, " [", bar, "] ", percentLabel)
	}

	// Display multiple progress bars
	fmt.Println(renderProgressBar("CPU", 67, lipgloss.Color("46")))
	fmt.Println(renderProgressBar("Memory", 82, lipgloss.Color("214")))
	fmt.Println(renderProgressBar("Disk", 45, lipgloss.Color("86")))
	fmt.Println(renderProgressBar("Network", 91, lipgloss.Color("196")))
}

// TableExample creates a formatted table with styling
// Concept: Using lipgloss to create structured data displays
func TableExample() {
	fmt.Println("\n=== HARD: Styled Table ===")

	// Define table data
	headers := []string{"Name", "Role", "Status", "Score"}
	rows := [][]string{
		{"Alice Johnson", "Engineer", "Active", "95"},
		{"Bob Smith", "Designer", "Active", "88"},
		{"Carol White", "Manager", "Away", "92"},
		{"David Brown", "Developer", "Active", "87"},
	}

	// Header style
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("63")).
		Padding(0, 1).
		Width(18).
		Align(lipgloss.Center)

	// Cell styles
	cellStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		Padding(0, 1).
		Width(18)

	alternateStyle := cellStyle.Copy().
		Background(lipgloss.Color("235"))

	// Render headers
	var headerRow []string
	for _, header := range headers {
		headerRow = append(headerRow, headerStyle.Render(header))
	}
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, headerRow...))

	// Render rows
	for i, row := range rows {
		var styledCells []string
		for _, cell := range row {
			// Alternate row colors
			if i%2 == 0 {
				styledCells = append(styledCells, cellStyle.Render(cell))
			} else {
				styledCells = append(styledCells, alternateStyle.Render(cell))
			}
		}
		fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, styledCells...))
	}
}

// AdaptiveLayoutExample demonstrates responsive-like behavior
// Concept: Adjusting styles based on content or conditions
func AdaptiveLayoutExample() {
	fmt.Println("\n=== HARD: Adaptive Layout ===")

	// Function to create a notification card based on type
	createNotification := func(notifType, title, message string) string {
		var (
			icon        string
			color       lipgloss.Color
			borderColor lipgloss.Color
		)

		// Adapt style based on notification type
		switch notifType {
		case "success":
			icon = "✓"
			color = lipgloss.Color("46")
			borderColor = lipgloss.Color("46")
		case "warning":
			icon = "⚠"
			color = lipgloss.Color("226")
			borderColor = lipgloss.Color("226")
		case "error":
			icon = "✗"
			color = lipgloss.Color("196")
			borderColor = lipgloss.Color("196")
		case "info":
			icon = "ℹ"
			color = lipgloss.Color("86")
			borderColor = lipgloss.Color("86")
		default:
			icon = "•"
			color = lipgloss.Color("252")
			borderColor = lipgloss.Color("241")
		}

		// Create icon style
		iconStyle := lipgloss.NewStyle().
			Foreground(color).
			Bold(true).
			Width(3).
			Align(lipgloss.Center)

		// Create title style
		titleStyle := lipgloss.NewStyle().
			Foreground(color).
			Bold(true)

		// Create message style
		messageStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")).
			Width(50)

		// Create border style
		boxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2)

		// Compose the notification
		content := lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Top, iconStyle.Render(icon), " ", titleStyle.Render(title)),
			"",
			messageStyle.Render(message),
		)

		return boxStyle.Render(content)
	}

	// Display different notification types
	fmt.Println(createNotification("success", "Success!", "Your changes have been saved successfully."))
	fmt.Println()
	fmt.Println(createNotification("warning", "Warning", "Your session will expire in 5 minutes."))
	fmt.Println()
	fmt.Println(createNotification("error", "Error", "Failed to connect to the database."))
	fmt.Println()
	fmt.Println(createNotification("info", "Information", "New updates are available for download."))
}

// RunAllLipglossExamples executes all lipgloss examples
func RunAllLipglossExamples() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("LIPGLOSS EXAMPLES - Terminal UI Styling")
	fmt.Println(strings.Repeat("=", 70))

	// Easy examples
	SimpleLipglossExample()
	BasicColorsExample()
	SimpleBordersExample()

	// Medium examples
	PaddingAndMarginsExample()
	AlignmentExample()
	JoinExample()
	StyleInheritanceExample()

	// Hard examples
	ComplexLayoutExample()
	ProgressBarExample()
	TableExample()
	AdaptiveLayoutExample()

	fmt.Println("\n" + strings.Repeat("=", 70))
}
