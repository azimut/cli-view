package fetch

import (
	"fmt"
	"net/http"
	"time"
)

const MAX_REDIRECTS = 15

func UrlLocation(url, ua string, timeout time.Duration) (string, error) {
	locations, err := urlLocations(url, ua, timeout)
	if err != nil {
		return "", err
	}
	return locations[len(locations)-1], err
}

func urlLocations(url, ua string, timeout time.Duration) (locations []string, err error) {

	locations = append(locations, url)

	var counter int
	client := &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if counter > MAX_REDIRECTS {
				return fmt.Errorf("Too many `%d` redirects!", counter)
			}
			location, err := req.Response.Location()
			if err != nil {
				return err
			}
			locations = append(locations, location.String())
			counter = counter + 1
			return nil
		},
	}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", ua)

	_, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	return
}
