// Author: Dr. med. Ramon Saccilotto, Saccilotto Consulting, Switzerland, 2015

// Package jsonwrapper provides some convenience function to work with
// wrapped json data. It follows the conventions outlined on http://www.restapitutorial.com.
// JSON responses are wrapped in a struct containing a code (http-status-code) as integer,
// a status (success, fail, error), an optional message for responses with status fail or error
// and the data containing the actual payload
package jsonwrapper

import "net/http"

const (
	StatusSuccess      string = "success"
	StatusFail         string = "fail"
	StatusError        string = "error"
	StatusUnclassified string = "unclassified"
)

// Response is the struct to contain the wrapped response
type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

// Wrap will encapsulate the given information into a wrapped json response.
// The status attribute will be set automatically depending on the given-status-code
// (400-499: error, 500-599: fail, 100-399: success, other: unclassified).
// Error information can be passed in as data attribute.
// Message should be used for fail and error responses and is omitted from serialization
// if empty.
func Wrap(statusCode int, data interface{}, message string) Response {

	// initialize our status as unclassified
	var status string = StatusUnclassified

	// find the status corresponding to our status-code
	switch {
	case between(statusCode, 100, 399):
		status = StatusSuccess

	case between(statusCode, 400, 499):
		status = StatusError

	case between(statusCode, 500, 599):
		status = StatusFail
	}

	return Response{
		Code:    statusCode,
		Status:  status,
		Message: message,
		Data:    data,
	}

}

// WrapManually will transform the given information into a wrapped response
// without applying any logic
func WrapManually(statusCode int, status string, data interface{}, message string) Response {
	return Response{
		Code:    statusCode,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

// Success will encapsulate the given data as successful response
func Success(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Status:  StatusSuccess,
		Message: "",
		Data:    data,
	}
}

// Error will encapsulate the given information into an error message.
// This kind of response should be used if the request of the client was faulty.
// The status-code will be set to 400 (bad request).
// Please note, that an empty message string will result in a missing message
// attribute when serializing to json.
func Error(message string, err error) Response {
	return Response{
		Code:    http.StatusBadRequest,
		Status:  StatusError,
		Message: message,
		Data:    err,
	}
}

// Fail will encapsulate the given information into a failure message.
// This kind of response should be used when there is a server side error.
// The status code will be set to 500 (internal server error)
// Please note, that an empty message string will result in a missing message
// attribute when serializing to json
func Fail(message string, err error) Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Status:  StatusFail,
		Message: message,
		Data:    err,
	}
}

// between is a utility function to check if the given value
// is between the lower and upper bounds (inclusive)
func between(value int, lower int, upper int) bool {
	return value >= lower && value <= upper
}
