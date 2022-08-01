package contract_test

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/contract"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract_Call(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI("ws://localhost:9944")
	assert.NoError(t, err)

	alice := signature.TestKeyringPairAlice
	contractAddr := "5GCAvjHNhyZFoztHvPj2PxDnwVP2Zo8gs66Lgy7Gr73ar1Zp"

	request := contract.CallRequest{
		Origin:              alice.Address,
		Dest:                contractAddr,
		Value:               0,
		GasLimit:            5000000000000,
		StorageDepositLimit: nil,
		InputData:           "0x633aa551",
	}

	res, err := api.RPC.Contract.Call(request)
	assert.NoError(t, err)
	assert.Nil(t, res.Result.Err)
	assert.NotNil(t, res.Result.Success)
}
