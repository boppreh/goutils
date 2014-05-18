package goutils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0666)
}

func ReadUrl(url string) ([]byte, error) {
	if !strings.Contains(url, "://") {
		url = "http://" + url
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, errors.New("Server returned error " + response.Status)
	}

	content, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return nil, err
	}

	return content, nil
}

func Read(path string) ([]byte, error) {
	if _, err := os.Stat(path); err == nil {
		return ReadFile(path)
	} else {
		return ReadUrl(path)
	}
}

func SearchContent(contentBytes []byte, stringPattern string) [][]string {
	pattern := regexp.MustCompile(stringPattern)
	content := string(contentBytes)
	return pattern.FindAllStringSubmatch(content, -1)
}

func Search(path string, stringPattern string) ([][]string, error) {
	contentBytes, err := Read(path)
	if err != nil {
		return nil, err
	}

	return SearchContent(contentBytes, stringPattern), nil
}

func SearchFile(path string, stringPattern string) ([][]string, error) {
	contentBytes, err := ReadFile(path)
	if err != nil {
		return nil, err
	}

	return SearchContent(contentBytes, stringPattern), nil
}

func SearchUrl(path string, stringPattern string) ([][]string, error) {
	contentBytes, err := ReadUrl(path)
	if err != nil {
		return nil, err
	}

	return SearchContent(contentBytes, stringPattern), nil
}
