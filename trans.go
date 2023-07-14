package main

import (
	"context"
	"log"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/unlocker"
)

func main() {
	tx := bt.NewTx()

	_ = tx.From(
		"11b476ad8e0a48fcd40807a111a050af51114877e09283bfa7f3505081a1819d",
		0,
		"76a914eb0bd5edba389198e73f8efabddfc61666969ff788ac6a0568656c6c6f",
		1500,
	)

	_ = tx.PayToAddress("1NRoySJ9Lvby6DuE2UQYnyT67AASwNZxGb", 1000)

	decodedWif, _ := wif.DecodeWIF("KznvCNc6Yf4iztSThoMH6oHWzH9EgjfodKxmeuUGPq5DEX5maspS")

	if err := tx.FillAllInputs(context.Background(), &unlocker.Getter{PrivateKey: decodedWif.PrivKey}); err != nil {
		log.Fatal(err.Error())
		
	}
	
	err := tx.FillInput(
		context.Background(),
		&unlocker.Simple{
			PrivateKey: decodedWif.PrivKey,
		},
		bt.UnlockerParams{uint32(index), sighash.AllForkID},
	)


	log.Printf("tx: %s\n", tx)
}
