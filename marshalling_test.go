package jsonwrapper

import (
	"bytes"
	"fmt"
	"testing"
)

func TestToJsonBytes(t *testing.T) {
	data := []string{"test", "data"}
	expectedJson := `{"code":200,"status":"success","data":["test","data"]}`
	response := Success(data)

	asJson, err := ToJsonBytes(&response)

	if err != nil {
		t.Error("Expected error to be nil, was", err)
	}

	if bytes.Equal(asJson, []byte(expectedJson)) != true {
		t.Error("Expected byte format did not match")
	}
}

func TestToJsonString(t *testing.T) {
	data := []string{"test", "data"}
	expectedJson := `{"code":200,"status":"success","data":["test","data"]}`
	response := Success(data)

	asJson, err := ToJsonString(&response)

	if err != nil {
		t.Error("Expected error to be nil, was", err)
	}

	if asJson != expectedJson {
		t.Error("Expected string format did not match")
	}
}

func TestFromBytes(t *testing.T) {
	data := []string{"test", "data"}
	jsonData := `{"code":200,"status":"success","data":["test","data"]}`

	response, err := FromBytes([]byte(jsonData))

	if err != nil {
		t.Error("Expected error to be nil, was", err)
	}

	if response.Code != 200 {
		t.Error("Expected code to be 200")
	}

	if response.Status != StatusSuccess {
		t.Error("Expected status to be success")
	}

	if response.Message != "" {
		t.Error("Expected message to be empty")
	}

	if fmt.Sprint(response.Data) != fmt.Sprint(data) {
		t.Error("Expected the data to match our definitions, was", response.Data, data)
	}
}

func TestString(t *testing.T) {
	data := []string{"test", "data"}
	jsonData := `{"code":200,"status":"success","data":["test","data"]}`

	response, err := FromString(jsonData)

	if err != nil {
		t.Error("Expected error to be nil, was", err)
	}

	if response.Code != 200 {
		t.Error("Expected code to be 200")
	}

	if response.Status != StatusSuccess {
		t.Error("Expected status to be success")
	}

	if response.Message != "" {
		t.Error("Expected message to be empty")
	}

	if fmt.Sprint(response.Data) != fmt.Sprint(data) {
		t.Error("Expected the data to match our definitions, was", response.Data, data)
	}
}
