package ki

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/logger"
	"go.uber.org/zap"
)

type PUB string

func (pub PUB) S() string {
	return string(pub)
}

func (pub PUB) VerifySignature(message string, signStr string) (bool, *errors.Error) {
	publicKeyBytes, nErr := hexutil.Decode(pub.S())
	if nErr != nil {
		logger.Logger.Error("hex.DecodeString(pub.S()) failed", zap.Error(nErr))
		return false, errors.Sys("hex.DecodeString failed: " + nErr.Error())
	}
	msgBytes := []byte(message)
	signBytes, nErr := hexutil.Decode(signStr)
	if nErr != nil {
		logger.Logger.Error("hex-util.Decode(signStr) failed", zap.Error(nErr))
		return false, errors.Sys("hex-util.Decode(signStr) failed: " + nErr.Error())
	}

	hash := crypto.Keccak256Hash(msgBytes)
	isValid := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signBytes[:64])
	return isValid, nil
}
