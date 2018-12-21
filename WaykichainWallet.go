package wiccwallet

import (
	"wiccwallet/mnemonic"
	"wiccwallet/common"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"wiccwallet/bip"
	"testing"
)

type Wallet struct {
	Entropy     string
	Mnemonic    string
	Seed        string
	MasterNode  *hdkeychain.ExtendedKey
	PurposeNode *hdkeychain.ExtendedKey
	Coins       []*Coin
}

type Coin struct {
	Name    string
	Coin    *hdkeychain.ExtendedKey
	Network *chaincfg.Params
}

func ImportMnemonic(t *testing.T){
	//seed := bip.NewSeed("announce parent popular hybrid fine maid exile impulse unknown school castle wage hand impulse wing", "")
	//fmt.Println(hex.EncodeToString(seed))
	
	mnemonic:="announce parent popular hybrid fine maid exile impulse unknown school castle wage hand impulse wing"

	entropy, _ := bip.EntropyFromMnemonic(mnemonic)
	entropyToHexString := hex.EncodeToString(entropy)

	seed, err := bip.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		fmt.Println(err)
	}
	seedToHexString := hex.EncodeToString(seed)

	//@ToDo: create network params for FLO and LTC, etc

	masterKey, _ := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	purposeNode, _ := masterKey.Child(0x8001869f)

	wallet := Wallet{
		Mnemonic: mnemonic,
		Seed:     seedToHexString,
		//ToDo: Derive entropy from mnemonic or seed
		Entropy:     entropyToHexString,
		MasterNode:  masterKey,
		PurposeNode: purposeNode,
	}
	fmt.Println(wallet)
}


//创建助记词
func CreateMnemonics() (string){
   mn:= mnemonic.NewMnemonicWithLanguage(common.ENGLISH)
   words,err:=mn.GenerateMnemonic()
   if err!=nil{
   	return ""
   }
   return words
}

