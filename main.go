package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type JsonSha struct {
	Type       string `json:"Type"`
	PrivateKey string `json:"PrivateKey"`
}

func main() {
	var mnemonic string
	var hdpath string

	flag.StringVar(&mnemonic, "mnemonic", "", "Mnemonic")
	flag.StringVar(&hdpath, "path", "", "HD path")
	flag.Parse()

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath(hdpath)
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes, err := wallet.PrivateKeyBytes(account)
	if err != nil {
		log.Fatal(err)
	}

	jsonSha := JsonSha{"secp256k1", base64.StdEncoding.EncodeToString(privateKeyBytes)}
	jsonData, err := json.Marshal(jsonSha)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("public address:", account.Address.Hex())
	fmt.Println("private key:", fmt.Sprintf("%s%s", "0x", privateKey))
	fmt.Printf("filecoin key: %x\n", []byte(string(jsonData)))
}
