package common

import "net/http"

//to do....
type HTTP_METHOD string
type HTTP_REPONSE_CALLBACK func(response *http.Response) error

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
