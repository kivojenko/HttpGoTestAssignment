package controller

import (
	"HttpGoTestAssignment/service"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetFiles(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	contents := service.GetReversedFiles()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(contents)
	if err != nil {
		panic(fmt.Errorf("failed to print contents in controller %w", err))
	}
}
