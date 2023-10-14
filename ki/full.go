package ki

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hootuu/utils/errors"
	"github.com/tyler-smith/go-bip39"
)

type Ki struct {
	ADR ADR      `json:"adr"`
	PUB PUB      `json:"pub"`
	PRI PRI      `json:"pri"`
	MIC Mnemonic `json:"mic"`
}

func NewKi() (*Ki, *errors.Error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, errors.Sys("generate private key failed: " + err.Error())
	}

	addrStr := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyStr := hexutil.Encode(privateKeyBytes)

	pubBytes := crypto.FromECDSAPub(&privateKey.PublicKey)
	pubKeyStr := hexutil.Encode(pubBytes)

	micStr, err := bip39.NewMnemonic(privateKeyBytes)
	if err != nil {
		return nil, errors.Sys("new mnemonic failed: " + err.Error())
	}

	return &Ki{
		ADR: ADR(addrStr),
		PUB: PUB(pubKeyStr),
		PRI: PRI(privateKeyStr),
		MIC: Mnemonic(micStr),
	}, nil
}
