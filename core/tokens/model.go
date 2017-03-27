package tokens

// High-level description of the cryptographic part of the
// Peernotify protocol

type Pubkey []byte
type Privkey []byte

type IdKey Pubkey

type KeySet interface {
	Pub(int) (Pubkey, error)
	Priv(int) (Pubkey, error)
	Id() IdKey
}

type B58Stringer interface {
	B58String() string
}

type Token struct {
	OneTimeKey Pubkey
	UserIdKey  IdKey
	UserSecret []byte
}

type TokenManager interface {
	NewContactKey() (Privkey, Pubkey, Privkey, Pubkey, error)
	VerifyToken(token *Token) IdKey
}

type PNClient interface {
	NewToken(A, B Pubkey) (*Token, error)
}
