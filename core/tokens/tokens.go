package tokens

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"
	"path"

	base58 "github.com/jbenet/go-base58"
	"github.com/mrwhythat/peernotify/storage"
)

//------------------------------------------------------------------------------
// Model
//
// High-level description of the cryptographic part of the Peernotify protocol

// type B58Stringer interface {
// 	B58String() string
// }

func b58String(data []byte) string {
	return base58.Encode(data)
}

type Token struct {
	OneTimeKey []byte
	UserIdKey  []byte
	UserSecret []byte
}

type TokenManager interface {
	// Returns binary representation of the set of keys needed for
	// Peernotify protocol and ID key, which can be used to index keyset
	// in storage
	NewKeyset() ([]byte, []byte, error)

	// Returns ID key of the keyset that generated given token
	Generator(tokenBytes []byte) ([]byte, error)
}

type PeernotifyClient interface {
	NewToken(keyset []byte) ([]byte, error)
}

const TokenSize = 32

var (
	IDSize    = 32
	MaxTokens = math.MaxInt16

	RandError           = errors.New("Ur randomz numbez iz broken, mate!")
	IncorrectTokenError = errors.New("Incorrect token")
)

//------------------------------------------------------------------------------
// Simple token manager implementation
type simpleTokenManager struct {
	keys   storage.Storage
	tokens storage.Storage
}

func NewTokenManager(storeDir string) (TokenManager, error) {
	keys, err := storage.NewStorage(path.Join(storeDir, "peernotify.keys"))
	if err != nil {
		return nil, err
	}
	tokens, err := storage.NewStorage(path.Join(storeDir, "peernotify.tokens"))
	if err != nil {
		return nil, err
	}
	return &simpleTokenManager{keys: keys, tokens: tokens}, nil
}

func (tm *simpleTokenManager) NewKeyset() ([]byte, []byte, error) {
	idKey := make([]byte, IDSize)
	if _, err := rand.Read(idKey); err != nil {
		return nil, nil, RandError
	}
	rootKey := make([]byte, TokenSize)
	if _, err := rand.Read(rootKey); err != nil {
		return nil, nil, RandError
	}
	if err := tm.keys.Store(idKey, rootKey); err != nil {
		return nil, nil, err
	}
	return idKey, append(idKey, rootKey...), nil
}

func (tm *simpleTokenManager) Generator(tokenBytes []byte) ([]byte, error) {
	idKey, token := tokenBytes[:IDSize], tokenBytes[IDSize:]
	rootKey, err := tm.keys.Get(idKey)
	if err != nil {
		return nil, err
	}
	var chainLink [TokenSize]byte
	copy(chainLink[:], rootKey)
	rawTokenSet, err := tm.tokens.Get(idKey)
	if err != nil {
		return nil, err
	}
	var found bool
	tokenSet := NewTokenSet(rawTokenSet)
	for i := 0; i < MaxTokens; i++ {
		chainLink = sha256.Sum256(chainLink[:])
		if bytes.Equal(chainLink[:], token) {
			if tokenSet.GetAt(i) {
				return nil, IncorrectTokenError
			}
			tokenSet = tokenSet.AddAt(i)
			found = true
			break
		}
	}
	copy(chainLink[:], rootKey)
	var i int
	for i = 0; tokenSet.GetAt(i); i++ {
		chainLink = sha256.Sum256(chainLink[:])
	}
	tm.keys.Store(idKey, chainLink[:])
	tm.tokens.Store(idKey, tokenSet.DropUntil(i))
	if found {
		return idKey, nil
	} else {
		return nil, IncorrectTokenError
	}
}

//------------------------------------------------------------------------------
// Simple client implementation

type simpleClient struct{}

func NewPeernotifyClient() (PeernotifyClient, error) {
	return &simpleClient{}, nil
}

func (*simpleClient) NewToken(keyset []byte) ([]byte, error) {
	idKey, rootKey := keyset[:IDSize], keyset[IDSize:]
	chainLink := sha256.Sum256(rootKey)
	return append(idKey, chainLink[:]...), nil
}
