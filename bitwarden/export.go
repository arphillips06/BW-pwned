package bitwarden

import (
	"bw-hibp-check/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func askExportChoice() int {
	fmt.Println("Export Results? Will show passwords in export")
	fmt.Println("  1. CSV")
	fmt.Println("  2. JSON")
	fmt.Println("  3. No")
	fmt.Print("Choose [1-3]: ")

	var choice int
	fmt.Scanln(&choice)
	return choice
}

func exportCSV(results []models.Result) {
	path := "./exports"
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Println(err)
	}

	filename := fmt.Sprintf("exports/hibp_results_%s.csv",
		time.Now().Format("2006-01-02_15-04-05"),
	)
	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	f.WriteString("uri,username,password,pwned_count\n")

	for _, r := range results {
		if !r.Pwned {
			continue
		}
		line := r.URI + "," +
			r.Username + "," +
			r.Password + "," +
			strconv.FormatUint(r.PwnedCount, 10) + "\n"

		f.WriteString(line)
	}

	fmt.Println("CSV export completed:", filename)
}

func exportJSON(results []models.Result) {
	path := "./exports"
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Println(err)
		return
	}

	filename := fmt.Sprintf("exports/hibp_results_%s.json",
		time.Now().Format("2006-01-02_15-04-05"),
	)

	f, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	type exportJSON struct {
		URI        string `json:"uri"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		PwnedCount uint64 `json:"pwnedCount"`
	}

	var data []exportJSON
	for _, r := range results {
		if r.Pwned {
			data = append(data, exportJSON{
				URI:        r.URI,
				Username:   r.Username,
				Password:   r.Password,
				PwnedCount: r.PwnedCount,
			})
		}
	}

	bytes, _ := json.MarshalIndent(data, "", "  ")
	f.Write(bytes)

	fmt.Println("JSON export completed:", filename)
}
