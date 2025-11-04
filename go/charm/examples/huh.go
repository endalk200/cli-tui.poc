package examples

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

// ==============================================================================
// EASY EXAMPLES - Basic form inputs
// ==============================================================================

// SimpleInputExample demonstrates a basic text input
// Concept: Creating a single input field and collecting user input
func SimpleInputExample() {
	fmt.Println("\n=== EASY: Simple Text Input ===")

	var name string

	// Create a form with a single text input
	// NewInput() creates a text input field
	// Title() sets the prompt text
	// Value() binds the input to a variable
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What is your name?").
				Placeholder("Enter your name here...").
				Value(&name),
		),
	)

	// Run() displays the form and waits for user input
	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hello, %s!\n", name)
}

// SimpleConfirmExample demonstrates a yes/no confirmation
// Concept: Getting boolean input from users
func SimpleConfirmExample() {
	fmt.Println("\n=== EASY: Confirmation Dialog ===")

	var confirmed bool

	// NewConfirm() creates a yes/no confirmation dialog
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Do you want to continue?").
				Description("This action cannot be undone.").
				Value(&confirmed),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirmed {
		fmt.Println("✓ Proceeding with the action...")
	} else {
		fmt.Println("✗ Action cancelled.")
	}
}

// SimpleSelectExample demonstrates a single-choice selection
// Concept: Presenting a list of options for user to choose one
func SimpleSelectExample() {
	fmt.Println("\n=== EASY: Single Selection ===")

	var favoriteColor string

	// NewSelect() creates a single-choice selection menu
	// Options() defines the available choices
	// NewOption() creates an option with label and value
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What's your favorite color?").
				Options(
					huh.NewOption("Red", "red"),
					huh.NewOption("Green", "green"),
					huh.NewOption("Blue", "blue"),
					huh.NewOption("Yellow", "yellow"),
				).
				Value(&favoriteColor),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("You selected: %s\n", favoriteColor)
}

// ==============================================================================
// MEDIUM EXAMPLES - Multi-field forms and validation
// ==============================================================================

// MultiFieldFormExample demonstrates a form with multiple inputs
// Concept: Collecting multiple pieces of information in one form
func MultiFieldFormExample() {
	fmt.Println("\n=== MEDIUM: Multi-Field Form ===")

	var (
		username     string
		email        string
		age          string
		agreeToTerms bool
	)

	// NewGroup() groups related fields together
	// Multiple groups can be used to create multi-page forms
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Placeholder("johndoe").
				Value(&username),

			huh.NewInput().
				Title("Email").
				Placeholder("john@example.com").
				Value(&email),

			huh.NewInput().
				Title("Age").
				Placeholder("25").
				Value(&age),

			huh.NewConfirm().
				Title("I agree to the terms and conditions").
				Value(&agreeToTerms),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n--- User Registration ---")
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Age: %s\n", age)
	fmt.Printf("Agreed to Terms: %v\n", agreeToTerms)
}

