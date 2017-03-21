package tokens

import (
	arbo "github.com/yurizhykin/arbo/keys"
)

func NewContactKey() (arbo.Privkey, arbo.Pubkey, arbo.Privkey, arbo.Pubkey, error) {
	a, A, err := arbo.NewKeyPair()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	b, B, err := arbo.NewKeyPair()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return a, A, b, B, nil
}

type Token struct {
	Au arbo.Pubkey
	B  arbo.Pubkey
	X  []byte
}

func NewToken(A, B arbo.Pubkey) (*Token, error) {
	r1, R1, err := arbo.NewKeyPair()
	if err != nil {
		return nil, err
	}
	Au := PubAdd(A, R1)
	X := encryptRandPrivkey(A, r1)
	return &Token{Au, B, X}
}

func encryptRandPrivkey(A arbo.Pubkey, r1 arbo.Privkey) []byte {
	return []byte{}
}
