//go:build windows

package main

import (
	"log"

	"github.com/go-co-op/gocron/v2"
	"github.com/scncore/scncore-ocsp-responder/internal/common"
	"github.com/scncore/utils"
	"golang.org/x/sys/windows/svc"
)

func main() {
	var err error
	w := common.NewWorker("scncore-ocsp-responder.txt")

	// Start Task Scheduler
	w.TaskScheduler, err = gocron.NewScheduler()
	if err != nil {
		log.Printf("[ERROR]: could not create task scheduler, reason: %s", err.Error())
		return
	}
	w.TaskScheduler.Start()
	log.Println("[INFO]: task scheduler has been started")

	if err := w.GenerateOCSPResponderConfig(); err != nil {
		log.Printf("[ERROR]: could not generate config for OCSP responder: %v", err)
		if err := w.StartGenerateOCSPResponderConfigJob(); err != nil {
			log.Fatalf("[FATAL]: could not start job to generate config for OCSP responder: %v", err)
		}
	}

	s := utils.NewscncoreWindowsService()
	s.ServiceStart = w.StartWorker
	s.ServiceStop = w.StopWorker

	// Run service

	if err := svc.Run("scncore-ocsp-responder", s); err != nil {
		log.Printf("[ERROR]: could not run service: %v", err)
	}
}
