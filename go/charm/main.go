package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/endalk200/charm.poc/examples"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n" + strings.Repeat("=", 70))
		fmt.Println("CHARM PACKAGE EXAMPLES - Learning & Experimentation")
		fmt.Println(strings.Repeat("=", 70))
		fmt.Println("\nSelect which package examples to run:")
		fmt.Println("1. Lipgloss - Terminal UI Styling")
		fmt.Println("2. Log - Structured Logging")
		fmt.Println("3. Huh - Interactive Forms")
		fmt.Println("4. All Examples")
		fmt.Println("5. Exit")
		fmt.Print("\nEnter your choice (1-5): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			examples.RunAllLipglossExamples()
		case "2":
			examples.RunAllLogExamples()
		case "3":
			examples.RunAllHuhExamples()
		case "4":
			examples.RunAllLipglossExamples()
			examples.RunAllLogExamples()
			examples.RunAllHuhExamples()
		case "5":
			fmt.Println("\nGoodbye! 👋")
			return
		default:
			fmt.Println("\n❌ Invalid choice. Please enter a number between 1 and 5.")
		}

		fmt.Print("\nPress Enter to continue...")
		reader.ReadString('\n')
	}
}
