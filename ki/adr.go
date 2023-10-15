package ki

import (
	"github.com/hootuu/utils/errors"
	"regexp"
)

type ADR string

const (
	UnkADR ADR = "?"
)

func (a ADR) S() string {
	return string(a)
}

func ADRVerify(adrStr string) *errors.Error {
	if adrStr == UnkADR.S() {
		return nil
	}
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{1,360}$", adrStr)
	if !matched {
		return errors.Verify("invalid ADR: ^[a-zA-Z0-9]{1,360}$")
	}
	return nil
}

func ADROf(adrStr string) (ADR, *errors.Error) {
	if adrStr == UnkADR.S() {
		return UnkADR, nil
	}
	if err := ADRVerify(adrStr); err != nil {
		return UnkADR, err
	}
	return ADR(adrStr), nil
}
