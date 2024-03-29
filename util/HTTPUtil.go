package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPResult struct {
	StatusCode      int
	StatusMessage   string
	ResponseContent string
	Url             string
	PostData        string
	HttpMethod      string
}

func HTTPGet(url string, timeOutInt int64) (response HTTPResult, err error) {
	response.Url = url
	response.HttpMethod = http.MethodGet
	response.PostData = ""

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return response, err
	}
	var timeOut time.Duration = time.Duration(timeOutInt) * time.Second
	http.DefaultClient.Timeout = timeOut

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return response, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return response, err
	}

	response.StatusMessage = ""
	response.StatusCode = res.StatusCode
	response.ResponseContent = string(resBody)
	return response, nil
}

func HTTPPost(url string, postData string, timeOutInt int64) (response HTTPResult, err error) {
	response.Url = url
	response.HttpMethod = http.MethodPost
	response.PostData = postData

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return response, err
	}
	var timeOut time.Duration = time.Duration(timeOutInt) * time.Second
	http.DefaultClient.Timeout = timeOut

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return response, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return response, err
	}

	response.StatusMessage = ""
	response.StatusCode = res.StatusCode
	response.ResponseContent = string(resBody)
	return response, nil
}
