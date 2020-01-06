package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jasondavindev/godontpadservice/services"
	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	services.GetCommands().Exec()
	time.Sleep(time.Second * 10)
	p.run()
}

func (p *program) Stop(s service.Service) error {
	fmt.Println("Stopping golang dontpad service")
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoDontpadService",
		DisplayName: "Go dontpad service",
		Description: "Go dontpad service",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)

	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)

	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
