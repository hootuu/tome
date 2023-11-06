package bk

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/crypto"
	"github.com/hootuu/utils/errors"
	"sort"
)

type SignBuilder struct {
	data map[string]string
}

func NewSignBuilder() *SignBuilder {
	return &SignBuilder{}
}

func (b *SignBuilder) Add(key string, val string) *SignBuilder {
	if b.data == nil {
		b.data = make(map[string]string)
	}
	b.data[key] = val
	return b
}

func (b *SignBuilder) Build() string {
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

func (b *SignBuilder) Hash() string {
	content := b.Build()
	return crypto.SHA256(content)
}

func (b *SignBuilder) Sign(privateKey ki.PRI) (string, *errors.Error) {
	content := b.Build()
	return privateKey.Sign(content)
}

func (b *SignBuilder) Verify(publicKey ki.PUB, signStr string) (bool, *errors.Error) {
	content := b.Build()
	return publicKey.VerifySignature(content, signStr)
}
