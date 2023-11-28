package nd

import (
	"fmt"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/sys"
)

type ID string

func (id ID) S() string {
	return string(id)
}

type Node struct {
	ID  ID
	ADR ki.ADR
	PUB ki.PUB
	PRI ki.PRI
}

func NewNode(id ID, pri ki.PRI) (*Node, *errors.Error) {
	n := &Node{
		ID:  id,
		ADR: "",
		PUB: "",
		PRI: pri,
	}
	adr, pub, err := pri.GetPUB()
	if err != nil {
		return nil, err
	}
	n.ADR = adr
	n.PUB = pub
	return n, nil
}

func (n *Node) Lead() *Lead {
	return &Lead{
		ID:  n.ID,
		ADR: n.ADR,
		PUB: n.PUB,
	}
}

type Lead struct {
	ID  ID
	ADR ki.ADR
	PUB ki.PUB
}

func (l *Lead) S() string {
	return fmt.Sprintf("%s:%s:%s", l.ID.S(), l.ADR.S(), l.PUB.S())
}

var gHere *Node

func Init(id ID, pri ki.PRI) {
	n, err := NewNode(id, pri)
	if err != nil {
		sys.Exit(err)
		return
	}
	gHere = n
}

func Here() *Node {
	if gHere == nil {
		sys.Error("do nd.Init first ...")
	}
	return gHere
}
