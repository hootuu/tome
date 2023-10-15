package uc

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Code string

const (
	NilCoinCode Code = ""
	HTG         Code = "HTG"
	HTC         Code = "HTC"
)

func (c Code) S() string {
	return string(c)
}

func CoinCodeVerify(cStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[0-9a-zA-Z]{1,32}$", cStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid coin code: %s", cStr))
	}
	if strings.HasPrefix(strings.ToUpper(cStr), "HT") ||
		strings.HasPrefix(strings.ToUpper(cStr), "HOT") ||
		strings.HasPrefix(strings.ToUpper(cStr), "HOTU") {
		return errors.Verify("invalid coin code with HT or HOTU")
	}
	return nil
}

func CoinCodeOf(cStr string) (Code, *errors.Error) {
	if err := CoinCodeVerify(cStr); err != nil {
		return NilCoinCode, err
	}
	return Code(strings.ToUpper(cStr)), nil
}
