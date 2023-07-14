package main

import (
	"fmt"

	"github.com/libsv/go-bk/bec"

	"github.com/libsv/go-bk/chaincfg"

	"encoding/hex"

	"github.com/libsv/go-bk/crypto"
	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2/bscript"
)

func main() {
	privKey, _ := bec.NewPrivateKey(bec.S256())
	wi, _ := wif.NewWIF(privKey, &chaincfg.MainNet, true)
	// fmt.Println("first", privKey)

	fmt.Printf("WIF: %s\n", wi)

	pubKey := privKey.PubKey().SerialiseCompressed()

	pubKey = crypto.Sha256(pubKey)

	pubKey = crypto.Ripemd160(pubKey)

	version := make([]byte, 0)

	version = append(version, 0x00) // 0x4d for testnet

	pubKey = append(version, pubKey...)

	address := bscript.Base58EncodeMissingChecksum(pubKey)

	a, _ := bscript.NewAddressFromPublicKey(privKey.PubKey(), true)

	// fmt.Printf("PubKeyPoint X: %s\n", privKey.PublicKey.X)
	// fmt.Printf("PubKeyPoint Y: %s\n", privKey.PublicKey.Y)
	fmt.Printf("Compressed PubKey: %s\n", hex.EncodeToString(privKey.PubKey().SerialiseCompressed()))
	fmt.Printf("Address: %s\n", address)
	fmt.Printf("verfication: %s\n", a.AddressString)

}
