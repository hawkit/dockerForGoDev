package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func routeHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "This is a demo for go develop environment with docker")
}

func main() {

	router := httprouter.New()
	router.GET("/", routeHandler)
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Server in production mode")
	} else {
		log.Println("Server in dev mode")
	}
	http.ListenAndServe(":8080", router)
}
