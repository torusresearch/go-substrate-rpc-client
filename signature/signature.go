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

package signature

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/vedhavyas/go-subkey/v2"
	"github.com/vedhavyas/go-subkey/v2/sr25519"
	"golang.org/x/crypto/blake2b"
)

type KeyringPair struct {
	// URI is the derivation path for the private key in subkey
	URI string
	// Address is an SS58 address
	Address string
	// PublicKey
	PublicKey []byte
}

// KeyringPairFromSecret creates KeyPair based on seed/phrase and network
// Leave network empty for default behavior
func KeyringPairFromSecret(seedOrPhrase string, network uint16) (KeyringPair, error) {
	scheme := sr25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, seedOrPhrase)
	if err != nil {
		return KeyringPair{}, err
	}

	ss58Address := kyr.SS58Address(network)

	var pk = kyr.Public()

	return KeyringPair{
		URI:       seedOrPhrase,
		Address:   ss58Address,
		PublicKey: pk,
	}, nil
}

var TestKeyringPairAlice = KeyringPair{
	URI:       "//Alice",
	PublicKey: []byte{0xd4, 0x35, 0x93, 0xc7, 0x15, 0xfd, 0xd3, 0x1c, 0x61, 0x14, 0x1a, 0xbd, 0x4, 0xa9, 0x9f, 0xd6, 0x82, 0x2c, 0x85, 0x58, 0x85, 0x4c, 0xcd, 0xe3, 0x9a, 0x56, 0x84, 0xe7, 0xa5, 0x6d, 0xa2, 0x7d}, //nolint:lll
	Address:   "5GrwvaEF5zXb26Fz9rcQpDWS57CtERHpNehXCPcNoHGKutQY",
}

var TestKeyringPairBob = KeyringPair{
	URI:       "//Bob",
	PublicKey: []byte{0x8e, 0xaf, 0x04, 0x15, 0x16, 0x87, 0x73, 0x63, 0x26, 0xc9, 0xfe, 0xa1, 0x7e, 0x25, 0xfc, 0x52, 0x87, 0x61, 0x36, 0x93, 0xc9, 0x12, 0x90, 0x9c, 0xb2, 0x26, 0xaa, 0x47, 0x94, 0xf2, 0x6a, 0x48}, //nolint:lll
	Address:   "5FHneW46xGXgs5mUiveU4sbTyGBzmstUspZC92UhjJM694ty",
}

var TestKeyringPairCharlie = KeyringPair{
	URI:       "//Charlie",
	PublicKey: []byte{0x90, 0xb5, 0xab, 0x20, 0x5c, 0x69, 0x74, 0xc9, 0xea, 0x84, 0x1b, 0xe6, 0x88, 0x86, 0x46, 0x33, 0xdc, 0x9c, 0xa8, 0xa3, 0x57, 0x84, 0x3e, 0xea, 0xcf, 0x23, 0x14, 0x64, 0x99, 0x65, 0xfe, 0x22}, //nolint:lll
	Address:   "5FLSigC9HGRKVhB9FiEo4Y3koPsNmBmLJbpXg2mp1hXcS59Y",
}

// Sign signs data with the private key under the given derivation path, returning the signature. Requires the subkey
// command to be in path
func Sign(data []byte, privateKeyURI string) ([]byte, error) {
	// if data is longer than 256 bytes, hash it first
	if len(data) > 256 {
		h := blake2b.Sum256(data)
		data = h[:]
	}

	scheme := sr25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, privateKeyURI)
	if err != nil {
		return nil, err
	}

	signature, err := kyr.Sign(data)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// Verify verifies data using the provided signature and the key under the derivation path. Requires the subkey
// command to be in path
func Verify(data []byte, sig []byte, privateKeyURI string) (bool, error) {
	// if data is longer than 256 bytes, hash it first
	if len(data) > 256 {
		h := blake2b.Sum256(data)
		data = h[:]
	}

	scheme := sr25519.Scheme{}
	kyr, err := subkey.DeriveKeyPair(scheme, privateKeyURI)
	if err != nil {
		return false, err
	}

	if len(sig) != 64 {
		return false, errors.New("wrong signature length")
	}

	v := kyr.Verify(data, sig)

	return v, nil
}

// LoadKeyringPairFromEnv looks up whether the env variable TEST_PRIV_KEY is set and is not empty and tries to use its
// content as a private phrase, seed or URI to derive a key ring pair. Panics if the private phrase, seed or URI is
// not valid or the keyring pair cannot be derived
// Loads Network from TEST_NETWORK variable
// Leave TEST_NETWORK empty or unset for default
func LoadKeyringPairFromEnv() (kp KeyringPair, ok bool) {
	networkString := os.Getenv("TEST_NETWORK")
	network, err := strconv.ParseInt(networkString, 10, 8)
	if err != nil {
		// defaults to generic substrate address
		// https://github.com/paritytech/substrate/wiki/External-Address-Format-(SS58)#checksum-types
		network = 42
	}
	priv, ok := os.LookupEnv("TEST_PRIV_KEY")
	if !ok || priv == "" {
		return kp, false
	}
	kp, err = KeyringPairFromSecret(priv, uint16(network))
	if err != nil {
		panic(fmt.Errorf("cannot load keyring pair from env or use fallback: %v", err))
	}
	return kp, true
}
