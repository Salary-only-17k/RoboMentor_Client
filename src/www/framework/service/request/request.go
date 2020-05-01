package RequestService

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type ResponseData struct {
	Code 	int 			`json:"code"`
	Data 	interface{} 	`json:"data"`
	Msg 	string 			`json:"msg"`
}

func Get(url string, params map[string]string, headers map[string]string, AppID string, AppSecret string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("new request is fail")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Robot-Token", AppID + "@" + AppSecret + "@" + strconv.Itoa(int(time.Now().Unix())))

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

func Post(url string, body map[string]string, params map[string]string, headers map[string]string, AppID string, AppSecret string) (*http.Response, error) {

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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Robot-Token", AppID + "@" + AppSecret + "@" + strconv.Itoa(int(time.Now().Unix())))

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