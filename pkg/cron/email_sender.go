package cron

import (
	"log"

	"github.com/MociW/store-api-golang/pkg/email"
	"github.com/robfig/cron/v3"
)

type EmailScheduler struct {
	cronScheduler *cron.Cron
	email         email.EmailService
}

func NewEmailScheduler(cronScheduler *cron.Cron, email email.EmailService) *EmailScheduler {
	return &EmailScheduler{
		cronScheduler: cronScheduler,
		email:         email,
	}
}

func (e *EmailScheduler) StartSEmailScheduler() {
	// Schedule the job to run every hour
	_, err := e.cronScheduler.AddFunc("@hourly", func() {
		// err := e.email.Send()
		// if err != nil {
		// 	log.Printf("Error sending emails: %v", err)
		// }
	})
	if err != nil {
		log.Fatalf("Error scheduling email job: %v", err)
	}

	e.cronScheduler.Start()
}
