package tokens

import "github.com/yurizhykin/monujo"

// cnPNClient is a client adhering to the cryptonight-based peernotify protocol
type cnPNClient struct{}

func (*cnPNClient) NewToken(A, B Pubkey) (*Token, error) {
	r1, R1, err := monujo.NewKeyPair()
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
