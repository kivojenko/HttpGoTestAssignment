package service

import (
	"fmt"
	"io"
	"net/http"
)

func fetchFileContent(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Errorf("failed to download %s: %w", url, err))
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to download %s: status %s", url, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read response body from %s: %w", url, err))
	}

	return body
}
