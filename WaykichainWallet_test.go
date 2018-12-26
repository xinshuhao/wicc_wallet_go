package wiccwallet

import (
	"testing"
	"fmt"
	"wiccwallet/commons"
)

func TestMnemonic(t *testing.T) {
	mnemonic := "empty regular curve turtle student prize toy accuse develop spike scatter ginger"
	//seed := bip.NewSeed(mnemonic, "")
	////fmt.Println(hex.EncodeToString(seed))
	address := commons.GenerateAddress(mnemonic, WAYKI_MAINTNET)
	fmt.Println("地址"+address)
}

func TestMnemonicWIF(t *testing.T) {
	mnemonic := "empty regular curve turtle student prize toy accuse develop spike scatter ginger"
	privateKey := commons.GeneratePrivateKey(mnemonic, WAYKI_MAINTNET)
	fmt.Println("私钥"+privateKey)

	fmt.Println("地址"+commons.ImportPrivateKey(privateKey,WAYKI_MAINTNET))
}

