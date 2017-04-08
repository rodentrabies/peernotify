package tokens

import "github.com/mrwhythat/monujo"

// cnTokenManager is a wrapper around CryptoNight cryptocurrency wallet
type cnTokenManager struct {
	wallet monujo.Wallet
}

func NewTokenManager() (*TokenManager, error) {
	wallet, err := monujo.NewWallet()
	wallet.Run(nil)
	return &cnTokenManager{wallet: wallet}, nil
}

func (tm *cnTokenManager) NewContactKey() (*Keyset, error) {
	a, A, err := monujo.NewKeyPair()
	if err != nil {
		return nil, err
	}
	b, B, err := monujo.NewKeyPair()
	if err != nil {
		return nil, err
	}
	return &Keyset{a, b, A, B}, nil
}

func (tm *cnTokenManager) VerifyToken(token *Token) IdKey {
	return token.UserIdKey
}
