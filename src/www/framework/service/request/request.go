package RequestService

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type ResponseData struct {
	Code 	int 			`json:"code"`
	Data 	interface{} 	`json:"data"`
	Msg 	string 			`json:"msg"`
}

func DefaultGet(url string, params map[string]string) (*http.Response, error) {
	return Get(url, params, nil)
}

func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("new request is fail ")
	}

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}

	return client.Do(req)
}

func DefaultPost(url string, body map[string]string) (*http.Response, error) {
	return Post(url, body, nil, nil)
}

func Post(url string, body map[string]string, params map[string]string, headers map[string]string) (*http.Response, error) {

	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			return nil, errors.New("http post body to json failed")
		}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, errors.New("new request is fail: %v \n")
	}

	req.Header.Set("Content-type", "application/json")

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}

	return client.Do(req)
}