package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var hibpCacheMap = map[string]string{} //keys are the prefix, value is the repsponse body
var hibpCacheMutex sync.Mutex

func ShaPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	shaHash := hex.EncodeToString(hash.Sum(nil))
	return strings.ToUpper(shaHash)
}

func StringSplitter(hash string) (prefix, suffix string) {
	if len(hash) < 5 {
		return hash, ""
	}
	prefix = hash[:5]
	suffix = hash[5:]
	return prefix, suffix
}

func GetPwned(prefix string) string {
	hibpCacheMutex.Lock()
	value, ok := hibpCacheMap[prefix]
	if ok {
		hibpCacheMutex.Unlock()
		fmt.Println("cache hit")
		return value
	}
	hibpCacheMutex.Unlock()
	fmt.Println("FETCHING FROM API:", prefix)
	url := fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", prefix)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}
	sbody := string(body)
	hibpCacheMutex.Lock()
	hibpCacheMap[prefix] = sbody
	hibpCacheMutex.Unlock()
	return sbody
}

func FindSuffixCount(suffix string, body string) int {
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		s := strings.Split(line, ":")
		if len(s) == 2 {
			hash := strings.Trim(s[0], " '\r\n")
			count := strings.TrimSpace(s[1])
			if hash == suffix {
				num, _ := strconv.Atoi(count)
				return num
			}
		}
	}
	return 0
}

func CheckPassword(password string) (hash, prefix, suffix string, count int) {
	sha := ShaPassword(password)
	prefix, suffix = StringSplitter(sha)
	body := GetPwned(prefix)
	getCount := FindSuffixCount(suffix, body)
	return sha, prefix, suffix, getCount
}
