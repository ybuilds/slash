package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatalln("error starting server: ", err)
	}
}
