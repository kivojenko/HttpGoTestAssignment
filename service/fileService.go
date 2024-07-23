package service

import (
	"HttpGoTestAssignment/model"
	"HttpGoTestAssignment/utils"
)

func GetReversedFiles() []model.File {
	urls := utils.ReadURLsFromFile()
	contents := downloadContents(urls)
	reversedFiles := utils.GetReversedFiles(urls, contents)

	return reversedFiles
}
