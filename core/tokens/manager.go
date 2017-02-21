package tokens

// TokenManager is a wrapper around Bitcoin wallet, which implement both
// HD wallet functionality and stealth payments.
type TokenManager struct {
}

func NewTokenManager() (*TokenManager, error) {
	return &TokenManager{}, nil
}
