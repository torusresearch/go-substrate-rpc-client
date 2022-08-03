package contract_test

import (
	"github.com/btcsuite/btcutil/base58"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/contract"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContract_Call(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI("ws://localhost:9944")
	assert.NoError(t, err)

	alice := signature.TestKeyringPairAlice
	contractAddr := "5Fm4n8sWxgtbSBd2aKXHanccNuL6vqKuJwuVXn91cDby7JG5"

	data := types.MustHexDecodeString("0xe303952a")
	ss58d := base58.Decode(contractAddr)
	data = append(data, ss58d...)
	hexData, err := types.Hex(data)
	assert.NoError(t, err)

	request := contract.CallRequest{
		Origin:              alice.Address,
		Dest:                contractAddr,
		Value:               0,
		GasLimit:            5000000000000,
		StorageDepositLimit: nil,
		InputData:           hexData,
	}

	res, err := api.RPC.Contract.Call(request)
	assert.NoError(t, err)
	assert.Nil(t, res.Result.Err)
	assert.NotNil(t, res.Result.Success)
}
