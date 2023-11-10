package value

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"regexp"
	"strings"
)

type Key string

func (k Key) S() string {
	return string(k)
}

const (
	NilKey Key = ""
)

const (
	KeyGMV              Key = "GMV"
	KeyOrders           Key = "ORDERS"
	KeyOrderUsers       Key = "ORDER_USERS"
	KeyPCT              Key = "PCT" //per customer transaction
	KeyDailyGMV         Key = "GMV_DAILY"
	KeyDailyOrders      Key = "ORDER_DAILY"
	KeyDailyPCT         Key = "PCT_DAILY"
	KeyDailyActiveUsers Key = "AU_DAILY"
)

func KeyVerify(keyStr string) *errors.Error {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_\\.]{1,108}$", keyStr)
	if !matched {
		return errors.Verify(fmt.Sprintf("invalid value.key: %s", keyStr))
	}
	return nil
}

func KeyOf(keyStr string) (Key, *errors.Error) {
	if err := KeyVerify(keyStr); err != nil {
		return NilKey, err
	}
	return Key(strings.ToUpper(keyStr)), nil
}

type Value map[Key]int64

func NewValue() Value {
	val := make(map[Key]int64)
	return val
}

func (m Value) Add(keyStr string, val int64) *errors.Error {
	key, err := KeyOf(keyStr)
	if err != nil {
		return err
	}
	if val < 0 {
		return errors.Verify("invalid value.value")
	}
	m[key] = val
	return nil
}

func (m Value) Get(key Key) int64 {
	v, ok := m[key]
	if !ok {
		return 0
	}
	return v
}
