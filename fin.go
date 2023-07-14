package main

import (
	"context"
	"fmt"
	"log"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/sighash"
	"github.com/libsv/go-bt/v2/unlocker"
)

func main() {
	WriteDataToBlockchain()
}

func WriteDataToBlockchain() {
	tx := bt.NewTx()
	// unlocking, err := bscript.NewFromASM("76a91493e3ac50804356fb8ed57cb911d8567740a2773888ac")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err := tx.From(
		"ca32b83f178c2aef51ca54a8a89f560091618161a43d489d2e239af77771e655",
		0,
		"76a91493e3ac50804356fb8ed57cb911d8567740a2773888ac",
		800,
	)

	if err != nil {
		log.Fatal(err)
	}
	//defined in library
	err = tx.AddHashPuzzleOutput("shubham", "76a91493e3ac50804356fb8ed57cb911d8567740a2773888ac", 600)
	if err != nil {
		log.Fatal(err)
	}

	decodedWif, _ := wif.DecodeWIF("L1Dgi5VxvdFuYiWrK3v8T51Wuu4t9Kjkn3qCs4XNkNEDhuDvUCG4")

	err = tx.FillInput(context.Background(), &unlocker.Simple{PrivateKey: decodedWif.PrivKey},
		bt.UnlockerParams{uint32(0), sighash.AllForkID},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.String())
}
