package common

import (
	"fmt"
	baseRequest "main/common/network/requests"
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
	var options baseRequest.Request
	options.Method = baseRequest.GET
	options.Url = "https://catfact.ninja/fact"
	var res ResPonseType
	options.Callback = func(response *http.Response) error {

		err := ParsResponse(response, &res)

		assert.NoError(t, err)
		fmt.Println(res)
		return err

	}
	error := Send(options)

	assert.NoError(t, error)
}

//========================================================================================================================
func TestSendFunctionalityPostMethod(t *testing.T) {
	type Body struct {
		id       int
		Customer string
		Quantity int
		Price    float64
	}
	type CustomerReponse struct {
		Success bool `json:"success"`
	}
	var options baseRequest.Request
	var body Body
	body.id = 78912
	body.Customer = "ason Sweet"
	body.Quantity = 2
	body.Price = 18.08
	options.Method = baseRequest.POST
	options.Body = body
	options.Url = "https://reqbin.com/echo/post/json"
	var res CustomerReponse
	options.Callback = func(response *http.Response) error {
		err := ParsResponse(response, &res)

		assert.NoError(t, err)
		fmt.Println(res)
		return err
	}
}
