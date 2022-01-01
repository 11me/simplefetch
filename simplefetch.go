package simplefetch

import (
	"net/http"
	"net/url"
)

type Options struct {
	URL     string
	Method  string
	Headers http.Header
}

func Get(o Options) (*http.Response, error) {
	url, err := url.Parse(o.URL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	var headers http.Header
	if o.Headers == nil {
		headers = map[string][]string{
			"User-Agent": {"simple-fetch"},
		}
	}

	req := &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
