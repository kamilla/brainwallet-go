// generator.go - Generator routines
// Copyright (c) 2015 Kamilla Productions Uninc. Author Joonas Greis  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.

package brainwallet

import (
        "golang.org/x/crypto/ripemd160"
        "github.com/jbenet/go-base58"
        "github.com/haltingstate/secp256k1-go"
        "crypto/sha256"
        "encoding/hex"
        "math/big"
	"sync"
	"hash"
)

// Generator
func Generator(input chan string, output chan string, done chan int, wg *sync.WaitGroup) {

	defer wg.Done()

        waitfordone:
        for {
                select {
                case passphrase := <- input:				// Receive passphrase

                        hasher := sha256.New()				// SHA256
			sha := SHA256(hasher, []byte(passphrase))

                        publicKeyBytes := secp256k1.UncompressedPubkeyFromSeckey(sha)	// ECDSA
                        privateKey := hex.EncodeToString(sha)		// Store Private Key

                        sha = SHA256(hasher, publicKeyBytes)		// SHA256
                        ripe := RIPEMD160(sha)				// RIPEMD160

                        versionripe := "00" + hex.EncodeToString(ripe)	// Add version byte 0x00
                        decoded, _ := hex.DecodeString(versionripe)

                        sha = SHA256(hasher, SHA256(hasher, decoded))	// SHA256x2

                        addressChecksum := hex.EncodeToString(sha)[0:8]	// Concencate Address Checksum and Extended RIPEMD160 Hash
                        hexBitcoinAddress := versionripe + addressChecksum

                        bigintBitcoinAddress, _ := new(big.Int).SetString((hexBitcoinAddress),16)	// Base58Encode the Address
                        base58BitcoinAddress := base58.Encode(bigintBitcoinAddress.Bytes())

                        line := "1" + base58BitcoinAddress + ":" + privateKey + ":" + passphrase	// Create a line for io output
                        output <- line					// Send line to output channel

                case <- done:						// Everything is done. Break out from the loop.
                        break waitfordone
                }
        }
}

// SHA256 Hasher function
func SHA256(hasher hash.Hash, input []byte) (hash []byte) {

	hasher.Reset()
       	hasher.Write(input)
        hash = hasher.Sum(nil)
	return hash

}

// RIPEMD160 Hasher function
func RIPEMD160(input []byte) (hash []byte) {

	riper := ripemd160.New()
        riper.Write(input)
        hash = riper.Sum(nil)
	return hash

}
