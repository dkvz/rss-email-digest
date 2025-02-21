package config

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

	c := &Config{
		Urls:          urls,
		Email:         email,
		SleepInterval: sleepInterval32,
		SmtpHost:      smtpHost,
	}

	return c, nil
}
