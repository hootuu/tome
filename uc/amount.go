package uc

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"math"
	"strconv"
	"strings"
)

type WEI int8

func (wei WEI) S() string {
	return fmt.Sprintf("%d", wei)
}

const (
	MaxWei     WEI = 16
	DefaultWei WEI = 10
)

type Amount struct {
	Wei   WEI   `bson:"wei" json:"wei"`
	Left  int64 `bson:"left" json:"left"`
	Right int64 `bson:"right" json:"right"`
}

func (a *Amount) S() string {
	return fmt.Sprintf("wei=%d;left=%d;right=%d;", a.Wei, a.Left, a.Right)
}

func NewAmount(wei WEI, left int64, right int64) (*Amount, *errors.Error) {
	if wei < 0 || wei > MaxWei {
		return nil, errors.Verify("wei must be (0, 10]")
	}
	rightStr := fmt.Sprintf("%d", right)
	if len(rightStr) > int(wei) {
		rightStr = rightStr[:wei]
	}
	rightVal, _ := strconv.ParseInt(rightStr, 10, 64)
	return &Amount{
		Wei:   wei,
		Left:  left,
		Right: rightVal,
	}, nil
}

func NewAmountFloat64(wei WEI, fVal float64) (*Amount, *errors.Error) {
	if wei < 0 || wei > MaxWei {
		return nil, errors.Verify("wei must be (0, 10]")
	}
	valStr := fmt.Sprintf("%f", fVal)
	parts := strings.Split(valStr, ".")
	leftPartStr := parts[0]
	leftVal, _ := strconv.ParseInt(leftPartStr, 10, 64)
	if len(parts) == 0 {
		return NewAmount(wei, leftVal, 0)
	}
	rightPartStr := parts[1]
	var rightVal int64 = 0
	if len(rightPartStr) > int(wei) {
		rightPartStr = rightPartStr[:wei]
	} else {
		if len(rightPartStr) < int(wei) {
			for i := len(rightPartStr); i < int(wei); i++ {
				rightPartStr += "0"
			}
		}
		rightVal, _ = strconv.ParseInt(rightPartStr, 10, 64)
	}
	return &Amount{
		Wei:   wei,
		Left:  leftVal,
		Right: rightVal,
	}, nil
}

func (a *Amount) Dec(decAmount *Amount) *Amount {
	decF := float64(decAmount.Left) + float64(a.Right)*math.Pow10(0-int(a.Wei))
	defA, _ := NewAmountFloat64(a.Wei, 0-decF)
	return a.Inc(defA)
}

func (a *Amount) Float64() float64 {
	f := float64(a.Left) + float64(a.Right)*math.Pow10(0-int(a.Wei))
	return f
}

func (a *Amount) Inc(incAmount *Amount) *Amount {
	originF := float64(a.Left) + float64(a.Right)*math.Pow10(0-int(a.Wei))
	totalF := float64(a.Left) +
		float64(incAmount.Left) +
		float64(a.Right)*math.Pow10(0-int(a.Wei)) +
		float64(incAmount.Right)*math.Pow10(0-int(a.Wei))
	newAmount, _ := NewAmountFloat64(a.Wei, totalF)
	a.Left = newAmount.Left
	a.Right = newAmount.Right
	incF := totalF - originF
	incA, _ := NewAmountFloat64(a.Wei, incF)
	return incA
	//willLeft := a.Left + incAmount.Left
	//willRight := a.Right + incAmount.Right
	//willRightStr := fmt.Sprintf("%d", willRight)
	//if len(willRightStr) > int(a.Wei) {
	//	willLeft += 1
	//	for i := len(willRightStr); i < int(a.Wei); i++ {
	//		willRightStr += "0"
	//	}
	//	willRight, _ = strconv.ParseInt(willRightStr, 10, 64)
	//}
	//pureIncAmount, _ := NewAmount(a.Wei, willLeft-a.Left, willRight-a.Right)
	//a.Left = willLeft
	//a.Right = willRight
	//return pureIncAmount
}

func (a *Amount) Lte(b *Amount) bool {
	if a.Left > b.Left {
		return false
	}
	if a.Left < b.Left {
		return true
	}
	return a.Right <= b.Right
}

func (a *Amount) Lt(b *Amount) bool {
	if a.Left > b.Left {
		return false
	}
	if a.Left < b.Left {
		return true
	}
	return a.Right < b.Right
}

func (a *Amount) Int64Round() int64 {
	f := a.getRightFirstDigit()
	if f >= 5 {
		return a.Left + 1
	}
	return a.Left
}

func (a *Amount) String() string {
	if a.Right == 0 {
		return fmt.Sprintf("%d", a.Left)
	}
	rightStr := fmt.Sprintf("%d", a.Right)
	rightLength := len(rightStr)
	padding := ""
	if rightLength < int(a.Wei) {
		for i := rightLength; i < int(a.Wei); i++ {
			padding += "0"
		}
	}
	return fmt.Sprintf("%d.%s%d", a.Left, padding, a.Right)
}

func (a *Amount) getRightFirstDigit() int64 {
	if a.Right == 0 {
		return 0
	}
	rightStr := fmt.Sprintf("%d", a.Right)
	rightLength := len(rightStr)
	if rightLength < int(a.Wei) {
		return 0
	}
	n := a.Right
	for n >= 10 {
		n /= 10
	}
	return n
}

func (a *Amount) getWeiBase() int64 {
	return int64(math.Pow10(int(a.Wei)))
}
