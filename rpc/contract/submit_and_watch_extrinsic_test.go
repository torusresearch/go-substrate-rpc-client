package contract_test

import (
	"math/big"
	"testing"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/rpc/author"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/assert"
)

func TestAuthor_SubmitAndWatchExtrinsic(t *testing.T) {
	api, err := gsrpc.NewSubstrateAPI("ws://localhost:9944")
	assert.NoError(t, err)

	meta, err := api.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	contractAddress := "5GtgbPRMtcEBAVSjcWqYjWBaz6h8Au9CBJa1WBxWrgG4eBRz"
	dest := types.NewMultiAddressFromAddress(contractAddress)
	amount := types.NewUCompact(big.NewInt(0))
	gasLimit := types.NewUCompact(big.NewInt(1280000000000))
	storageDepositLimit := types.NewOptionU128Empty()
	data := types.MustHexDecodeString("0x633aa551")

	c, err := types.NewCall(meta, "Contracts.call", dest, amount, gasLimit, storageDepositLimit, data)
	assert.NoError(t, err)

	var sub *author.ExtrinsicStatusSubscription
	for {
		ext := types.NewExtrinsic(c)
		genesisHash, err := api.RPC.Chain.GetBlockHash(0)
		assert.NoError(t, err)

		rv, err := api.RPC.State.GetRuntimeVersionLatest()
		assert.NoError(t, err)

		// Get the nonce for Alice
		key, err := types.CreateStorageKey(meta, "System", "Account", signature.TestKeyringPairAlice.PublicKey)
		assert.NoError(t, err)

		var accountInfo types.AccountInfo
		ok, err := api.RPC.State.GetStorageLatest(key, &accountInfo)
		assert.NoError(t, err)
		assert.True(t, ok)
		nonce := uint32(accountInfo.Nonce)
		o := types.SignatureOptions{
			BlockHash:          genesisHash,
			Era:                types.ExtrinsicEra{IsMortalEra: false},
			GenesisHash:        genesisHash,
			Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
			SpecVersion:        rv.SpecVersion,
			Tip:                types.NewUCompactFromUInt(0),
			TransactionVersion: rv.TransactionVersion,
		}

		// Sign the transaction using Alice's default account
		err = ext.Sign(signature.TestKeyringPairAlice, o)
		assert.NoError(t, err)

		// Do the transfer and track the actual status
		sub, err = api.RPC.Author.SubmitAndWatchExtrinsic(ext)
		if err != nil {
			t.Logf("extrinsic submit failed: %v", err)
			continue
		}

		break
	}
	defer sub.Unsubscribe()
	for {
		status := <-sub.Chan()

		// wait until finalisation
		if status.IsInBlock || status.IsFinalized {
			break
		}

		t.Log("waiting for the extrinsic to be included/finalized")
	}

}