package main

import (
	"encoding/json"
)

func GetBuddyVersion(address string, apiKey string, username string, password string) buddyVersion {
	resp := accessBuddyApi("version", address, apiKey, username, password)

	var result buddyVersion

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetBuddyFiles(address string, apiKey string, username string, password string) buddyFiles {
	resp := accessBuddyApi("files", address, apiKey, username, password)

	var result buddyFiles

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetBuddyJob(address string, apiKey string, username string, password string) buddyJob {
	resp := accessBuddyApi("job", address, apiKey, username, password)

	var result buddyJob

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetBuddyPrinter(address string, apiKey string, username string, password string) buddyPrinter {
	resp := accessBuddyApi("printer", address, apiKey, username, password)

	var result buddyPrinter

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}
