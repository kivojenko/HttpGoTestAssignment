package utils

import (
	"HttpGoTestAssignment/model"
	"sync"
)

func GetReversedFiles(urls []string, contents *sync.Map) []model.File {
	urlsLen := len(urls)
	filesResponse := make([]model.File, 0, urlsLen)

	for i := urlsLen - 1; i >= 0; i-- {
		url := urls[i]
		if content, exists := contents.Load(url); exists {
			file := model.MakeFile(url, content.([]byte))
			filesResponse = append(filesResponse, file)
		}
	}
	return filesResponse
}
