package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"sort"
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
		if item.Type != 1 || len(item.Login.URIs) == 0 {
			continue
		}
		uri := item.Login.URIs[0].URI
		jobs = append(jobs, models.Job{
			Password: item.Login.Password,
			Username: item.Login.Username,
			URI:      uri,
			ItemName: item.Name,
		})
	}

	results := runWorkerPool(jobs)

	sortResults(results)

	printResults(results)

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
