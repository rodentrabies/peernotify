package tokens

import "github.com/yurizhykin/monujo"

// High-level description of the cryptographic part of the
// Peernotify protocol

type Pubkey monujo.Pubkey
type Privkey monujo.Privkey

type UserID Pubkey

type Token struct {
	OneTimeKey Pubkey
	UserIdKey  Pubkey
	UserSecret []byte
}

type TokenManager interface {
	NewContactKey() (Privkey, Pubkey, Privkey, Pubkey, error)
	VerifyToken(token *Token) (UserID, error)
}

type PNClient interface {
	NewToken(A, B Pubkey) (*Token, error)
}
