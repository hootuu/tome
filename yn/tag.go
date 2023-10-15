package yn

import (
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Tag string

const (
	NilTag Tag = ""
)

func (t Tag) S() string {
	return string(t)
}

func TagOf(tagStr string) (Tag, *errors.Error) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]{2,60}$", tagStr)
	if !matched {
		return NilTag, errors.Verify("invalid tag: ^[a-zA-Z0-9]{2,60}$")
	}
	return Tag(strings.ToUpper(tagStr)), nil
}

func TagAppend(tags []Tag, tag Tag) []Tag {
	if len(tags) == 0 {
		tags = append(tags, tag)
		return tags
	}
	exists := false
	for _, t := range tags {
		if t == tag {
			exists = true
			break
		}
	}
	if !exists {
		tags = append(tags, tag)
	}
	return tags
}
