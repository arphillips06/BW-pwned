package helper

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/term"
)

func PromptPassword() string {
	fmt.Printf("Enter vault password (hidden): ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	return string(bytePassword)
}
