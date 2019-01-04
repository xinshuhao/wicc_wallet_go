package wiccwallet

import (
	"wiccwallet/commons"
)

const WAYKI_TESTNET int = 1
const WAYKI_MAINTNET int = 2

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

//注册账户交易签名
func SignRegisterTx(height int64, fees int64,privateKey string) string {
	var waykiRegister commons.WaykiRegisterTxParams
	waykiRegister.BaseSignTxParams.PrivateKey=privateKey
	waykiRegister.BaseSignTxParams.ValidHeight=height
	waykiRegister.BaseSignTxParams.Fees=fees
	waykiRegister.BaseSignTxParams.TxType=commons.TX_REGISTERACCOUNT
	waykiRegister.BaseSignTxParams.Version=1
	hash:=waykiRegister.SignTX()
	return hash
}


//普通交易签名
func SignCommonTx(value int64,regid string,toAddr string,height int64, fees int64,privateKey string) string {
	var waykicommon commons.WaykiCommonTxParams
	waykicommon.Value=value
	waykicommon.DestAddress=toAddr
	waykicommon.BaseSignTxParams.PrivateKey=privateKey
	waykicommon.BaseSignTxParams.RegId=regid
	waykicommon.BaseSignTxParams.ValidHeight=height
	waykicommon.BaseSignTxParams.Fees=fees
	waykicommon.BaseSignTxParams.TxType=commons.TX_COMMON
	waykicommon.BaseSignTxParams.Version=1
	hash:=waykicommon.SignTX()
	return hash
}

//投票交易签名
func SignDelegateTx(regid string,height int64, fees int64,privateKey string,str string) string {
	var waykiDelegate commons.WaykiDelegateTxParams
	waykiDelegate.BaseSignTxParams.PrivateKey=privateKey
	waykiDelegate.BaseSignTxParams.RegId=regid
	waykiDelegate.BaseSignTxParams.ValidHeight=height
	waykiDelegate.BaseSignTxParams.Fees=fees
	waykiDelegate.BaseSignTxParams.TxType=commons.TX_DELEGATE
	waykiDelegate.BaseSignTxParams.Version=1

	//delegateList:=[]commons.OperVoteFund{commons.OperVoteFund{
	//	commons.MINUS_FUND,pubKeyList[0],10000,
	//}}
	//waykiDelegate.OperVoteFunds=delegateList
	hash:=waykiDelegate.SignTX()
	return hash
}



