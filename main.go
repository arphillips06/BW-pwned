package main

import (
	"fmt"

	"github.com/arphillips06/bw-pwned/bitwarden"
	"github.com/arphillips06/bw-pwned/helper"
)

func promptMenu() int {
	fmt.Println("What do you want to do?")
	fmt.Println("  1. Check status")
	fmt.Println("  2. Get single item")
	fmt.Println("  3. List all items")
	return helper.PromptInt("Choose an option [1-3]: ")
}

func main() {
	fmt.Println("Bitwarden → HIBP checker starting ...")
	choice := promptMenu()
	switch choice {
	case 1:
		bitwarden.HandleCheckStatus()
	case 2:
		bitwarden.HandleGetSingleItem()
	case 3:
		bitwarden.HandleListAllItems()
	default:
		fmt.Println("Invalid choice. Exiting.")
	}
	fmt.Println("Done.")
}
