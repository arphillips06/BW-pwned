package bitwarden

import (
	"fmt"
	"sort"
	"time"

	"github.com/arphillips06/bw-pwned/helper"
	"github.com/arphillips06/bw-pwned/models"
)

func sortResults(results []models.Result) {
	sort.Slice(results, func(i, j int) bool {
		return results[i].PwnedCount > results[j].PwnedCount
	})
}

func ListAllItems() (*models.BitwardenItemsListResponse, error) {
	var resp models.BitwardenItemsListResponse
	err := helper.DoRequest("GET", "http://localhost:8087/list/object/items", nil, &resp)
	if err != nil {
		return nil, err
	}
	var jobs []models.Job
	for _, item := range resp.Data.Data {
		if item.Type != 1 {
			continue
		}
		if len(item.Login.URIs) == 0 {
			continue
		}
		if item.Login.Password == "" && len(item.Login.Fido2Credentials) > 0 {
			continue
		}
		jobs = append(jobs, models.Job{
			Password: item.Login.Password,
			Username: item.Login.Username,
			URI:      item.Login.URIs[0].URI,
			ItemName: item.Name,
		})
	}
	start := time.Now()
	results := runWorkerPool(jobs)
	duration := time.Since(start)
	sortResults(results)
	printResults(results)
	fmt.Printf("Scan completed in %s\n", duration.Round(time.Millisecond))
	exportChoice := askExportChoice()
	switch exportChoice {
	case 1:
		exportCSV(results)
	case 2:
		exportJSON(results)
	case 3:
	}
	return &resp, nil
}
