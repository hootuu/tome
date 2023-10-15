package yn

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Act string

const (
	NilAct Act = ""
)

func (a Act) S() string {
	return string(a)
}

func ActVerify(actStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9\\._]{3,60}$", actStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid currency: %s", actStr))
	}
	return nil
}

func ActOf(actStr string) (Act, *errors.Error) {
	if err := ActVerify(actStr); err != nil {
		return NilAct, err
	}
	return Act(strings.ToUpper(actStr)), nil
}
