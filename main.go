package main

import (
	"flag"
	"fmt"
	"log"
	tg_cl "read-adviser-bot/clients/telegram"
	econsumer "read-adviser-bot/consumer/e-consumer"
	tg_ev "read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	host, token := mustToken()

	fmt.Println(host, " ", token)

	tgClient := tg_cl.New(host, token)

	fmt.Println(tgClient)

	s := files.New(storagePath)

	eventProcesser := tg_ev.New(tgClient, s)

	log.Print("service started\n")

	cons := econsumer.New(eventProcesser, eventProcesser, batchSize)

	if err := cons.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}

}

func mustToken() (string, string) {
	token := flag.String("t", "", "token for telegram")
	host := flag.String("h", "", "host for telegram")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	if *host == "" {
		log.Fatal("host is not specified")
	}

	return *token, *host
}
