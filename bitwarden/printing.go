package bitwarden

import (
	"bw-hibp-check/models"
	"fmt"
	"strings"
)

func printResults(results []models.Result) {
	fmt.Println("Show passwords in output? (y/n):")
	var choice string
	fmt.Scanln(&choice)
	showPw := strings.ToLower(choice) == "y"

	for _, r := range results {
		if !r.Pwned {
			continue
		}

		pw := r.Password
		if !showPw {
			pw = "********"
		}

		fmt.Println("BREACHED")
		fmt.Println("Account:", r.URI)
		fmt.Println("Username:", r.Username)
		fmt.Println("Password:", pw)
		fmt.Println("Seen in breaches:", r.PwnedCount)
		fmt.Println()
	}
}
