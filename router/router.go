package router

import (
	"HttpGoTestAssignment/configuration"
	"HttpGoTestAssignment/controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func StartServer() {
	cfg := configuration.GetConfiguration()

	router := httprouter.New()
	router.GET("/files", controller.GetFiles)

	port := cfg.HTTPport
	fmt.Println("Starting server at port " + port)
	server := &http.Server{Addr: port, Handler: router}

	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("error starting server:", err))
	}
}
