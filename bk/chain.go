package bk

import (
	"fmt"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/tome/yn"
	"github.com/hootuu/utils/errors"
	"regexp"
)

type Chain string

func (c Chain) Verify() *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{1,200}$`, c.S())
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid chain: %s", c.S()))
	}
	return nil
}

func (c Chain) S() string {
	return string(c)
}

func VNChain() Chain {
	return "hotu.vn"
}

func CoinChain() Chain {
	return Chain(fmt.Sprintf("hotu.coin"))
}

func SPChain(vnID vn.ID) Chain {
	return Chain(fmt.Sprintf("hotu.%s.sp", vnID.S()))
}

func YinChain(vnID vn.ID, yin yn.Code) Chain {
	return Chain(fmt.Sprintf("hotu.%s.yin.%s", vnID.S(), yin.S()))
}
