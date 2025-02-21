package main

import (
	"fmt"
	"log"

	rssemaildigest "github.com/dkvz/rss-email-digest"
)

func main() {
	conf, err := rssemaildigest.ConfigFromDotEnv()
	if err != nil {
		log.Fatal("Could not load configuration, check if .env exists")
	}

	fmt.Printf("%v\n", conf)
}
