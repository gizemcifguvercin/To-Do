package http_client

import (
	"bytes"
	"encoding/json"
	error_pkg "github.com/gizemcifguvercin/To-Do/src/todo/pkg/error"
	"io/ioutil"
	"net/http"
)

type ClientBase struct {
	BaseUrl string
	Client  *http.Client
	Retries int
}

type Client interface {
	ConvertBodyToString(response *http.Response) (bodyString string, err error)
	GetWith(endpoint string) (*http.Response, error)
	DeleteWith(endpoint string) (*http.Response, error)
	PostWith(endpoint string, params interface{}) (*http.Response, error)
}

func New(baseUrl string) Client {
	return &ClientBase{BaseUrl: baseUrl, Client: &http.Client{}, Retries: 3}
}

func (h ClientBase) Send(req *http.Request) (resp *http.Response, err error) {
	defaultRetryCount := h.Retries

	for defaultRetryCount > 0 {
		resp, err = h.Client.Do(req)
		if err != nil {
			defaultRetryCount -= 1
		} else {
			break
		}
	}
	return CreateResponse(resp, err)
}

func (h ClientBase) ConvertBodyToString(response *http.Response) (bodyString string, err error) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", &error_pkg.CustomError{Code: response.StatusCode, Message: err.Error()}
	}
	bodyString = string(bodyBytes)
	return bodyString, nil
}

func CreateResponse(resp *http.Response, error error) (*http.Response, error) {
	if error != nil {
		if resp == nil || resp.Body == nil {
			return nil, &error_pkg.CustomError{Code: 500, Message: error.Error()}
		}

		response_error := error_pkg.GetErrorMessage(resp.Body)
		return nil, &error_pkg.CustomError{Code: resp.StatusCode, Message: response_error.Error_Message}
	}

	error = CheckHttpStatus(resp)
	if error != nil {
		return nil, error
	}

	return resp, nil
}

func CheckHttpStatus(resp *http.Response) error {
	if !(resp.StatusCode >= 200 && resp.StatusCode <= 299) {
		response_error := error_pkg.GetErrorMessage(resp.Body)
		return &error_pkg.CustomError{Code: resp.StatusCode, Message: response_error.Error_Message}
	}
	return nil
}

func (h ClientBase) GetWith(endpoint string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	return h.Send(req)
}

func (h ClientBase) DeleteWith(endpoint string) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	return h.Send(req)
}

func (h ClientBase) PostWith(endpoint string, params interface{}) (resp *http.Response, err error) {
	json, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(json))
	return h.Send(req)
}
