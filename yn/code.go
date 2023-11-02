package yn

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Code string

const (
	NilCode Code = ""
)

func (a Code) S() string {
	return string(a)
}

func CodeVerify(codeStr string) *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,100}$`, codeStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid yin.code: %s", codeStr))
	}
	return nil
}

func CodeOf(codeStr string) (Code, *errors.Error) {
	if err := CodeVerify(codeStr); err != nil {
		return NilCode, err
	}
	return Code(strings.ToUpper(codeStr)), nil
}
