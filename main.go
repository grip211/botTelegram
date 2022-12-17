package main

import (
	"flag"
	"log"

	tgClient "botTelegram/clients/telegram"
	event_consumer "botTelegram/consumer/event-consumer"
	"botTelegram/events/telegram"
	"botTelegram/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "shared"
	bathSize    = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stooped", err)

	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")

	}
	return *token
}
