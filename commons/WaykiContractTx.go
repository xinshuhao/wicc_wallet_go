package commons

import (
	"bytes"
	"encoding/hex"
	"github.com/btcsuite/btcutil"
)

type WaykiContractTxParams struct {
	BaseSignTxParams
	Value int64
	Appid string
	ContractBytes []byte
}

func (waykiContract WaykiContractTxParams)SignTX()string{
	regId:=parseRegId(waykiContract.RegId)
	bytesBuffer := bytes.NewBuffer([]byte{})
	bytesBuffer.WriteByte(byte(waykiContract.TxType))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Version))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.ValidHeight))
	bytesBuffer.Write(EncodeInOldWay(regId[0]))
	bytesBuffer.Write(EncodeInOldWay(regId[1]))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Fees))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Value))
	bytesBuffer.Write(EncodeInOldWay(int64(len(waykiContract.ContractBytes))))
	bytesBuffer.Write(waykiContract.ContractBytes)
	ss9:=signContractTX(waykiContract)
	bytesBuffer.Write(EncodeInOldWay(int64(len(ss9))))
	bytesBuffer.Write(ss9)
	signHex:=hex.EncodeToString(bytesBuffer.Bytes())
	return signHex
}

func signContractTX(waykiContract WaykiContractTxParams) []byte{
	regId:=parseRegId(waykiContract.RegId)
	bytesBuffer := bytes.NewBuffer([]byte{})
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Version))
	bytesBuffer.WriteByte(byte(waykiContract.TxType))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.ValidHeight))
	bytesBuffer.Write(EncodeInOldWay(regId[0]))
	bytesBuffer.Write(EncodeInOldWay(regId[1]))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Fees))
	bytesBuffer.Write(EncodeInOldWay(waykiContract.Value))
	bytesBuffer.Write(EncodeInOldWay(int64(len(waykiContract.ContractBytes))))
	bytesBuffer.Write(waykiContract.ContractBytes)
	data1,_:=HashDoubleSha256(bytesBuffer.Bytes())
	wif,_ := btcutil.DecodeWIF(waykiContract.PrivateKey)
	key:=wif.PrivKey
	ss,_:=key.Sign(data1)
	return ss.Serialize()
}
