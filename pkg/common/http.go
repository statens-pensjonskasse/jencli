package common

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func HttpRequest(method string, url string, user string, token string, params ...map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	if params != nil && len(params) > 0 {
		q := req.URL.Query()
		for key, val := range params[0] {
			if strings.HasPrefix(key, "Header ") {
				headerKey := strings.TrimPrefix(key, "Header ")
				req.Header.Set(headerKey, val)
			} else {
				q.Add(key, val)
			}
		}
		req.URL.RawQuery = q.Encode()
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, token)))
	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func PostRequest(url string, user string, token string, params map[string]string) (*http.Response, error) {
	return HttpRequest("POST", url, user, token, params)
}
