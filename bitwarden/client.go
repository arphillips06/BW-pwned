package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"fmt"
	"log"
	"path"
)

const bwBaseURL = "http://localhost:8087"

func GetStatus() (*models.VaultStatus, error) {
	var resp models.VaultStatus
	url := path.Join(bwBaseURL, "status")
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, fmt.Errorf("status request failed: %w", err)
	}
	return &resp, nil
}

func UnlockVault(password string) (*models.UnlockResponse, error) {
	var resp models.UnlockResponse
	url := path.Join(bwBaseURL, "unlock")
	body := models.UnlockRequest{Password: password}
	if err := helper.DoRequest("POST", url, body, &resp); err != nil {
		return nil, fmt.Errorf("unlock vault failed: %w", err)
	}
	log.Printf("Unlocked: %v | Message: %s", resp.Success, resp.Data.Title)
	return &resp, nil
}

func GetItem(id string) (*models.BitwardenItemResponse, error) {
	var resp models.BitwardenItemResponse
	url := path.Join(bwBaseURL, "object", "item", id)
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, fmt.Errorf("get item %s failed: %w", id, err)
	}
	return &resp, nil
}
