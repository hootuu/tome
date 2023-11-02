package bk

import (
	"fmt"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/tome/yn"
)

type Chain string

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
