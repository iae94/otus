package service

//Каркас микросервиса
//Реализовать "каркас" микросервиса, считывающий конфиг из файла,
//создающий логгер/логгеры с указанными уровнями детализации.

import (
	serviceConfig "Microservice/internal/config"
	serviceLogger "Microservice/internal/logger"
	"Microservice/internal/routes"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type Microservice interface {
	Init()
	Start()
}

type WebService struct {
	Config   *serviceConfig.Config
	Logger   *zap.Logger
	Handlers map[string]func(http.ResponseWriter, *http.Request)
}

func (s *WebService) Init() error {

	config, err := serviceConfig.ReadConfig()
	if err != nil {
		log.Printf("Reading config error: %v\n", err)
		return err
	}

	logger, err := serviceLogger.CreateLogger()
	if err != nil {
		log.Printf("Create logger error: %v\n", err)
		return err
	}

	s.Config = config
	s.Logger = logger

	if s.Handlers == nil || len(s.Handlers) == 0 {
		s.Handlers = make(map[string]func(http.ResponseWriter, *http.Request))
		s.Handlers["/"] = routes.HelloHandler
	}

	sugar := s.Logger.Sugar()
	for url, handler := range s.Handlers {
		http.HandleFunc(url, handler)
		sugar.Infof("Register handler: %v", url)
	}

	return nil
}

func (s *WebService) Start() error {
	defer s.Logger.Sync()
	err := http.ListenAndServe(fmt.Sprintf(":%v", s.Config.Port), nil)
	if err != nil {
		s.Logger.Error("Listener error: %v\n", zap.String("err", err.Error()))
	}
	return err
}
