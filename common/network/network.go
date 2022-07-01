package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//to do....
type HTTP_METHOD string
type HTTP_REPONSE_CALLBACK func(response *http.Response, reftype interface{}) error

//...................................................................
const (
	GET    HTTP_METHOD = "GET"
	POST   HTTP_METHOD = "POST"
	PUT    HTTP_METHOD = "PUT"
	DELETE HTTP_METHOD = "DELETE"
)

//...................................................................
type Header struct {
	Key   string
	Value string
}

//...................................................................
type Request struct {
	Method   HTTP_METHOD
	Url      string
	Headers  []Header
	Body     interface{}
	Callback HTTP_REPONSE_CALLBACK
}

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
func PrepareHeaders(options *Request, req *http.Request) {
	for _, header := range options.Headers {
		req.Header.Set(header.Key, header.Value)
	}
}

//========================================================================================================================

func Send(options Request, reftype interface{}) error {

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
	error := options.Callback(resp, reftype)
	if error != nil {
		return error
	}
	return nil

}
