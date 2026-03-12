package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ybuilds/slash/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/api/slash/v1/create", controllers.CreateMapping)
	router.GET("/:encode", controllers.GetMapping)

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatalln("error starting server: ", err)
	}
}
