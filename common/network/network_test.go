package common

import (
	"fmt"
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
	var options Request
	var body Body
	body.id = 78912
	body.Customer = "ason Sweet"
	body.Quantity = 2
	body.Price = 18.08
	options.Method = POST
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
