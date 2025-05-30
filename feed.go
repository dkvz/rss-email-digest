package rssemaildigest

import (
	"fmt"
	"log"
	"strings"

	"github.com/dkvz/rss-email-digest/notifications"
	"github.com/k3a/html2text"
	"github.com/mmcdole/gofeed"
)

// Function might panic in case of FS write error
func ProcessUrls(
	parser *gofeed.Parser,
	state *State,
	mailer *notifications.Mailer,
	urls []string,
) {
	for _, url := range urls {
		feed, err := parser.ParseURL(url)
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
			// Get all the feeds until we reach the latest GUID
			// Assemble the data for the notification
			newItems := []*gofeed.Item{feed.Items[0]}
			// Only add multiple items if this feed was fetched
			// before. Otherwise we just report the latest item.
			if latestGuid != "" {
				for _, it := range feed.Items[1:] {
					if it.GUID == latestGuid {
						break
					}
					newItems = append(newItems, it)
				}
			}
			// Create the notification:
			notifications := processNotifications(newItems)
			// Create the full body (we could use a template for this)
			// It's beautiful isn't it?
			fullBody := strings.Join(notifications, "\r\n\r\n\r\n====================\r\n\r\n\r\n")
			// Sending email:
			err = mailer.SendNotification(fullBody)
			if err != nil {
				// We don't update the state but go on with our life
				log.Printf("error sending notification email: %v", err)
			} else {
				err := state.SaveLastestGUID(url, feed.Items[0].GUID)
				if err != nil {
					log.Fatal("cannot write state file")
				}
			}
		}
	}
}

// Generates the notification bodies
func processNotifications(items []*gofeed.Item) []string {
	ret := make([]string, len(items))
	for n, it := range items {
		content := fmt.Sprintf(
			"%s\r\n%s\r\n\r\n---\r\n\r\n%v\r\n\r\nLink: %s",
			it.Title,
			it.Published,
			html2text.HTML2Text(it.Description),
			it.Link,
		)
		ret[n] = content
	}
	return ret
}