// ValidationExample demonstrates input validation
// Concept: Ensuring user input meets certain criteria
func ValidationExample() {
	fmt.Println("\n=== MEDIUM: Input Validation ===")

	var (
		email    string
		password string
	)

	form := huh.NewForm(
		huh.NewGroup(
			// Email validation
			huh.NewInput().
				Title("Email Address").
				Placeholder("user@example.com").
				Value(&email).
				Validate(func(s string) error {
					// Simple email validation
					if !strings.Contains(s, "@") || !strings.Contains(s, ".") {
						return fmt.Errorf("please enter a valid email address")
					}
					return nil
				}),

			// Password validation
			huh.NewInput().
				Title("Password").
				Placeholder("Enter a strong password").
				EchoMode(huh.EchoModePassword). // Hide password input
				Value(&password).
				Validate(func(s string) error {
					if len(s) < 8 {
						return fmt.Errorf("password must be at least 8 characters long")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("✓ Validation successful!")
	fmt.Printf("Email: %s\n", email)
	fmt.Printf("Password: %s\n", strings.Repeat("*", len(password)))
}

// MultiSelectExample demonstrates selecting multiple options
// Concept: Allowing users to choose multiple items from a list
func MultiSelectExample() {
	fmt.Println("\n=== MEDIUM: Multiple Selection ===")

	var selectedLanguages []string

	// NewMultiSelect() allows selecting multiple options
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Which programming languages do you know?").
				Description("Select all that apply").
				Options(
					huh.NewOption("Go", "go"),
					huh.NewOption("Python", "python"),
					huh.NewOption("JavaScript", "javascript"),
					huh.NewOption("Java", "java"),
					huh.NewOption("C++", "cpp"),
					huh.NewOption("Rust", "rust"),
					huh.NewOption("TypeScript", "typescript"),
				).
				Value(&selectedLanguages).
				Limit(3), // Optional: limit number of selections
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nYou know these languages:")
	for i, lang := range selectedLanguages {
		fmt.Printf("%d. %s\n", i+1, lang)
	}
}

// TextAreaExample demonstrates multi-line text input
// Concept: Collecting longer text input from users
func TextAreaExample() {
	fmt.Println("\n=== MEDIUM: Text Area Input ===")

	var feedback string

	form := huh.NewForm(
		huh.NewGroup(
			// NewText() creates a multi-line text area
			huh.NewText().
				Title("Please provide your feedback").
				Description("Tell us what you think about our service").
				Placeholder("Enter your feedback here...").
				CharLimit(500). // Optional: limit character count
				Value(&feedback),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n--- Your Feedback ---")
	fmt.Println(feedback)
}

// ==============================================================================
// HARD EXAMPLES - Complex forms and advanced patterns
// ==============================================================================

// MultiPageFormExample demonstrates a multi-step form
// Concept: Creating wizard-like forms with multiple pages
func MultiPageFormExample() {
	fmt.Println("\n=== HARD: Multi-Page Form (Wizard) ===")

	var (
		// Page 1: Personal Information
		firstName string
		lastName  string
		email     string

		// Page 2: Preferences
		theme         string
		notifications bool

		// Page 3: Confirmation
		confirmSubmit bool
	)

	// Each NewGroup() creates a separate page
	form := huh.NewForm(
		// Page 1: Personal Information
		huh.NewGroup(
			huh.NewNote().
				Title("Personal Information").
				Description("Let's start with your basic details"),

			huh.NewInput().
				Title("First Name").
				Value(&firstName).
				Validate(func(s string) error {
					if len(s) < 2 {
						return fmt.Errorf("first name must be at least 2 characters")
					}
					return nil
				}),

			huh.NewInput().
				Title("Last Name").
				Value(&lastName).
				Validate(func(s string) error {
					if len(s) < 2 {
						return fmt.Errorf("last name must be at least 2 characters")
					}
					return nil
				}),

			huh.NewInput().
				Title("Email").
				Value(&email).
				Validate(func(s string) error {
					if !strings.Contains(s, "@") {
						return fmt.Errorf("invalid email address")
					}
					return nil
				}),
		),

		// Page 2: Preferences
		huh.NewGroup(
			huh.NewNote().
				Title("Preferences").
				Description("Customize your experience"),

			huh.NewSelect[string]().
				Title("Theme").
				Options(
					huh.NewOption("Light", "light"),
					huh.NewOption("Dark", "dark"),
					huh.NewOption("Auto", "auto"),
				).
				Value(&theme),

			huh.NewConfirm().
				Title("Enable email notifications").
				Value(&notifications),
		),

		// Page 3: Confirmation
		huh.NewGroup(
			huh.NewNote().
				Title("Review and Confirm").
				Description("Please review your information"),

			huh.NewConfirm().
				Title("Submit your information?").
				Description(fmt.Sprintf(
					"Name: %s %s\nEmail: %s\nTheme: %s\nNotifications: %v",
					firstName, lastName, email, theme, notifications,
				)).
				Value(&confirmSubmit),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if confirmSubmit {
		fmt.Println("\n✓ Registration completed successfully!")
		fmt.Printf("Name: %s %s\n", firstName, lastName)
		fmt.Printf("Email: %s\n", email)
		fmt.Printf("Theme: %s\n", theme)
		fmt.Printf("Notifications: %v\n", notifications)
	} else {
		fmt.Println("\n✗ Registration cancelled.")
	}
}

// DynamicFormExample demonstrates conditional form fields
// Concept: Showing/hiding fields based on user input
func DynamicFormExample() {
	fmt.Println("\n=== HARD: Dynamic Form (Conditional Fields) ===")

	var (
		userType      string
		companyName   string
		jobTitle      string
		studentID     string
		university    string
		freelanceRate string
	)

	// First, ask user type
	typeForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What type of user are you?").
				Options(
					huh.NewOption("Corporate Employee", "corporate"),
					huh.NewOption("Student", "student"),
					huh.NewOption("Freelancer", "freelancer"),
				).
				Value(&userType),
		),
	)

	err := typeForm.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create conditional form based on user type
	var detailsForm *huh.Form

	switch userType {
	case "corporate":
		detailsForm = huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Corporate Information"),

				huh.NewInput().
					Title("Company Name").
					Value(&companyName),

				huh.NewInput().
					Title("Job Title").
					Value(&jobTitle),
			),
		)

	case "student":
		detailsForm = huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Student Information"),

				huh.NewInput().
					Title("University").
					Value(&university),

				huh.NewInput().
					Title("Student ID").
					Value(&studentID),
			),
		)

	case "freelancer":
		detailsForm = huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Freelancer Information"),

				huh.NewInput().
					Title("Hourly Rate (USD)").
					Placeholder("50").
					Value(&freelanceRate).
					Validate(func(s string) error {
						if s == "" {
							return fmt.Errorf("rate is required")
						}
						return nil
					}),
			),
		)
	}

	err = detailsForm.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display results based on user type
	fmt.Println("\n--- User Profile ---")
	fmt.Printf("Type: %s\n", userType)

	switch userType {
	case "corporate":
		fmt.Printf("Company: %s\n", companyName)
		fmt.Printf("Job Title: %s\n", jobTitle)
	case "student":
		fmt.Printf("University: %s\n", university)
		fmt.Printf("Student ID: %s\n", studentID)
	case "freelancer":
		fmt.Printf("Hourly Rate: $%s\n", freelanceRate)
	}
}

// ComplexWorkflowExample demonstrates a complete application workflow
// Concept: Building a full user workflow with multiple forms and logic
func ComplexWorkflowExample() {
	fmt.Println("\n=== HARD: Complete Application Workflow ===")

	var (
		// Authentication
		existingUser bool
		username     string
		password     string

		// Profile
		displayName string
		bio         string
		interests   []string

		// Settings
		privacyLevel string
		theme        string
		language     string
	)

	// Step 1: Authentication
	authForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Welcome! 👋").
				Description("Let's get you set up"),

			huh.NewConfirm().
				Title("Do you have an existing account?").
				Value(&existingUser),
		),
	)

	err := authForm.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Step 2: Login/Register
	credentialsForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Value(&username).
				Validate(func(s string) error {
					if len(s) < 3 {
						return fmt.Errorf("username must be at least 3 characters")
					}
					return nil
				}),

			huh.NewInput().
				Title("Password").
				EchoMode(huh.EchoModePassword).
				Value(&password).
				Validate(func(s string) error {
					if len(s) < 8 {
						return fmt.Errorf("password must be at least 8 characters")
					}
					return nil
				}),
		),
	)

	err = credentialsForm.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if existingUser {
		fmt.Println("✓ Logged in successfully!")
	} else {
		// Step 3: Profile Setup (only for new users)
		profileForm := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					Title("Complete Your Profile").
					Description("Tell us a bit about yourself"),

				huh.NewInput().
					Title("Display Name").
					Placeholder("John Doe").
					Value(&displayName),

				huh.NewText().
					Title("Bio").
					Placeholder("A short bio about yourself...").
					CharLimit(200).
					Value(&bio),

				huh.NewMultiSelect[string]().
					Title("Interests").
					Options(
						huh.NewOption("Technology", "tech"),
						huh.NewOption("Sports", "sports"),
						huh.NewOption("Music", "music"),
						huh.NewOption("Gaming", "gaming"),
						huh.NewOption("Reading", "reading"),
						huh.NewOption("Travel", "travel"),
					).
					Value(&interests),
			),
		)

		err = profileForm.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	// Step 4: Settings
	settingsForm := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("Customize Your Experience").
				Description("Set your preferences"),

			huh.NewSelect[string]().
				Title("Privacy Level").
				Options(
					huh.NewOption("Public - Anyone can see your profile", "public"),
					huh.NewOption("Friends - Only friends can see", "friends"),
					huh.NewOption("Private - Only you can see", "private"),
				).
				Value(&privacyLevel),

			huh.NewSelect[string]().
				Title("Theme").
				Options(
					huh.NewOption("Light", "light"),
					huh.NewOption("Dark", "dark"),
					huh.NewOption("System", "system"),
				).
				Value(&theme),

			huh.NewSelect[string]().
				Title("Language").
				Options(
					huh.NewOption("English", "en"),
					huh.NewOption("Spanish", "es"),
					huh.NewOption("French", "fr"),
					huh.NewOption("German", "de"),
					huh.NewOption("Japanese", "ja"),
				).
				Value(&language),
		),
	)

	err = settingsForm.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Display final summary
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("✓ Setup Complete!")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("\nUsername: %s\n", username)

	if !existingUser {
		fmt.Printf("Display Name: %s\n", displayName)
		fmt.Printf("Bio: %s\n", bio)
		fmt.Printf("Interests: %v\n", interests)
	}

	fmt.Printf("\nSettings:\n")
	fmt.Printf("  Privacy: %s\n", privacyLevel)
	fmt.Printf("  Theme: %s\n", theme)
	fmt.Printf("  Language: %s\n", language)
	fmt.Println(strings.Repeat("=", 60))
}

