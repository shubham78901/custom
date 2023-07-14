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

	MultipleTransaction()
}

type InputUtxo struct {
	Txid   string
	Index  int
	Script string
	Value  float64
	WifKey string
}

type OutputUtxo struct {
	Amount  uint64
	Address string
}

// func MultipleTransaction(numInputUtxo int, inputUtxo []InputUtxo, outputUtxo []OutputUtxo, changeaddress string, numoutputTUtxo int) {
func MultipleTransaction() {
	fmt.Println("hello rohan 1")
	numInputUtxo := 1
	inputUtxo := []InputUtxo{
		{
			Txid:   "469da0434b1f03409c8e6ed30adef13bee55f182a9ee81a469314afa894b5f89",
			Index:  0,
			Script: "76a9145bc4c9228dc6a243f545e43fe867e295c001739088ac",
			Value:  8000,
			WifKey: "L5V6fPzoAF19jgGK3gaBZTEkBHNy2J27XH1tiGom9ZW5U43GSuVk",
		},
	}
	outputUtxo := []OutputUtxo{
		{
			Amount:  3000,
			Address: "1K81b9eR7iAaAZAKjoKHXWMQYfKyCc1zKP",
		},
	}
	changeAddress := "1K81b9eR7iAaAZAKjoKHXWMQYfKyCc1zKP"
	numoutputUtxo := 1

	howManyUTXOs := numInputUtxo

	tx := bt.NewTx()
	fmt.Println("hello rohan 2")
	var wifKeys = []string{}

	for i := 0; i < howManyUTXOs; i++ {

		//	utxo := GetInput("Enter funding UTXO [TXID index script value wifKey]")
		fmt.Println("hello rohan 3")
		txId := inputUtxo[i].Txid
		index := inputUtxo[i].Index
		script := inputUtxo[i].Script
		value := inputUtxo[i].Value
		wifKeys = append(wifKeys, inputUtxo[i].WifKey)

		//Add input UTXO
		_ = tx.From(
			txId,
			uint32(index),
			script,
			uint64(value),
		)

	}

	howManyOutputs := numoutputUtxo

	for i := 0; i < howManyOutputs; i++ {

		//	utxo := GetInput("Enter the reciever address and how many satoshis to be sent [address satoshis]")
		fmt.Println("hello rohan 4")
		satoshis := outputUtxo[i].Amount

		_ = tx.AddP2PKHOutputFromAddress(outputUtxo[i].Address, satoshis)
	}

	changeaddress := changeAddress

	//Calc Fee
	_ = tx.ChangeToAddress(changeaddress, bt.NewFeeQuote())

	//Unlock input UTXOs
	for index, _ := range wifKeys {

		decodedWif, _ := wif.DecodeWIF(wifKeys[index])
		err := tx.FillInput(
			context.Background(),
			&unlocker.Simple{
				PrivateKey: decodedWif.PrivKey,
			},
			bt.UnlockerParams{uint32(index), sighash.AllForkID},
		)

		if err != nil {
			log.Fatalf(err.Error())
		}

	}

	fmt.Printf("%s\n", tx)
	fmt.Println("hello rohan 4")

}
