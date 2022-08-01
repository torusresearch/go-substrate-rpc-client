package contract

import (
	"github.com/mitchellh/mapstructure"
	_ "github.com/mitchellh/mapstructure"
)

const method = "contracts_call"

func (c *contract) Call(callRequest CallRequest) (*CallResponse, error) {
	var res interface{}
	err := c.client.Call(&res, method, callRequest)
	if err != nil {
		return nil, err
	}

	var callResponse CallResponse
	err = mapstructure.Decode(res, &callResponse)
	if err != nil {
		return nil, err
	}

	return &callResponse, nil
}
