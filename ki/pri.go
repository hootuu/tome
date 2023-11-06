package ki

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

type PRI string

func (pri PRI) S() string {
	return string(pri)
}

func (pri PRI) Sign(message string) (string, *errors.Error) {
	privateKey, nErr := crypto.HexToECDSA(pri.S()[2:])
	if nErr != nil {
		logger.Logger.Error("crypto.HexToECDSA(pri.S())", zap.Error(nErr))
		return "", errors.Sys("crypto.HexToECDSA error: " + nErr.Error())
	}

	msgBytes := []byte(message)
	hash := crypto.Keccak256Hash(msgBytes)
	signStr, nErr := crypto.Sign(hash.Bytes(), privateKey)
	if nErr != nil {
		logger.Logger.Error("crypto.Sign(hash.Bytes(), privateKey)", zap.Error(nErr))
		return "", errors.Sys("crypto.Sign error: " + nErr.Error())
	}
	return hexutil.Encode(signStr), nil
}

func (pri PRI) GetPUB() (ADR, PUB, *errors.Error) {
	privateKey, nErr := crypto.HexToECDSA(pri.S()[2:])
	if nErr != nil {
		logger.Logger.Error("crypto.HexToECDSA(pri.S())", zap.Error(nErr))
		return "", "", errors.Sys("crypto.HexToECDSA error: " + nErr.Error())
	}
	addrStr := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	pubBytes := crypto.FromECDSAPub(&privateKey.PublicKey)
	pubKeyStr := hexutil.Encode(pubBytes)
	return ADR(addrStr), PUB(pubKeyStr), nil
}
