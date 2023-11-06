package pr

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/errors"
	"sync"
)

type Peer struct {
	ID  string `json:"i"`
	PUB ki.PUB `json:"p"`
}

type Local struct {
	ID   string
	PRI  ki.PRI
	peer *Peer
	lock sync.Mutex
}

func (l *Local) Peer() (*Peer, *errors.Error) {
	if l.peer != nil {
		return l.peer, nil
	}
	l.lock.Lock()
	defer l.lock.Unlock()
	l.peer = &Peer{
		ID: l.ID,
	}
	var err *errors.Error
	_, l.peer.PUB, err = l.PRI.GetPUB()
	if err != nil {
		return nil, err
	}
	return l.peer, nil
}
