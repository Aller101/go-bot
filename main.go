package main

import (
	"flag"
	"fmt"
	"log"
	"read-adviser-bot/clients/telegram"
)

func main() {

	host, token := mustToken()

	fmt.Println(host, " ", token)

	tgClient := telegram.New(host, token)

	fmt.Println(tgClient)

	// processor := processor.New(tgClient)

	// fetcher := fetcher.New(tgClient)

	// consumer.Start(fetcher, processor)

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
