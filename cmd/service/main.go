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
	log.Println("Routes successfully registered")
	log.Println("Start serving server on port :4040")

	log.Fatal(http.ListenAndServe(":4040", router))

	//service := service2.NewExternalService(10, time.Second)
	//log.Println("started 1")
	//err := service.Process(context.Background(), batch.Batch{batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}})
	//if err != nil {
	//	log.Println(1)
	//	log.Fatal(err)
	//}
	//log.Println("started 2")
	//err = service.Process(context.Background(), batch.Batch{batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}})
	//if err != nil {
	//	log.Println(2)
	//	log.Println(err)
	//}
	//log.Println("wait second")
	//time.Sleep(time.Second)
	//log.Println("started 3")
	//err = service.Process(context.Background(), batch.Batch{batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}, batch.Item{}})
	//if err != nil {
	//	log.Println(2)
	//	log.Println(err)
	//}
	//
	//time.Sleep(time.Second * 10)
}
