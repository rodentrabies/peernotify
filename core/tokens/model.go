package tokens

// High-level description of the cryptographic part of the
// Peernotify protocol

type Pubkey []byte
type Privkey []byte

type IdKey Pubkey

// KeySet interface represents the set of keys that is required for a
// certain token generation protocol to function
// type KeySet interface {
// 	Pub(int) (Pubkey, error)
// 	Priv(int) (Pubkey, error)
// 	Id() IdKey
// }

// type ClientKeySet interface {
// 	ClientPubKeys() []Pubkey
// 	ClientPrivKeys() []Privkey
// }

type Keyset struct {
	PrivA Privkey
	PrivB Privkey
	PubA  Pubkey
	PubB  Pubkey
}

// Base58 encoder interface
type B58Stringer interface {
	B58String() string
}

type Token struct {
	OneTimeKey Pubkey
	UserIdKey  IdKey
	UserSecret []byte
}

type TokenManager interface {
	NewContactKey() (*Keyset, error)
	VerifyToken(token *Token) IdKey
}

type PeernotifyClient interface {
	NewToken(A, B Pubkey) (*Token, error)
}
