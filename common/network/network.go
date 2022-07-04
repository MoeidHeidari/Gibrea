package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	baseRequest "main/common/network/requests"
	"net/http"
)

//========================================================================================================================
func ParsResponseBody(body []byte, reftype *interface{}) interface{} {
	stringBuffer := string(body)

	err := json.Unmarshal([]byte(stringBuffer), &reftype)
	if err != nil {
		return err
	}

	return &reftype
}

//========================================================================================================================
func ParsResponse(reposnse *http.Response, reftype interface{}) error {
	body, err := ioutil.ReadAll(reposnse.Body)
	if err != nil {
		return err
	}
	ParsResponseBody(body, &reftype)
	return nil
}

//========================================================================================================================
func PrepareBody(data interface{}) (*bytes.Reader, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req_body := bytes.NewReader(payloadBytes)
	return req_body, nil
}

//========================================================================================================================
func PrepareHeaders(options *baseRequest.Request, req *http.Request) {
	for _, header := range options.Headers {
		req.Header.Set(header.Key, header.Value)

	}
}

//========================================================================================================================

func Send(options baseRequest.Request) error {

	data, reqBodyError := PrepareBody(options.Body)
	if reqBodyError != nil {
		return reqBodyError
	}

	req, err := http.NewRequest(string(options.Method), options.Url, data)
	if err != nil {
		return err
	}
	PrepareHeaders(&options, req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if options.Callback != nil {
		error := options.Callback(resp)
		if error != nil {
			return error
		}
	}

	return nil

}
