package jsonwrapper

import (
	"errors"
	"reflect"
	"testing"
)

func TestConstants(t *testing.T) {

	if StatusUnclassified != "unclassified" {
		t.Error("Expected status unclassified to be «unclassified», was", StatusUnclassified)
	}

	if StatusSuccess != "success" {
		t.Error("Expected status success to be «success», was", StatusSuccess)
	}

	if StatusError != "error" {
		t.Error("Expected status error to be «error», was", StatusError)
	}

	if StatusFail != "fail" {
		t.Error("Expected status fail to be «fail», was", StatusFail)
	}

}

func TestWrap(t *testing.T) {

	// --- status unclassified

	statusCode := 700
	data := []string{"test one", "test two", "test three", "test four"}
	message := "a custom message"

	response := Wrap(statusCode, data, message)

	if response.Code != 700 {
		t.Error("Expected status code to be 700, was", response.Code)
	}

	if response.Status != StatusUnclassified {
		t.Error("Expected the status to be unclassified, was", response.Status)
	}

	if response.Message != message {
		t.Error("Expected the message to equal our custom message")
	}

	if reflect.DeepEqual(response.Data, data) != true {
		t.Error("Expeted the data to contain our data")
	}

	// --- status success ---

	statusCode = 222
	response = Wrap(statusCode, data, message)

	if response.Code != 222 {
		t.Error("Expected status code to be 222, was", response.Code)
	}

	if response.Status != StatusSuccess {
		t.Error("Expected the status to be success, was", response.Status)
	}

	// --- status fail ---

	statusCode = 555
	response = Wrap(statusCode, data, message)

	if response.Code != 555 {
		t.Error("Expected status code to be 555, was", response.Code)
	}

	if response.Status != StatusFail {
		t.Error("Expected the status to be fail, was", response.Status)
	}

	// --- status error ---

	statusCode = 444
	response = Wrap(statusCode, data, message)

	if response.Code != 444 {
		t.Error("Expected status code to be 444, was", response.Code)
	}

	if response.Status != StatusError {
		t.Error("Expected the status to be error, was", response.Status)
	}

}

func TestWrapManually(t *testing.T) {

	statusCode := 700
	status := "undefined status"
	data := []string{"test one", "test two", "test three", "test four"}
	message := "a custom message"

	response := WrapManually(statusCode, status, data, message)

	if response.Code != 700 {
		t.Error("Expected status code to be 700")
	}

	if response.Status != status {
		t.Error("Expected the status to equal our custom status")
	}

	if response.Message != message {
		t.Error("Expected the message to equal our custom message")
	}

	if reflect.DeepEqual(response.Data, data) != true {
		t.Error("Expeted the data to contain our data")
	}

}

func TestSuccess(t *testing.T) {

	data := []string{"test one", "test two", "test three", "test four"}

	response := Success(data)

	if response.Code != 200 {
		t.Error("Expected status code to be 200")
	}

	if response.Status != StatusSuccess {
		t.Error("Expected the status to be success")
	}

	if response.Message != "" {
		t.Error("Expected the message to be empty")
	}

	if reflect.DeepEqual(response.Data, data) != true {
		t.Error("Expeted the data to contain our data")
	}
}

func TestError(t *testing.T) {

	var message = "The request was faulty"
	var errorData = "This is a error"

	err := errors.New(errorData)
	response := Error(message, err)

	if response.Code != 400 {
		t.Error("Expected status code to be 400")
	}

	if response.Status != StatusError {
		t.Error("Expected the status to be error")
	}

	if response.Message != message {
		t.Error("Expected the message to equal our message")
	}

	if response.Data != err {
		t.Error("Expeted the data to contain the error-information")
	}

}

func TestFail(t *testing.T) {

	var message = "The request failed"
	var errorData = "This is a fail error"

	err := errors.New(errorData)
	response := Fail(message, err)

	if response.Code != 500 {
		t.Error("Expected status code to be 500")
	}

	if response.Status != StatusFail {
		t.Error("Expected the status to be fail")
	}

	if response.Message != message {
		t.Error("Expected the message to equal our message")
	}

	if response.Data != err {
		t.Error("Expeted the data to contain the error-information")
	}

}

func TestBetween(t *testing.T) {

	// test if our between function is returning the expected results
	if between(100, 50, 200) != true {
		t.Error("Expected the result of between() to be true")
	}

	if between(100, 100, 101) != true {
		t.Error("Expected the result of between() to be true")
	}

	if between(101, 100, 101) != true {
		t.Error("Expected the result of between() to be true")
	}

	if between(10, 100, 101) != false {
		t.Error("Expected the result of between() to be false")
	}
}
