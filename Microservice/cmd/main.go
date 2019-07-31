package main

//Каркас микросервиса
//Реализовать "каркас" микросервиса, считывающий конфиг из файла,
//создающий логгер/логгеры с указанными уровнями детализации.

import (
	"Microservice/internal/service"
	"log"
)

func main() {

	webService := service.WebService{}
	err := webService.Init()
	if err != nil {
		log.Fatalf("WebService initialization has failed with error: %v\n", err)
	}
	webService.Logger.Sugar().Infof("Start service at port: %v", webService.Config.Port)
	err = webService.Start()
	if err != nil {
		log.Fatalf("WebService Fatal error: %v\n", err)
	}
}
