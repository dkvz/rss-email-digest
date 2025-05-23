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

	// Create the state:
	state := rssemaildigest.State{}

	fp := gofeed.NewParser()
	for _, url := range conf.Urls {
		feed, err := fp.ParseURL(url)
		if err != nil {
			log.Fatal("Error fetching feed: " + err.Error())
		}
		for _, item := range feed.Items {
			fmt.Printf("%v; %v; %v; %s\n", item.Title, item.Link, item.GUID, item.Published)
		}

	}

}
