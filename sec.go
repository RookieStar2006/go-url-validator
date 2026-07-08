package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your url: ")
	scanner.Scan()
	inputURL := strings.TrimSpace(scanner.Text())

	parsedURL, err := url.Parse(inputURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		fmt.Println("❌ Format galat hai! Example: https://google.com")
		return
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(parsedURL.String())
	if err != nil {
		fmt.Printf("❌ URL is not reachable.: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Printf("✅ Valid & Working URL! (Status: %d)\n", resp.StatusCode)
	} else {
		fmt.Printf("⚠️ got URL but status %d aaya\n", resp.StatusCode)
	}
}