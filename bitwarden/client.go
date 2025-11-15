package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/hibp"
	"bw-hibp-check/models"
	"fmt"
	"log"
	"sort"
	"sync"
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

func ListAllItems() (*models.BitwardenItemsListResponse, error) {
	jobs := make(chan models.Job)
	results := make(chan models.Result)
	var all []models.Result
	go func() {
		for r := range results {
			all = append(all, r)
		}
	}()
	const numWorkers = 20
	wg := new(sync.WaitGroup)
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(WorkerID int) {
			for job := range jobs {
				hash, prefix, suffix, count := hibp.CheckPassword(job.Password)
				result := models.Result{
					Username:   job.Username,
					URI:        job.URI,
					ItemName:   job.ItemName,
					Password:   job.Password,
					PwnedCount: uint64(count),
					Prefix:     prefix,
					Suffix:     suffix,
					Hash:       hash,
					Pwned:      count > 0,
					WorkerID:   fmt.Sprintf("worker-%d", WorkerID),
				}
				results <- result
			}
			wg.Done()
		}(i)
	}
	var resp models.BitwardenItemsListResponse
	err := helper.DoRequest("GET", "http://localhost:8087/list/object/items", nil, &resp)
	if err != nil {
		return nil, err
	}
	for _, item := range resp.Data.Data {
		if item.Type != 1 {
			continue
		}
		if len(item.Login.URIs) == 0 {
			continue
		}
		name := item.Login.URIs[0]
		jobs <- models.Job{
			Password: item.Login.Password,
			Username: item.Login.Username,
			URI:      name.URI,
			ItemName: item.Name,
		}
	}
	close(jobs)
	wg.Wait()
	close(results)
	sort.Slice(all, func(i, j int) bool {
		return all[i].PwnedCount > all[j].PwnedCount
	})
	for _, r := range all {
		if r.Pwned {
			fmt.Printf("BREACHED \n")
			fmt.Printf("Account name: %s\n", r.URI)
			fmt.Printf("Username: %s\n", r.Username)
			fmt.Printf("Password: %s\n", r.Password)
			fmt.Printf("Seen in breaches %d times\n", r.PwnedCount)
			fmt.Println()
		}
	}
	return &resp, nil
}
