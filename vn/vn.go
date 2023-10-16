package vn

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type ID string

const (
	HOTU ID = "HOTU"
)

func (id ID) S() string {
	return string(id)
}

const (
	NilID ID = ""
)

func IDVerify(idStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{1,108}$", idStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid vn id: %s", idStr))
	}
	return nil
}

func IDOf(idStr string) (ID, *errors.Error) {
	err := IDVerify(idStr)
	if err != nil {
		return NilID, err
	}
	return ID(strings.ToUpper(idStr)), nil
}

type VN struct {
	ID         string `bson:"id" json:"id"`
	Originator ki.ADR `bson:"originator" json:"originator"`
	Guardian   ki.ADR `bson:"guardian" json:"guardian"`
	Dob        int64  `bson:"dob" json:"dob"` //date of birth
}
