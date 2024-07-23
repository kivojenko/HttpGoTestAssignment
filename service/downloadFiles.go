package service

import (
	"sync"
)

func downloadContents(urls []string) *sync.Map {
	var wg sync.WaitGroup
	contentMap := &sync.Map{}

	for _, url := range urls {
		wg.Add(1)
		go downloadWorker(url, &wg, contentMap)
	}

	wg.Wait()
	return contentMap
}

func downloadWorker(url string, wg *sync.WaitGroup, contentMap *sync.Map) {
	defer wg.Done()

	body := fetchFileContent(url)
	contentMap.Store(url, body)
}
