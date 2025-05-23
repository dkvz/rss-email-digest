package main

import (
	"fmt"
	"log"
	"time"

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
	for {
		for _, url := range conf.Urls {
			feed, err := fp.ParseURL(url)
			if err != nil {
				log.Fatal("Error fetching feed: " + err.Error())
			}
			if len(feed.Items) == 0 {
				// Ignore this run and print a warning:
				log.Printf("feed %v was empty - ignoring\n", url)
				continue
			}

			// Check first item:
			if state.IsNewGUID(url, feed.Items[0].GUID) {
				latestGuid := state.LatestGUID(url)
				err := state.SaveLastestGUID(url, feed.Items[0].GUID)
				if err != nil {
					log.Fatal("cannot write state file")
				}
				// Get all the feeds until we reach the latest GUID:

			}

			for _, item := range feed.Items {
				fmt.Printf("%v; %v; %v; %s\n", item.Title, item.Link, item.GUID, item.Published)
			}
		}
		time.Sleep(time.Duration(conf.SleepInterval) * time.Second)
	}

}
