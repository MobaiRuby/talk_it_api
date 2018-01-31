package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpConPool struct {
	Conn *http.Client
}

var Hpool *HttpConPool

func init() {
	Hpool = new(HttpConPool)
	Hpool.Conn = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
		Timeout: time.Duration(1000 * time.Millisecond),
	}
}

func (h *HttpConPool) Request(url string, method string, data string, header map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	response, err := h.Conn.Do(req)

	if err != nil {
		return nil, err
	} else if response != nil {
		defer response.Body.Close()

		r_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			return string(r_body), nil
		}
	} else {
		return nil, nil
	}
}
