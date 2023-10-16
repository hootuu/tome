package sp

import (
	"fmt"
	"github.com/hootuu/tome/fq"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type ID string

func (id ID) S() string {
	return string(id)
}

const (
	NilID ID = ""
)

func IDVerify(idStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{1,108}$", idStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid sp id: %s", idStr))
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

type SP struct {
	VN         vn.ID    `bson:"vn" json:"vn"`
	ID         ID       `bson:"id" json:"id"`
	Link       []ID     `bson:"link" json:"link"`
	Originator ki.ADR   `bson:"originator" json:"originator"`
	Guardian   ki.ADR   `bson:"guardian" json:"guardian"`
	Dob        int64    `bson:"dob" json:"dob"` //date of birth
	FQ         []*fq.FQ `bson:"fq" json:"fq"`
}

func (s *SP) FQHas(fqCode fq.Code) bool {
	if len(s.FQ) == 0 {
		return false
	}
	for _, f := range s.FQ {
		if f.Fq == fqCode {
			return true
		}
	}
	return false
}
