package model

import (
	"github.com/google/uuid"
	"strings"
)

type File struct {
	ID       string `json:"id"`
	FileName string `json:"file_name"`
	FileURL  string `json:"file_url"`
	Content  string `json:"content"`
}

func MakeFile(url string, content []byte) File {
	fileName := getFileName(url)
	fileContent := string(content)

	return NewFile(fileName, url, fileContent)
}

func NewFile(fileName, fileURL string, content string) File {
	newID := uuid.New().String()
	return File{newID, fileName, fileURL, content}
}

func getFileName(url string) string {
	urlParts := strings.Split(url, "/")
	return urlParts[len(urlParts)-1]
}
