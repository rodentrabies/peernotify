package core

type KeyPair struct {
	// privkey []byte
	// pubkey  []byte
}

func NewKeyPair() KeyPair {
	return KeyPair{}
}

type Token struct {
	Key []byte
}
