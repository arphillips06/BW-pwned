package hibp

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

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

func GetPwned(prefix string) {
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
	lines := strings.Split(string(body), "\n")
	for i := 0; i < 10; i++ {
		fmt.Println(lines[i])
	}
}
