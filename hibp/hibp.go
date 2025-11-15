package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	hibpCache     = map[string]string{} //keys are the prefix, value is the repsponse body
	hibpCacheLock sync.RWMutex
)

func ShaPassword(password string) string {
	sum := sha1.Sum([]byte(password))
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}

func StringSplitter(hash string) (prefix, suffix string) {
	if len(hash) < 5 {
		return hash, ""
	}
	return hash[:5], hash[5:]
}

func GetPwned(prefix string) (string, error) {
	hibpCacheLock.RLock()
	body, ok := hibpCache[prefix]
	hibpCacheLock.RUnlock()
	if ok {
		return body, nil
	}
	url := fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefix)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	body = string(b)
	hibpCacheLock.Lock()
	hibpCache[prefix] = body
	hibpCacheLock.Unlock()
	return body, nil
}

func FindSuffixCount(suffix string, body string) int {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		if len(line) < len(suffix)+2 {
			continue
		}
		hash, countStr, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		if hash == suffix {
			num, _ := strconv.Atoi(strings.TrimSpace(countStr))
			return num
		}
	}
	return 0
}

func CheckPassword(password string) (sha, prefix, suffix string, count int) {
	sha = ShaPassword(password)
	prefix, suffix = StringSplitter(sha)
	body, err := GetPwned(prefix)
	if err != nil {
		fmt.Printf("WARN: HIBP lookup failed for prefix %s (%v) — treating as zero\n", prefix, err)
		body = ""
	}
	count = FindSuffixCount(suffix, body)
	return
}
