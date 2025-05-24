package rssemaildigest

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const DefaultSleepInterval uint = 1600

type Config struct {
	Urls          []string
	Email         string
	SleepInterval uint
	SmtpHost      string
	EmailFrom     string
	EmailSubject  string
}

func ConfigFromDotEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// We do not check if the email is valid.
	email := os.Getenv("EMAIL")
	if email == "" {
		return nil, errors.New("missing EMAIL environement variable")
	}

	// Split URLS by space:
	urls := strings.Fields(os.Getenv("URLS"))
	if len(urls) == 0 {
		return nil, errors.New("no feed urls set, use the URLS env variable")
	}

	sleepInterval, _ := strconv.ParseUint(os.Getenv("SLEEP_INTERVAL"), 10, 32)
	sleepInterval32 := uint(sleepInterval)
	if sleepInterval32 == 0 {
		sleepInterval32 = DefaultSleepInterval
	}

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		smtpHost = "localhost"
	}

	emailFrom := os.Getenv("EMAIL_FROM")
	if emailFrom == "" {
		emailFrom = "rss-feed-alerts@localhost.localdomain"
	}

	emailSubject := os.Getenv("EMAIL_SUBJECT")
	if emailSubject == "" {
		emailSubject = "New feed items from watched feeds"
	}

	c := &Config{
		Urls:          urls,
		Email:         email,
		SleepInterval: sleepInterval32,
		SmtpHost:      smtpHost,
		EmailFrom:     emailFrom,
	}

	return c, nil
}
