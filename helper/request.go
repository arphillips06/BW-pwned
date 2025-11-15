package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func DoRequest(method, url string, body interface{}, v interface{}) error {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("%s %s: marshal body: %w", method, url, err)
		}
		reqBody = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("%s %s: build request: %w", method, url, err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s %s: do request: %w", method, url, err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s %s: read body: %w", method, url, err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf(
			"%s %s: unexpected status %d: %s",
			method, url, resp.StatusCode, string(bodyBytes),
		)
	}
	if v == nil {
		return nil
	}
	if len(bodyBytes) == 0 || string(bodyBytes) == "null" {
		return nil
	}
	if err := json.Unmarshal(bodyBytes, v); err != nil {
		return fmt.Errorf("%s %s: unmarshal response: %w\nBody: %s",
			method, url, err, string(bodyBytes))
	}

	return nil
}
