package common

import (
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/chaincfg"
	"errors"
)

type Network int16
type WICCNet uint32
type WalletStatus int64
type Purpose uint32
type CoinType uint32
type ChangeType uint32

const (
	// MainNet represents the main wicc network.
	MainNet wire.BitcoinNet = 0xff421d1a

	// TestNet represents the  test wicc network.
	TestNet wire.BitcoinNet =0xfd7d5cd7

	// Default entropy size for mnemonic
	DefaultEntropySize = 128
	// Default seed pass. it used to generate seed from mnemonic( BIP39 ). Don't change if determined
	DefaultSeedPass = ""

	HardenedKeyZeroIndex =0x8001869f

	BIP44Purpose Purpose = 44
	WICCCoinType CoinType=99999

	 MAINNET  Network = 1
	TESTNET Network = 2


	ExternalChangeType ChangeType = 0
	InternalChangeType ChangeType = 1
)

type MnemonicLanguage string

// List Mnemonic language support
const (
	ENGLISH  MnemonicLanguage = "EN"
	JAPANESE                  = "JP"
	FRENCH                    = "FR"
	ITALIAN                   = "IT"
	KOREAN                    = "KR"
	SPANISH                   = "ES"
)

func NetworkToChainConfig(net Network) (*chaincfg.Params, error) {
	switch net {
	case TESTNET:
		return &WaykiTestParams, nil

	case MAINNET:
		return &WaykiMainNetParams, nil
	}

	return nil, errors.New("invalid network")
}