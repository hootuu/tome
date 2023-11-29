package du

import (
	"fmt"
	"github.com/hootuu/utils/errors"
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
