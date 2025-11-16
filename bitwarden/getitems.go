package bitwarden

import (
	"fmt"
	"log"

	"github.com/arphillips06/bw-pwned/helper"
)

func ensureVaultUnlocked() {
	status, err := GetStatus()
	if err != nil {
		log.Fatalf("Check status failed: %v", err)
	}

	if status.Data.Template.Status == "locked" {
		password := helper.PromptPassword()
		if _, err := UnlockVault(password); err != nil {
			log.Fatalf("Failed to unlock vault: %v", err)
		}
	}
}

func HandleGetSingleItem() {
	ensureVaultUnlocked()
	id := helper.PromptItemID()
	item, err := GetItem(id)
	if err != nil {
		log.Fatalf("Get item failed: %v", err)
	}
	fmt.Printf(
		"\nItem: %s\nUsername: %s\nPassword: %s\n",
		item.Data.Name,
		item.Data.Login.Username,
		item.Data.Login.Password,
	)
}

func HandleListAllItems() {
	ensureVaultUnlocked()
	if _, err := ListAllItems(); err != nil {
		log.Fatalf("List all items failed: %v", err)
	}
}
