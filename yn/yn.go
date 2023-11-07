package yn

import (
	"fmt"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type YID = string

type Yin struct {
	Yin   Code     `bson:"yin" json:"yin"`
	Vn    vn.ID    `bson:"vn" json:"vn" cbor:"vn"`
	Sp    sp.ID    `bson:"sp" json:"sp" cbor:"sp"`
	Who   *Who     `bson:"who" json:"who"`
	When  When     `bson:"when" json:"when"`
	Act   Act      `bson:"act" json:"act"`
	What  *What    `bson:"what" json:"what"`
	Exp   *Expense `bson:"exp" json:"exp"`
	Title Title    `bson:"title" json:"title"`
	Tag   []Tag    `bson:"tag" json:"tag"`
	Ctx   Ctx      `bson:"ctx" json:"ctx"`
}

func (y *Yin) GetType() string {
	return "YIN"
}

func (y *Yin) GetVN() vn.ID {
	return y.Vn
}

func (y *Yin) GetSP() sp.ID {
	return y.Sp
}

func (y *Yin) Summary() string {
	return fmt.Sprintf("%s %s Exp[%s%d]", y.Who.Ref.Ref, y.Act, y.Exp.Currency, y.Exp.Amount)
}

func NewYin(vnIDStr string, spIDStr string) (*Yin, *errors.Error) {
	vnID, err := vn.IDOf(vnIDStr)
	if err != nil {
		return nil, errors.Verify("invalid vn: " + err.Error())
	}
	spID, err := sp.IDOf(spIDStr)
	if err != nil {
		return nil, errors.Verify("invalid sp: " + err.Error())
	}
	return &Yin{
		Vn:   vnID,
		Sp:   spID,
		When: NewWhen(),
		Tag:  []Tag{},
		Ctx:  NewCtx(),
	}, nil
}

func (y *Yin) WithWho(whoStr string, refCode string, refRef string) *errors.Error {
	who, err := NewWho(whoStr, refCode, refRef)
	if err != nil {
		return errors.Verify("invalid yin.who: " + err.Error())
	}
	y.Who = who
	return nil
}

func (y *Yin) WithAct(actStr string) *errors.Error {
	act, err := ActOf(actStr)
	if err != nil {
		return errors.Verify("invalid yin.act: " + err.Error())
	}
	y.Act = act
	return nil
}

func (y *Yin) WithWhat(codeStr string, refStr string) *errors.Error {
	what, err := NewWhat(codeStr, refStr)
	if err != nil {
		return errors.Verify("invalid yin.what: " + err.Error())
	}
	y.What = what
	return nil
}

func (y *Yin) WithExpense(currencyStr string, amount int64) *errors.Error {
	exp, err := NewExpense(currencyStr, amount)
	if err != nil {
		return errors.Verify("invalid yin.expense: " + err.Error())
	}
	y.Exp = exp
	return nil
}

func (y *Yin) WithTitle(tilStr string) *errors.Error {
	til, err := TitleOf(tilStr)
	if err != nil {
		return errors.Verify("invalid yin.title: " + err.Error())
	}
	y.Title = til
	return nil
}

func (y *Yin) WithTag(tagStrArr ...string) *errors.Error {
	if len(tagStrArr) == 0 {
		return nil
	}
	for _, tStr := range tagStrArr {
		t, err := TagOf(tStr)
		if err != nil {
			return errors.Verify("invalid tag: " + err.Error())
		}
		y.Tag = TagAppend(y.Tag, t)
	}
	return nil
}

func (y *Yin) WithCtx(codeStr string, valStr string) *errors.Error {
	err := y.Ctx.Put(codeStr, valStr)
	if err != nil {
		return errors.Verify("invalid ctx: " + err.Error())
	}
	return nil
}

func (y *Yin) Verify() *errors.Error {
	if len(y.Vn) == 0 {
		return errors.Verify("require yin.vn")
	}
	if len(y.Sp) == 0 {
		return errors.Verify("require yin.sp")
	}
	if y.Who == nil {
		return errors.Verify("require yin.who")
	}
	if len(y.Act) == 0 {
		return errors.Verify("require yin.act")
	}
	if len(y.Title) == 0 {
		return errors.Verify("require yin.title")
	}
	if y.What == nil {
		return errors.Verify("require yin.what")
	}
	if y.Exp == nil {
		return errors.Verify("require yin.exp")
	}
	if y.Exp.Ex == nil {
		y.Exp.PutEx("exp", &ExpenseItem{
			Amount:   y.Exp.Amount,
			Currency: y.Exp.Currency,
		})
	}
	return nil
}

func (y *Yin) Digest() string {
	return fmt.Sprintf("[%s:%s][%s] %s %s %s EXP %d #%s#",
		y.Vn.S(),
		y.Sp.S(),
		y.When.Time().Format("2006-01-02 15:04:05"),
		y.Who.Who,
		y.Act.S(),
		y.What.Code,
		y.Exp.Amount,
		y.Title)
}
