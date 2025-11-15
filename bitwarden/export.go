package bitwarden

import (
	"bw-hibp-check/helper"
	"bw-hibp-check/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ensureExportDir() {
	if err := os.MkdirAll("exports", 0755); err != nil {
		log.Println("Failed to create exports directory:", err)
	}
}

func escapeCSV(v string) string {
	v = strings.ReplaceAll(v, `"`, `""`)
	return `"` + v + `"`
}

func timestampedFilename(ext string) string {
	return fmt.Sprintf("exports/hibp_results_%s.%s",
		time.Now().Format("2006-01-02_17-04-05"),
		ext,
	)
}

func askExportChoice() int {
	fmt.Println("Export Results? Will show passwords in export")
	fmt.Println("  1. CSV")
	fmt.Println("  2. JSON")
	fmt.Println("  3. No")
	return helper.PromptInt("Choose [1-3]: ")
}

func exportCSV(results []models.Result) {
	ensureExportDir()
	filename := timestampedFilename("csv")
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Failed to create CSV file:", err)
		return
	}
	defer f.Close()
	_, _ = f.WriteString("uri,username,password,pwned_count\n")
	for _, r := range results {
		if !r.Pwned {
			continue
		}
		line := escapeCSV(r.URI) + "," +
			escapeCSV(r.Username) + "," +
			escapeCSV(r.Password) + "," +
			strconv.FormatUint(r.PwnedCount, 10) + "\n"

		_, _ = f.WriteString(line)
	}
	fmt.Println("CSV export completed:", filename)
}

func exportJSON(results []models.Result) {
	ensureExportDir()
	filename := timestampedFilename("json")
	f, err := os.Create(filename)
	if err != nil {
		log.Println("Failed to create JSON file:", err)
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
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Failed to marshal JSON:", err)
		return
	}
	if _, err := f.Write(bytes); err != nil {
		log.Println("Failed to write JSON:", err)
		return
	}
	fmt.Println("JSON export completed:", filename)
}

func ExportResults(results []models.Result) {
	switch askExportChoice() {
	case 1:
		exportCSV(results)
	case 2:
		exportJSON(results)
	default:
		fmt.Println("No export.")
	}
}
