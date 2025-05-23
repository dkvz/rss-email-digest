package rssemaildigest

import (
	"os"
	"strings"
	"time"
)

const statusFilename = "./rss-email-digest.status"

type State struct {
	latestGUID string
}

func (s *State) SaveLastestGUID(guid string) error {
	if guid == "" {
		// Not supposed to happen.
		guid = time.Now().String()
	}
	// First save it to hard drive:
	err := os.WriteFile(statusFilename, []byte(guid), 0644)
	if err != nil {
		return err
	}
	s.latestGUID = guid
	return nil
}

func (s *State) IsNewGUID(guid string) bool {
	if guid == "" {
		return false
	}
	// If latestGUID is an empty string, load from disk
	if s.latestGUID == "" {
		// We ignore read errors, program will crash on
		// write errors.
		content, _ := os.ReadFile(statusFilename)
		// Chop possible line returns and extra space:
		s.latestGUID = strings.TrimRight(string(content), " \n\r\t")
	}
	return s.latestGUID != guid
}