// FormWithInlineHelpExample demonstrates help text and documentation
// Concept: Providing contextual help to users
func FormWithInlineHelpExample() {
	fmt.Println("\n=== HARD: Form with Inline Help ===")

	var (
		apiKey   string
		endpoint string
		timeout  string
		retries  string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("API Configuration").
				Description("Configure your API settings for optimal performance"),

			huh.NewInput().
				Title("API Key").
				Description("Your secret API key (keep this secure!)").
				Placeholder("sk_live_xxxxxxxxxxxxx").
				EchoMode(huh.EchoModePassword).
				Value(&apiKey).
				Validate(func(s string) error {
					if len(s) < 20 {
						return fmt.Errorf("API key appears to be invalid")
					}
					return nil
				}),

			huh.NewInput().
				Title("API Endpoint").
				Description("The base URL for API requests (e.g., https://api.example.com)").
				Placeholder("https://api.example.com").
				Value(&endpoint).
				Validate(func(s string) error {
					if !strings.HasPrefix(s, "http") {
						return fmt.Errorf("endpoint must start with http:// or https://")
					}
					return nil
				}),

			huh.NewInput().
				Title("Timeout (seconds)").
				Description("How long to wait for a response (default: 30)").
				Placeholder("30").
				Value(&timeout),

			huh.NewInput().
				Title("Max Retries").
				Description("Number of times to retry failed requests (default: 3)").
				Placeholder("3").
				Value(&retries),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n✓ API Configuration saved!")
	fmt.Printf("Endpoint: %s\n", endpoint)
	fmt.Printf("Timeout: %s seconds\n", timeout)
	fmt.Printf("Max Retries: %s\n", retries)
}

// RunAllHuhExamples executes all huh examples
func RunAllHuhExamples() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("HUH EXAMPLES - Interactive Terminal Forms")
	fmt.Println(strings.Repeat("=", 70))

	// Easy examples
	SimpleInputExample()
	SimpleConfirmExample()
	SimpleSelectExample()

	// Medium examples
	MultiFieldFormExample()
	ValidationExample()
	MultiSelectExample()
	TextAreaExample()

	// Hard examples
	MultiPageFormExample()
	DynamicFormExample()
	ComplexWorkflowExample()
	FormWithInlineHelpExample()

	fmt.Println("\n" + strings.Repeat("=", 70))
}
