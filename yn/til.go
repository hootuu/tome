package yn

import (
	"github.com/hootuu/utils/errors"
)

type Title string

const (
	NilTitle Title = ""
)

func (t Title) S() string {
	return string(t)
}

func TitleVerify(tilStr string) *errors.Error {
	length := len(tilStr)
	if length == 0 || length > 360 {
		return errors.Verify("length of title must be in (0, 360]")
	}
	return nil
}

func TitleOf(tilStr string) (Title, *errors.Error) {
	if err := TitleVerify(tilStr); err != nil {
		return NilTitle, err
	}
	return Title(tilStr), nil
}
