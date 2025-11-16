package bitwarden

import (
	"fmt"
	"sync"

	"github.com/arphillips06/bw-pwned/hibp"
	"github.com/arphillips06/bw-pwned/models"
)

func runWorkerPool(jobs []models.Job) []models.Result {
	const numWorkers = 20
	jobChan := make(chan models.Job)
	resultChan := make(chan models.Result)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := range numWorkers {
		workerID := fmt.Sprintf("worker-%d", i)
		go func(id string) {
			defer wg.Done()
			for job := range jobChan {
				hash, prefix, suffix, count := hibp.CheckPassword(job.Password)
				resultChan <- models.Result{
					Username:   job.Username,
					URI:        job.URI,
					ItemName:   job.ItemName,
					Password:   job.Password,
					PwnedCount: uint64(count),
					Prefix:     prefix,
					Suffix:     suffix,
					Hash:       hash,
					Pwned:      count > 0,
					WorkerID:   id,
				}
			}
		}(workerID)
	}
	go func() {
		for _, j := range jobs {
			jobChan <- j
		}
		close(jobChan)
	}()
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	var results []models.Result
	for r := range resultChan {
		results = append(results, r)
	}
	return results
}
