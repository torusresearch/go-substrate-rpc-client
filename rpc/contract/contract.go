package contract

import "github.com/centrifuge/go-substrate-rpc-client/v4/client"

type Contract interface {
	Call(result interface{}, callRequest CallRequest) error
}

// contract exposes methods for contract rpc calls
type contract struct {
	client client.Client
}

// NewContract creates a new contract struct
func NewContract(cl client.Client) Contract {
	return &contract{cl}
}
