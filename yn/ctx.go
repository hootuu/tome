package yn

import "github.com/hootuu/utils/errors"

type Ctx map[string]string

func NewCtx() Ctx {
	return make(map[string]string)
}

func (c Ctx) Put(code string, val string) *errors.Error {
	codeLen := len(code)
	if codeLen == 0 || codeLen > 60 {
		return errors.Verify("length of code must be in (0, 60]")
	}
	valLen := len(val)
	if valLen == 0 || valLen > 360 {
		return errors.Verify("length of value must be in (0, 360]")
	}
	c[code] = val
	return nil
}
