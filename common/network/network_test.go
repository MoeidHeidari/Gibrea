package network

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResPonseType struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

//To do...
func TestSendFunctionalityGetMethod(t *testing.T) {
	var options Request
	options.Method = GET
	options.Url = "https://catfact.ninja/fact"
	var response ResPonseType
	options.Callback = func(response *http.Response, reftype interface{}) error {

		var resp ResPonseType
		err := ParsResponse(response, &resp)
		assert.NoError(t, err)

		return err

	}
	error := Send(options, response)

	assert.NoError(t, error)
}

//========================================================================================================================
func TestSendFunctionalityPostMethod(t *testing.T) {

}
