package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func makeRequest(url, ua string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)
	return req, nil
}

func getResponse(req *http.Request, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func handleResponse(resp *http.Response) (string, error) {
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid http status code %d", resp.StatusCode)
	}
	if b, err := ioutil.ReadAll(resp.Body); err == nil {
		return string(b), nil
	}
	defer resp.Body.Close()
	return "", fmt.Errorf("no body read")
}

func Fetch(url, ua string, timeout time.Duration) (string, error) {
	request, err := makeRequest(url, ua)
	if err != nil {
		return "", err
	}
	response, err := getResponse(request, timeout)
	if err != nil {
		return "", err
	}
	body, err := handleResponse(response)
	if err != nil {
		return "", err
	}
	return body, nil
}
