package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/icholy/digest"
)

func head(s string) bool {
	r, e := http.Head(s)
	return e == nil && r.StatusCode == 200
}

func connTest(s string) bool {
	r, e := http.Get(s)
	return e == nil && r.StatusCode == 200
}

func getURL(path string, address string) string {
	return fmt.Sprintf("http://%s/api/%s", address, path) // http is used in internal network only
}

func accessBuddyApi(path string, address string, apiKey string, username string, password string) []byte {
	url := getURL(path, address)
	var res *http.Response
	var err error
	var body []byte
	if apiKey == "" {
		client := &http.Client{
			Transport: &digest.Transport{
				Username: username,
				Password: password,
			},
		}
		res, err = client.Get(url)
		if err != nil {
			logger.Error(err.Error())
		}
	} else {
		req, _ := http.NewRequest("GET", url, nil)
		client := &http.Client{}
		req.Header.Add("X-Api-Key", apiKey)
		res, err = client.Do(req)
		if err != nil {
			logger.Error(err.Error())
		}
	}
	if err == nil {
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	} else {
		logger.Error(err.Error())
	}

	return body
}

func accessEinsyApi(path string, address string, apiKey string) ([]byte, error) {
	url := getURL(path, address)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Api-Key", apiKey)
	client := &http.Client{Timeout: time.Duration(scrapeTimeout) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	} else {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger.Error(err.Error())
		}
		return body, nil
	}
}
