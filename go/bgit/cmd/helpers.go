package cmd

import (
	"fmt"
	"os"
)

// exitWithError standardizes error reporting. In a production-grade CLI we want
// consistent, user-friendly messages while still allowing rich errors internally.
// Using stderr distinguishes errors from normal command output making piping safer.
func exitWithError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

// NOTE: As functionality grows consider extracting reusable output formatting
// (colors, table layouts, column wrapping) into a dedicated package. Keeping
// helpers minimal here maintains approachability for learning phase.
