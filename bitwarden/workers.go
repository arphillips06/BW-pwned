package bitwarden

import (
	"bw-hibp-check/hibp"
	"bw-hibp-check/models"
	"fmt"
	"sync"
)

func runWorkerPool(jobs []models.Job) []models.Result {
	jobChan := make(chan models.Job)
	resultChan := make(chan models.Result)

	const numWorkers = 20
	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			for job := range jobChan {
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
					WorkerID:   fmt.Sprintf("worker-%d", workerID),
				}

				resultChan <- result
			}
			wg.Done()
		}(i)
	}
	var all []models.Result
	done := make(chan struct{})

	go func() {
		for r := range resultChan {
			all = append(all, r)
		}
		close(done)
	}()
	go func() {
		for _, j := range jobs {
			jobChan <- j
		}
		close(jobChan)
	}()
	wg.Wait()
	close(resultChan)
	<-done

	return all
}
