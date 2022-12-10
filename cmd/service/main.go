package main

import (
	service2 "github.com/DiasOrazbaev/kazanexpress-test-task/internal/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Creating instance of service...")
	service := service2.NewExternalService(10, time.Second*5)
	log.Println("Service instance currently created")

	log.Println("Registering routes...")
	router := httprouter.New()

	router.POST("/batch", service.BatchHandler)
	router.GET("/batch", service.GetLimitsHandler)

	log.Println("Routes successfully registered")
	log.Println("Start serving server on port :4040")

	log.Fatal(http.ListenAndServe(":4040", router))
}
