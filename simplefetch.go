package simplefetch

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

var headers http.Header
var client http.Client

type Data map[string]interface{}

type Options struct {
	URL     string
	Method  string
	Headers http.Header
	Data    Data
}

type bytesReadCloser struct {
	io.Reader
}

func (brc bytesReadCloser) Close() error { return nil }

func Get(o Options) (*http.Response, error) {
	url, err := url.Parse(o.URL)
	if err != nil {
		return nil, err
	}

	client = http.Client{}

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

func Post(o Options) (*http.Response, error) {
	client = http.Client{}

	url, err := url.Parse(o.URL)
	if err != nil {
		return nil, err
	}

	if o.Headers == nil {
		headers = map[string][]string{
			"User-Agent": {"simple-fetch"},
		}
	}

	json_data, err := json.Marshal(o.Data)
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		URL:    url,
		Method: http.MethodPost,
		Header: headers,
		Body:   bytesReadCloser{bytes.NewBuffer(json_data)},
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
