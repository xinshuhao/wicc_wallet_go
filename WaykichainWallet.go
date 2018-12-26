package wiccwallet

import "wiccwallet/commons"

//创建助记词
func CreateMnemonics() (string){
   mn:= commons.NewMnemonicWithLanguage(commons.ENGLISH)
   words,err:=mn.GenerateMnemonic()
   if err!=nil{
   	return ""
   }
   return words
}

//助记词转换地址
func Mnemonic2Address(words string,netType int)(string){
	address := commons.GenerateAddress(words,netType)
	return address
}

//助记词转私钥
func Mnemonic2PrivateKey(words string,netType int)(string){
	privateKey := commons.GeneratePrivateKey(words,netType)
	return privateKey
}

//私钥转地址
func PrivateKey2Address(words string,netType int)(string){
	address := commons.ImportPrivateKey(words,netType)
	return address
}

const WAYKI_TESTNET int = 1
const WAYKI_MAINTNET int = 2


