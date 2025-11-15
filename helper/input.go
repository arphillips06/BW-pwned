package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func PromptInt(prompt string) int {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if n, err := strconv.Atoi(line); err == nil {
			return n
		}

		fmt.Println("Invalid number. Try again.")
	}
}

func PromptString(prompt string) string {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line != "" {
			return line
		}

		fmt.Println("Input cannot be empty.")
	}
}

func PromptItemID() string {
	return PromptString("Enter item ID: ")
}

func PromptPassword() string {
	return PromptString("Enter Bitwarden password: ")
}

func PromptYesNo(prompt string) bool {
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(strings.ToLower(line))

		if line == "y" || line == "yes" {
			return true
		}
		if line == "n" || line == "no" {
			return false
		}

		fmt.Println("Please enter y/n.")
	}
}
