package du

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type Du struct {
	W Wei   `bson:"w" json:"w"`
	N bool  `bson:"n" json:"n"` //negative
	L int64 `bson:"l" json:"l"`
	R int64 `bson:"r" json:"r"`
}

func FromLR(w Wei, neg bool, iL int64, iR int64) Du {
	return Du{
		W: w,
		N: neg,
		L: iL,
		R: iR,
	}
}

func FromInt64(w Wei, d int64) Du {
	m := Du{
		W: w,
		N: false,
		L: d,
		R: 0,
	}
	if d < 0 {
		m.N = true
		m.L = 0 - d
	}
	return m
}

func FromFloat64(w Wei, f float64) Du {
	nStr := "+"
	if f < 0 {
		nStr = "-"
	}
	str := fmt.Sprintf("%s%d:%.*f", nStr, w, int(w), math.Abs(f))
	dM, _ := FromString(str)
	return dM
}

func FromByte(bArr []byte) (Du, *errors.Error) {
	return FromString(string(bArr))
}

func MustFromString(str string) Du {
	dM, _ := FromString(str)
	return dM
}

func FromString(str string) (Du, *errors.Error) {
	if !strings.HasPrefix(str, "+") && !strings.HasPrefix(str, "-") {
		return Du{}, errors.Verify("Invalid Du string, start with +-")
	}
	isN := false
	nStr := str[0:1]
	switch nStr {
	case "+":
		isN = false
	case "-":
		isN = true
	default:
		isN = false
	}
	str = str[1:]
	weiAndNumbArr := strings.Split(str, ":")
	if len(weiAndNumbArr) != 2 {
		return Du{}, errors.Verify("Invalid Du string, no wei info")
	}
	weiStr := weiAndNumbArr[0]
	weiVal := strs.Int64Of(weiStr, -1)
	if weiVal < 0 {
		return Du{}, errors.Verify("Invalid Du string")
	}
	w := Wei(weiVal)
	if err := WeiVerify(w); err != nil {
		return Du{}, errors.Verify("Invalid Du string:" + err.Error())
	}
	numbStr := weiAndNumbArr[1]
	if b := IsPositiveOrDecimalNumber(numbStr); !b {
		return Du{}, errors.Verify("Not a valid integer or decimal")
	}
	strLRArr := strings.Split(numbStr, ".")
	strLRLen := len(strLRArr)
	if strLRLen < 1 || strLRLen > 2 {
		return Du{}, errors.Verify("Not a valid integer or decimal")
	}
	dM := Du{
		W: w,
		N: isN,
		L: 0,
		R: 0,
	}
	dL, nErr := strconv.ParseInt(strLRArr[0], 10, 64)
	if nErr != nil {
		return Du{}, errors.Verify("Not a valid integer")
	}
	dM.L = dL
	if strLRLen == 2 {
		strR := strLRArr[1]
		strRLen := len(strR)
		for i := strRLen; i < int(w); i++ {
			strR += "0"
		}
		dR, nErr := strconv.ParseInt(strR, 10, 64)
		if nErr != nil {
			return Du{}, errors.Verify("Not a valid decimal")
		}
		dM.R = dR
	}

	return dM, nil
}

func (d Du) Bytes() []byte {
	str := d.String()
	return []byte(str)
}

func (d Du) String() string {
	nStr := "+"
	if d.N {
		nStr = "-"
	}
	return fmt.Sprintf("%s%d:%s", nStr, d.W, d.toBigString())
}

func (d Du) View() string {
	nStr := ""
	if d.N {
		nStr = "-"
	}
	return fmt.Sprintf("%s%s", nStr, d.toBigString())
}

// Compare
//
// -1 if d <  other
//
//	0 if d == other
//
// +1 if d >  other
func (d Du) Compare(other Du) int {
	dB := d.toBigInt()
	oB := other.toBigInt()
	return dB.Cmp(oB)
}

func (d Du) Add(other Du) (Du, *errors.Error) {
	if other.W != d.W {
		return Du{}, errors.Verify("It is not allowed to add two values of different WEI")
	}
	if other.L > math.MaxInt64-d.L {
		return Du{}, errors.Verify("Numerical overflow")
	}
	return fromBigInt(d.W,
		new(big.Int).Add(
			d.toBigInt(),
			other.toBigInt(),
		)), nil
}

func (d Du) Subtract(other Du) (Du, *errors.Error) {
	if other.W != d.W {
		return Du{}, errors.Verify("It is not allowed to subtract two values of different WEI")
	}
	if d.L < math.MinInt64+other.L {
		return Du{}, errors.Verify("Numerical overflow")
	}
	return fromBigInt(d.W,
		new(big.Int).Sub(
			d.toBigInt(),
			other.toBigInt(),
		)), nil
}

func (d Du) Multiply(mul int64) (Du, *errors.Error) {
	return fromBigInt(d.W, new(big.Int).Mul(
		d.toBigInt(), big.NewInt(mul),
	)), nil
}

func (d Du) MultiplyF(mul float64) (Du, *errors.Error) {
	mulF := new(big.Float).Mul(big.NewFloat(mul), d.W.PowF())
	mulI, acy := mulF.Int(nil)
	if acy != big.Exact {
		return Du{}, errors.Verify("Numerical overflow")
	}
	return fromBigInt(d.W, new(big.Int).Div(
		new(big.Int).Mul(
			d.toBigInt(),
			mulI,
		), d.W.Pow())), nil
}

func (d Du) Divide(div int64) (Du, *errors.Error) {
	return fromBigInt(d.W, new(big.Int).Div(d.toBigInt(), big.NewInt(div))), nil
}

func (d Du) DivideF(div float64) (Du, *errors.Error) {
	divF := new(big.Float).Mul(big.NewFloat(div), d.W.PowF())
	divI, acy := divF.Int(nil)
	if acy != big.Exact {
		return Du{}, errors.Verify("Numerical overflow")
	}
	return fromBigInt(d.W,
		new(big.Int).Div(
			new(big.Int).Mul(
				d.toBigInt(),
				d.W.Pow(),
			),
			divI,
		)), nil
}

func (d Du) doGetR() string {
	rStr := fmt.Sprintf("%d", d.R)
	for i := len(rStr); i < int(d.W); i++ {
		rStr = "0" + rStr
	}
	return strings.TrimRight(rStr, "0")
}

func (d Du) toBigString() string {
	pureR := d.doGetR()
	if len(pureR) == 0 {
		return fmt.Sprintf("%d", d.L)
	}
	return fmt.Sprintf("%d.%s", d.L, pureR)
}

func (d Du) toBigFloat() *big.Float {
	bf := new(big.Float).SetPrec(uint(d.W)).SetInt(d.toBigInt())
	return bf
}

func (d Du) toBigInt() *big.Int {
	bL := new(big.Int).Mul(big.NewInt(d.L), d.W.Pow())
	bR := big.NewInt(d.R)
	t := new(big.Int).Add(bL, bR)
	if d.N {
		t = new(big.Int).Sub(big.NewInt(0), t)
	}
	return t
}

func fromBigInt(w Wei, bi *big.Int) Du {
	wp := w.Pow()
	xbi := new(big.Int).Abs(bi)
	nL := new(big.Int).Div(xbi, wp)
	nR := new(big.Int).Mod(xbi, wp)
	return Du{
		W: w,
		N: bi.Cmp(big.NewInt(0)) < 0,
		L: nL.Int64(),
		R: nR.Int64(),
	}

}

func IsPositiveOrDecimalNumber(str string) bool {
	pattern := `^[-+]?\d+(\.\d+)?$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
