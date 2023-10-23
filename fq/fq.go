package fq

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Code string

const (
	NilCode Code = ""

	SLIVER  Code = "SLIVER"
	GOLD    Code = "GOLD"
	DIAMOND Code = "DIAMOND"
)

func (c Code) S() string {
	return string(c)
}

func CodeVerify(codeStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{1,360}$", codeStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid FQ code: %s", codeStr))
	}
	return nil
}

func CodeOf(codeStr string) (Code, *errors.Error) {
	if err := CodeVerify(codeStr); err != nil {
		return NilCode, err
	}
	return Code(strings.ToUpper(codeStr)), nil
}

//type Meta struct {
//	Description string `bson:"description" json:"description"`
//	Image       string `bson:"image" json:"image"`
//	ExternalUrl string `bson:"external_url" json:"external_url"`
//	License     string `bson:"license" json:"license"`
//}

type FQ struct {
	Fq Code `bson:"fq" json:"fq"`
	//Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
}

func (f *FQ) Verify() *errors.Error {
	if err := CodeVerify(f.Fq.S()); err != nil {
		return err
	}
	return nil
}

func Include(arr []*FQ, fq Code) bool {
	if len(arr) == 0 {
		return false
	}
	for _, f := range arr {
		if strings.ToUpper(f.Fq.S()) == strings.ToUpper(fq.S()) {
			return true
		}
	}
	return false
}
