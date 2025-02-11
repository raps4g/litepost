package core

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func SendHttpRequest(reqData Request) (*http.Response, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	var body io.Reader
	if reqData.Body != "" {
		body = bytes.NewBufferString(reqData.Body)
	}

	req, err := http.NewRequest(reqData.Methods[reqData.SelectedMethod], reqData.Url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range reqData.Headers {
        if (key != "" && value != "") {
            canonKey := http.CanonicalHeaderKey(key)
            req.Header.Set(canonKey, value)
        }
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
