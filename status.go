package rssemaildigest

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

const statusFilename = "./rss-email-digest.status"

// Map key is going to be the RSS URL
// Value is the latest GUID
type State struct {
	latestGUIDs map[string]string
}

func ReadState() (error, *State) {
	// Attempt to read it from disk:
	content, err := os.ReadFile(statusFilename)
	if err != nil || len(content) == 0 {
		return nil, &State{
			latestGUIDs: make(map[string]string),
		}
	}

	var result map[string]string
	err = json.Unmarshal(content, &result)
	if err != nil {
		return errors.New("invalid state file (parsing failed)"), nil
	}

	return nil, &State{
		latestGUIDs: result,
	}
}

func (s *State) SaveLastestGUID(url string, guid string) error {
	if guid == "" {
		// Not supposed to happen.
		guid = time.Now().String()
	}
	s.latestGUIDs[url] = guid
	jsonData, err := json.Marshal(s.latestGUIDs)
	if err != nil {
		panic("can't convert state to JSON")
	}
	// First save it to hard drive:
	err = os.WriteFile(statusFilename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *State) IsNewGUID(url string, guid string) bool {
	if guid == "" {
		return false
	}
	if latestGuid, ok := s.latestGUIDs[url]; ok {
		return latestGuid != guid
	}
	return false
}

func (s *State) LatestGUID(url string) string {
	return s.latestGUIDs[url]
}
