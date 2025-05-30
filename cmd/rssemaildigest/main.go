package main

import (
	"log"
	"time"

	rssemaildigest "github.com/dkvz/rss-email-digest"
	"github.com/dkvz/rss-email-digest/notifications"
	"github.com/mmcdole/gofeed"
)

func main() {
	// Comment the following line to have log output
	// timestamps with each line. Unwanted for docker
	// and systemd.
	log.SetFlags(0)

	conf, err := rssemaildigest.ConfigFromDotEnv()
	if err != nil {
		log.Fatal("Could not load configuration: " + err.Error())
	}

	// Create the Mailer:
	mailer := notifications.NewMailer(
		conf.SmtpHost,
		conf.EmailFrom,
		conf.Email,
		conf.EmailSubject,
	)

	// Create the state:
	state, err := rssemaildigest.ReadState()
	if err != nil {
		log.Fatal("Could not parse state file, check format or remove it")
	}

	fp := gofeed.NewParser()
	for {
		rssemaildigest.ProcessUrls(fp, state, mailer, conf.Urls)
		time.Sleep(time.Duration(conf.SleepInterval) * time.Second)
	}

}
