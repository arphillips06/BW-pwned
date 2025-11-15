package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"fmt"
	"log"
)

func GetStatus() (*models.VaultStatus, error) {
	var resp models.VaultStatus
	err := helper.DoRequest("GET", "http://localhost:8087/status", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func UnlockVault(password string) (*models.UnlockResponse, error) {
	var resp models.UnlockResponse
	err := helper.DoRequest("POST", "http://localhost:8087/unlock",
		models.UnlockRequest{Password: password}, &resp)
	if err != nil {
		return nil, err
	}
	log.Printf("Unlocked: %v | Message: %s", resp.Success, resp.Data.Title)
	return &resp, nil
}

func GetItem(id string) (*models.BitwardenItemResponse, error) {
	var resp models.BitwardenItemResponse
	url := fmt.Sprintf("http://localhost:8087/object/item/%s", id)
	if err := helper.DoRequest("GET", url, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
