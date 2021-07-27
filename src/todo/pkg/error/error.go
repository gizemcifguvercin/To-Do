package error

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type error interface {
	Error() string
}
type CustomError struct {
	Code int
	Message string
}

type ResponseError struct {
	Error_Message string `json:"error_message,omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d:%s: custom error", e.Code, e.Message)
}

func GetErrorMessage(responseBody io.ReadCloser) ResponseError {
	bodyBytes, err := ioutil.ReadAll(responseBody)
	if err != nil {
		log.Println(err)
	}
	bodyString := string(bodyBytes)
	var response_error ResponseError
	json.Unmarshal([]byte(bodyString), &response_error)

	return response_error;
}
