package yn

import (
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Ref struct {
	Code string `bson:"code" json:"code"`
	Ref  string `bson:"ref" json:"ref"`
}

func (ref *Ref) S() string {
	return ref.Ref + "@" + ref.Code
}

func RefOf(codeStr string, refStr string) (*Ref, *errors.Error) {
	codeMatched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,100}$`, codeStr)
	if !codeMatched {
		return nil, errors.Verify("invalid ref.code: [a-zA-Z0-9._]{3,100}$")
	}
	refLen := len(refStr)
	if refLen == 0 || refLen > 360 {
		return nil, errors.Verify("the length of ref.ref must be in (0, 360]")
	}
	return &Ref{
		Code: strings.ToUpper(codeStr),
		Ref:  refStr,
	}, nil
}
