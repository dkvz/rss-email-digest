package main

import (
	"fmt"
	"log"

	rssemaildigest "github.com/dkvz/rss-email-digest"
	"github.com/mmcdole/gofeed"
)

func main() {
	conf, err := rssemaildigest.ConfigFromDotEnv()
	if err != nil {
		log.Fatal("Could not load configuration: " + err.Error())
	}

	fp := gofeed.NewParser()
	for _, url := range conf.Urls {
		feed, err := fp.ParseURL(url)
		if err != nil {
			log.Fatal("Error fetching feed: " + err.Error())
		}
		fmt.Printf("%v\n", feed.Items)

	}

}
