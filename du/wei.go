package du

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"math"
	"math/big"
)

type Wei uint8

const (
	MaxWei     Wei = 18
	DefaultWei Wei = 2

	ErrWeiInvalid = "Invalid Wei: 0 <= wei <= 18"
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
		return errors.Verify(ErrWeiInvalid)
	}
	return nil
}
