package main

import (
	"bw-hibp-check/bitwarden"
	"bw-hibp-check/helper"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Bitwarden → HIBP checker starting ...")
	fmt.Println("What do you want to do?")
	fmt.Println("  1. Check status")
	fmt.Println("  2. Get single item")
	fmt.Println("  3. List all items")
	fmt.Print("Choose an option [1-3]: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		bitwarden.GetStatus()

	case 2:
		status, err := bitwarden.GetStatus()
		if err != nil {
			fmt.Printf("failed to get status: %s", err)
		}
		if status.Data.Template.Status == "locked" {
			password := helper.PromptPassword()
			_, err := bitwarden.UnlockVault(password)
			if err != nil {
				fmt.Printf("failed to unlock vault: %s", err)
			}
		}
		fmt.Print("Enter Bitwarden item ID: ")
		var itemID string
		fmt.Scanln(&itemID)

		item, err := bitwarden.GetItem(itemID)
		if err != nil {
			log.Fatalf("Get item failed: %v", err)
		}

		fmt.Printf("\nItem: %s\nUsername: %s\nPassword: %s\n",
			item.Data.Name, item.Data.Login.Username, item.Data.Login.Password)

	case 3:
		status, err := bitwarden.GetStatus()
		if err != nil {
			fmt.Printf("failed to get status: %s", err)
		}
		if status.Data.Template.Status == "locked" {
			password := helper.PromptPassword()
			_, err := bitwarden.UnlockVault(password)
			if err != nil {
				fmt.Printf("failed to unlock vault: %s", err)
			}
		}
		itemsResp, err := bitwarden.ListAllItems()
		if err != nil {
			log.Fatalf("List all items failed: %v", err)
		}
		fmt.Printf("Found %d vault items\n", len(itemsResp.Data.Data))

	default:
		fmt.Println("Invalid choice. Exiting.")
	}
}
