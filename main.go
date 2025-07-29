package main

import (
	"flag"
	"log"
)

func main() {

	token := mustToken()

	// tgClient := telegram.New(token)

	// processor := processor.New(tgClient)

	// fetcher := fetcher.New(tgClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String("t", "", "token for telegram")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
