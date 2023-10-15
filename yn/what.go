package yn

import "github.com/hootuu/utils/errors"

type What Ref

func NewWhat(codeStr string, refStr string) (*What, *errors.Error) {
	ref, err := RefOf(codeStr, refStr)
	if err != nil {
		return nil, errors.Verify("invalid what: " + err.Error())
	}
	return (*What)(ref), nil
}
