// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types_test

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"testing"

	"github.com/stretchr/testify/assert"

	fuzz "github.com/google/gofuzz"

	. "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

var (
	optionAccountIDFuzzOpts = []fuzzOpt{
		withFuzzFuncs(func(o *OptionAccountID, c fuzz.Continue) {
			if c.RandBool() {
				*o = NewOptionAccountIDEmpty()
				return
			}
			var accId AccountID
			c.Fuzz(&accId)

			*o = NewOptionAccountID(accId)
		}),
	}
)

func TestOptionAccountID_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[OptionAccountID](t, 100, optionAccountIDFuzzOpts...)
	assertEncodeEmptyObj[OptionAccountID](t, 1)
}

func TestOptionAccountID_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{NewOptionAccountID(NewAccountID([]byte{171, 18, 52})), MustHexDecodeString("0x01ab12340000000000000000000000000000000000000000000000000000000000")},
		{NewOptionAccountIDEmpty(), MustHexDecodeString("0x00")},
	})
}

func TestOptionAccountID_Decode(t *testing.T) {
	assertDecode(t, []decodingAssert{
		{MustHexDecodeString("0x01ab12340000000000000000000000000000000000000000000000000000000000"), NewOptionAccountID(NewAccountID([]byte{171, 18, 52}))},
		{MustHexDecodeString("0x00"), NewOptionAccountIDEmpty()},
	})
}

func TestOptionAccountID_OptionMethods(t *testing.T) {
	o := NewOptionAccountIDEmpty()
	o.SetSome(NewAccountID([]byte("acc-id")))

	ok, v := o.Unwrap()
	assert.True(t, ok)
	assert.NotNil(t, v)

	o.SetNone()

	ok, v = o.Unwrap()
	assert.False(t, ok)
	assert.Equal(t, NewAccountID([]byte{}), v)
}

func TestAccountID_EncodeDecode(t *testing.T) {
	assertRoundTripFuzz[AccountID](t, 100, withNilChance(0.01))
	assertDecodeNilData[AccountID](t)
	assertEncodeEmptyObj[AccountID](t, 32)
}

func TestAccountID_EncodedLength(t *testing.T) {
	assertEncodedLength(t, []encodedLengthAssert{
		{NewAccountID([]byte{}), 32},
		{NewAccountID([]byte{7, 6, 5, 4, 3, 2, 1, 0}), 32},
	})
}

func TestAccountID_Encode(t *testing.T) {
	assertEncode(t, []encodingAssert{
		{NewAccountID([]byte{0, 0, 0}), MustHexDecodeString("0x0000000000000000000000000000000000000000000000000000000000000000")},     //nolint:lll
		{NewAccountID([]byte{171, 18, 52}), MustHexDecodeString("0xab12340000000000000000000000000000000000000000000000000000000000")}, //nolint:lll
	})
}

func TestAccountID_Hash(t *testing.T) {
	assertHash(t, []hashAssert{
		{NewAccountID([]byte{0, 42, 254}), MustHexDecodeString(
			"0x7834db8eb04aefe8272c32d8160ce4fa3cb31fc95882e5bd53860715731c8198")},
		{NewAccountID([]byte{0, 0}), MustHexDecodeString(
			"0x89eb0d6a8a691dae2cd15ed0369931ce0a949ecafa5c3f93f8121833646e15c3")},
	})
}

func TestAccountID_Hex(t *testing.T) {
	assertEncodeToHex(t, []encodeToHexAssert{
		{NewAccountID([]byte{0, 0, 0}), "0x0000000000000000000000000000000000000000000000000000000000000000"},
		{NewAccountID([]byte{171, 18, 52}), "0xab12340000000000000000000000000000000000000000000000000000000000"},
	})
}

func TestAccountID_String(t *testing.T) {
	assertString(t, []stringAssert{
		{NewAccountID([]byte{0, 0, 0}), "[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"},
		{NewAccountID([]byte{171, 18, 52}), "[171 18 52 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]"},
	})
}

func TestAccountID_Eq(t *testing.T) {
	assertEq(t, []eqAssert{
		{NewAccountID([]byte{1, 0, 0}), NewAccountID([]byte{1, 0}), true},
		{NewAccountID([]byte{0, 0, 1}), NewAccountID([]byte{0, 1}), false},
		{NewAccountID([]byte{0, 0, 0}), NewAccountID([]byte{0, 0}), true},
		{NewAccountID([]byte{12, 48, 255}), NewAccountID([]byte{12, 48, 255}), true},
		{NewAccountID([]byte{0}), NewAccountID([]byte{0}), true},
		{NewAccountID([]byte{1}), NewBool(true), false},
		{NewAccountID([]byte{0}), NewBool(false), false},
	})
}

func TestAccountID_ConvertToHex(t *testing.T) {
	hex, err := ConvertSS58AddressToHex("5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY")
	assert.NoError(t, err)
	assert.Equal(t, "0xd43593c715fdd31c61141abd04a99fd6822c8558854ccde39a5684e7a56da27d", hex)
}

func TestAccountID_ConvertToSS58(t *testing.T) {
	ss58 := ConvertAccountIDToSS58Address(NewAccountID(signature.TestKeyringPairAlice.PublicKey))
	assert.Equal(t, signature.TestKeyringPairAlice.Address, ss58)
}
