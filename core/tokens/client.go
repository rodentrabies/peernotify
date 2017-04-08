package tokens

import "github.com/mrwhythat/monujo"

// cnPNClient is a client adhering to the CryptoNote-based peernotify protocol
type cryptoNoteClient struct{}

func (*cryptoNoteClient) NewToken(A, B Pubkey) (*Token, error) {
	r1, R1, err := monujo.NewKeyPair()
	if err != nil {
		return nil, err
	}
	Au := addPublicKeys(A, R1)
	X := encryptRandPrivkey(A, r1)
	return &Token{Au, B, X}
}

func encryptRandPrivkey(A Pubkey, r1 Privkey) []byte {
	return []byte{}
}

func addPublicKeys(A, B Pubkey) Pubkey {
	return nil
}
