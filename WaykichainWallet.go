package wicc_wallet_go

import (
	"wicc_wallet_go/commons"
	"encoding/hex"
	"encoding/json"
)

var NetWorkType int = 2

func UseTestNet(useTest bool) {
	if useTest {
		NetWorkType = 1
	} else {
		NetWorkType = 2
	}
}

//创建助记词
func CreateMnemonics() (string) {
	mn := commons.NewMnemonicWithLanguage(commons.ENGLISH)
	words, err := mn.GenerateMnemonic()
	if err != nil {
		return ""
	}
	return words
}

//助记词转换地址
func Mnemonic2Address(words string) (string) {
	address := commons.GenerateAddress(words, NetWorkType)
	return address
}

//助记词转私钥
func Mnemonic2PrivateKey(words string, ) (string) {
	privateKey := commons.GeneratePrivateKey(words, NetWorkType)
	return privateKey
}

func ValidatePrivateKey(privateArray []byte) bool{
	err:=	commons.ValidatePrivateKey(privateArray)
	if err!=nil{
		return false
	}else{
		return true
	}
}

//私钥转地址
func PrivateKey2Address(words string) (string) {
	address := commons.ImportPrivateKey(words, NetWorkType)
	return address
}

//注册账户交易签名
func SignRegisterTx(height int64, fees int64, privateKey string) string {
	var waykiRegister commons.WaykiRegisterTxParams
	waykiRegister.BaseSignTxParams.PrivateKey = privateKey
	waykiRegister.BaseSignTxParams.ValidHeight = height
	waykiRegister.BaseSignTxParams.Fees = fees
	waykiRegister.BaseSignTxParams.TxType = commons.TX_REGISTERACCOUNT
	waykiRegister.BaseSignTxParams.Version = 1
	hash := waykiRegister.SignTX()
	return hash
}

//普通交易签名
func SignCommonTx(value int64, regid string, toAddr string, height int64, fees int64, privateKey string) string {
	var waykicommon commons.WaykiCommonTxParams
	waykicommon.Value = value
	waykicommon.DestAddress = toAddr
	waykicommon.BaseSignTxParams.PrivateKey = privateKey
	waykicommon.BaseSignTxParams.RegId = regid
	waykicommon.BaseSignTxParams.ValidHeight = height
	waykicommon.BaseSignTxParams.Fees = fees
	waykicommon.BaseSignTxParams.TxType = commons.TX_COMMON
	waykicommon.BaseSignTxParams.Version = 1
	hash := waykicommon.SignTX()
	return hash
}

type Vote struct {
	VoteType  int    `json:"VoteType"`
	PubKey    string `json:"PubKey"`
	VoteValue int64  `json:"VoteValue"`
}

func (this Vote) MarshalJSON() ([]byte, error) {
	coder := new(struct {
		VoteType  int    `json:"VoteType"`
		PubKey    string `json:"PubKey"`
		VoteValue int64  `json:"VoteValue"`
	})
	coder.VoteValue = this.VoteValue
	coder.VoteType = this.VoteType
	coder.PubKey = this.PubKey
	return json.Marshal(coder)
}

func  unmarshalJSON(in []byte) (Vote,error) {
	decoder := new(struct {
		VoteType  int    `json:"VoteType"`
		PubKey    string `json:"PubKey"`
		VoteValue int64  `json:"VoteValue"`
	})
	err := json.Unmarshal(in, decoder)
	votes :=Vote{}
	if err == nil {
		votes.VoteValue = decoder.VoteValue
		votes.VoteType = decoder.VoteType
		votes.PubKey = decoder.PubKey
	}
	return votes,err
}

//投票交易签名
func SignDelegateTx(regid string, height int64, fees int64, privateKey string, voteListJson string) string {
	var waykiDelegate commons.WaykiDelegateTxParams
	var voteListByte [][]byte
	waykiDelegate.BaseSignTxParams.PrivateKey = privateKey
	waykiDelegate.BaseSignTxParams.RegId = regid
	waykiDelegate.BaseSignTxParams.ValidHeight = height
	waykiDelegate.BaseSignTxParams.Fees = fees
	waykiDelegate.BaseSignTxParams.TxType = commons.TX_DELEGATE
	waykiDelegate.BaseSignTxParams.Version = 1
	json.Unmarshal([]byte(voteListJson), &voteListByte)
	votes := []commons.OperVoteFund{}
	for _, vote := range voteListByte {
		voter, _ :=unmarshalJSON(vote)
		var vote commons.OperVoteFund
		pk,_ := hex.DecodeString(voter.PubKey)
		vote.PubKey = pk
		vote.VoteType = voter.VoteType
		vote.VoteValue = voter.VoteValue
		votes = append(votes, vote)
	}

	waykiDelegate.OperVoteFunds = votes
	hash := waykiDelegate.SignTX()
	return hash
}

//智能合约交易签名
func SignContractTx(value int64, height int64, fees int64, privateKey string, regId string, appid string, contractStr string) string {
	var waykiContract commons.WaykiContractTxParams
	waykiContract.Value = value
	waykiContract.BaseSignTxParams.PrivateKey = privateKey
	waykiContract.BaseSignTxParams.RegId = regId
	waykiContract.Appid = appid
	waykiContract.BaseSignTxParams.ValidHeight = height
	waykiContract.BaseSignTxParams.Fees = fees
	waykiContract.BaseSignTxParams.TxType = commons.TX_CONTRACT
	waykiContract.BaseSignTxParams.Version = 1
	binary, _ := hex.DecodeString(contractStr)
	waykiContract.ContractBytes = []byte(binary)
	hash := waykiContract.SignTX()
	return hash
}
