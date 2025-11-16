package bitwarden

import (
	"fmt"
	"log"

	"github.com/arphillips06/bw-pwned/helper"
	"github.com/arphillips06/bw-pwned/models"
)

const bwBaseURL = "http://localhost:8087"

func GetStatus() (*models.VaultStatus, error) {
	var resp models.VaultStatus
	url := bwBaseURL + "/status"
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, fmt.Errorf("status request failed: %w", err)
	}
	return &resp, nil
}

func UnlockVault(password string) (*models.UnlockResponse, error) {
	var resp models.UnlockResponse
	url := bwBaseURL + "/unlock"
	body := models.UnlockRequest{Password: password}
	if err := helper.DoRequest("POST", url, body, &resp); err != nil {
		return nil, fmt.Errorf("unlock vault failed: %w", err)
	}
	log.Printf("Unlocked: %v | Message: %s", resp.Success, resp.Data.Title)
	return &resp, nil
}

func GetItem(id string) (*models.BitwardenItemResponse, error) {
	var resp models.BitwardenItemResponse
	url := bwBaseURL + "/object/item/" + id
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, fmt.Errorf("get item %s failed: %w", id, err)
	}
	return &resp, nil
}
