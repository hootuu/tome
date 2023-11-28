package kt

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/crypto"
	"github.com/hootuu/utils/errors"
	"sort"
)

type Signing struct {
	data map[string]string
}

func NewSigning() *Signing {
	return &Signing{}
}

func (b *Signing) Add(key string, val string) *Signing {
	if b.data == nil {
		b.data = make(map[string]string)
	}
	b.data[key] = val
	return b
}

func (b *Signing) doBuild() string {
	if len(b.data) == 0 {
		return ""
	}
	var keyValuePairs []struct {
		Key   string
		Value string
	}
	for k, v := range b.data {
		keyValuePairs = append(keyValuePairs, struct {
			Key   string
			Value string
		}{k, v})
	}

	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Key < keyValuePairs[j].Key
	})

	content := ""
	for _, pair := range keyValuePairs {
		content += pair.Key + "=" + pair.Value + ";"
	}

	return content
}

type Signature struct {
	Nonce     Nonce     `json:"nonce"`
	Timestamp Timestamp `json:"timestamp"`
	Signer    ki.ADR    `json:"signer"`
	Pub       ki.PUB    `json:"pub"`
	Hash      string    `json:"hash"`
	Signature []byte    `json:"signature"`
}

func NewSignature() *Signature {
	return &Signature{
		Nonce:     NewNonce(),
		Timestamp: NewTimestamp(),
	}
}

func (s *Signature) sign(inv Invariable, privateKey ki.PRI) *errors.Error {
	signer, pub, err := privateKey.GetPUB()
	if err != nil {
		return err
	}
	s.Signer = signer
	s.Pub = pub
	builder := s.doGetSigning(inv)
	signContent := builder.doBuild()
	s.Hash = crypto.SHA256(signContent)
	signatureStr, err := privateKey.Sign(signContent)
	if err != nil {
		return err
	}
	s.Signature = []byte(signatureStr)
	inv.SetSignature(s)
	return nil
}

func (s *Signature) verify(inv Invariable) *errors.Error {
	builder := s.doGetSigning(inv)
	signContent := builder.doBuild()
	ok, err := s.Pub.VerifySignature(signContent, string(s.Signature))
	if err != nil {
		return err
	}
	if !ok {
		return errors.Verify("invalid signature")
	}
	return nil
}

func (s *Signature) doGetSigning(inv Invariable) *Signing {
	signing := inv.Signing()
	signing.Add("type", inv.GetType().S())
	signing.Add("version", inv.GetVersion().S())
	signing.Add("vn", inv.GetVN().S())
	signing.Add("nonce", s.Nonce.S())
	signing.Add("timestamp", s.Timestamp.S())
	signing.Add("signer", s.Signer.S())
	signing.Add("pub", s.Pub.S())
	return signing
}
