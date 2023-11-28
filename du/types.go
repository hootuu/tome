package du

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"math"
	"math/big"
	"regexp"
)

type Code string

func (c Code) S() string {
	return string(c)
}

func CodeVerify(cStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[0-9a-zA-Z]{1,200}$", cStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid coin code: %s", cStr))
	}
	return nil
}

type Wei uint8

const (
	MaxWei     Wei = 28
	DefaultWei Wei = 10
)

func (w Wei) S() string {
	return fmt.Sprintf("%d", w)
}

func (w Wei) PowInt64() int64 {
	return int64(math.Pow(10, float64(w)))
}

func (w Wei) Pow() *big.Int {
	return big.NewInt(int64(math.Pow(10, float64(w))))
}

func (w Wei) PowF() *big.Float {
	return big.NewFloat(math.Pow(10, float64(w)))
}

func (w Wei) Get(f float64) int64 {
	return int64(f * math.Pow(10, float64(w)))
}

func WeiVerify(wei Wei) *errors.Error {
	if wei < 0 || wei > MaxWei {
		return errors.Verify("invalid wei: [0, " + MaxWei.S() + "]")
	}
	return nil
}
