package tokens

import "github.com/yurizhykin/monujo"

// cnTokenManager is a wrapper around CryptoNight cryptocurrency wallet
type cnTokenManager struct {
	wallet monujo.Wallet
}

func NewTokenManager() (*TokenManager, error) {
	wallet, err := monujo.NewWallet()
	wallet.Run(nil)
	return &cnTokenManager{wallet: wallet}, nil
}

func (tm *cnTokenManager) NewContactKey() (Privkey, Pubkey, Privkey, Pubkey, error) {
	a, A, err := monujo.NewKeyPair()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	b, B, err := monujo.NewKeyPair()
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return a, A, b, B, nil
}

func (tm *cnTokenManager) VerifyToken(token *Token) IdKey {
	return token.UserIdKey
}
