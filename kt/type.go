package kt

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
)

type Type string

func (t Type) Verify() *errors.Error {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._]{3,200}$`, t.S())
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid type: %s", t.S()))
	}
	return nil
}

func (t Type) S() string {
	return string(t)
}
