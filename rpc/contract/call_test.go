package contract_test

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/contract"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract_Call(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI("ws://localhost:9944")
	assert.NoError(t, err)

	var res interface{}
	alice := signature.TestKeyringPairAlice

	request := contract.CallRequest{
		Origin:              alice.Address,
		Dest:                "5CfVtCYLzTiqsiXZE4p3exSEMQH1bxrnFGvon2FXb55e2uzi",
		Value:               0,
		GasLimit:            5000000000000,
		StorageDepositLimit: nil,
		InputData:           "0x2f865bd9",
	}

	err = api.RPC.Contract.Call(&res, request)
	assert.NoError(t, err)

	log.Info("res: ", res)
}
