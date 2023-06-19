package main

import (
	"encoding/json"
)

func GetEinsyJob(address string, apiKey string) einsyJob {
	resp, _ := accessEinsyApi("job", address, apiKey)

	var result einsyJob

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyCameras(address string, apiKey string) einsyCameras {
	resp, _ := accessEinsyApi("v1/cameras", address, apiKey)

	var result einsyCameras

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyPrinter(address string, apiKey string) einsyPrinter {
	resp, _ := accessEinsyApi("printer", address, apiKey)

	var result einsyPrinter

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyStorage(address string, apiKey string) einsyStorage { // currently unused
	resp, _ := accessEinsyApi("v1/storage", address, apiKey)

	var result einsyStorage

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsySettings(address string, apiKey string) einsySettings {
	resp, _ := accessEinsyApi("settings", address, apiKey)

	var result einsySettings

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyConnection(address string, apiKey string) einsyConection { // currently unused
	resp, _ := accessEinsyApi("connection", address, apiKey)

	var result einsyConection

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyFiles(address string, apiKey string) einsyFiles {
	resp, _ := accessEinsyApi("files", address, apiKey)

	var result einsyFiles

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyLogs(address string, apiKey string) einsyLogs {
	resp, _ := accessEinsyApi("logs", address, apiKey)

	var result einsyLogs

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyInfo(address string, apiKey string) einsyInfo {
	resp, _ := accessEinsyApi("v1/info", address, apiKey)

	var result einsyInfo

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyStatus(address string, apiKey string) einsyStatus { // currently unused
	resp, _ := accessEinsyApi("v1/status", address, apiKey)

	var result einsyStatus

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyVersion(address string, apiKey string) einsyVersion {
	resp, _ := accessEinsyApi("version", address, apiKey)

	var result einsyVersion

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}

func GetEinsyPorts(address string, apiKey string) einsyPorts {
	resp, _ := accessEinsyApi("ports", address, apiKey)

	var result einsyPorts

	if err := json.Unmarshal(resp, &result); err != nil {
		logger.Error("Can not unmarshal JSON")
	}

	return result
}
