package jsonwrapper

import "encoding/json"

// ToJsonBytes will marshal the response struct to a byte-slice
func ToJsonBytes(response *Response) ([]byte, error) {
	return json.Marshal(response)
}

// ToJsonString will marshal the response struct to a json string
func ToJsonString(response *Response) (string, error) {
	data, err := json.Marshal(response)
	if err != nil {
		return "", err
	}

	return string(data[:]), nil
}

// FromBytes will unmarshal a json given as byte-slice into a response struct
func FromBytes(jsonData []byte) (*Response, error) {
	response := &Response{}
	err := json.Unmarshal(jsonData, response)
	return response, err
}

// FromString will unmarshal a json given as string into a response struct
func FromString(jsonData string) (*Response, error) {
	return FromBytes([]byte(jsonData))
}
