package tokens

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"
	"path"

	"github.com/mrwhythat/peernotify/storage"
)

//------------------------------------------------------------------------------
// Model
type Token struct {
	OneTimeKey []byte
	UserIdKey  []byte
	UserSecret []byte
}

type TokenManager interface {
	// Returns binary representation of the set of keys and ID key
	// which can be used to index keyset in storage
	NewKeyset() ([]byte, []byte, error)

	// Returns ID key of the keyset that generated given token
	Generator(tokenBytes []byte) ([]byte, error)
}

type PeernotifyClient interface {
	// Generate new token based on the keyset
	NewToken() ([]byte, error)
}

const (
	IDSize    = 8
	KeySize   = 32
	MaskSize  = 32
	TokenSize = 40
	MaxTokens = math.MaxInt16
)

var (
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
	id := make([]byte, IDSize)
	if _, err := rand.Read(id); err != nil {
		return nil, nil, RandError
	}
	key := make([]byte, KeySize)
	if _, err := rand.Read(key); err != nil {
		return nil, nil, RandError
	}
	mask := make([]byte, MaskSize)
	if _, err := rand.Read(mask); err != nil {
		return nil, nil, RandError
	}
	rootKeyset := append(key, mask...)
	if err := tm.keys.Store(id, rootKeyset); err != nil {
		return nil, nil, err
	}
	return id, append(id, rootKeyset...), nil
}

func (tm *simpleTokenManager) Generator(tokenBytes []byte) ([]byte, error) {
	// Break raw token bytes into ID and token data
	id, token := tokenBytes[:IDSize], tokenBytes[IDSize:]
	// 	fmt.Printf(`
	//      vid: %s
	// `, base58.Encode(id))
	// Get root keyset for ID
	rootKeyset, err := tm.keys.Get(id)
	if err != nil {
		return nil, err
	}
	// Get token set for ID
	rawTokenSet, err := tm.tokens.Get(id)
	if err != nil {
		return nil, err
	}
	tokenSet := NewTokenSet(rawTokenSet)
	// Break root keyset into pieces
	var (
		key  [KeySize]byte
		mask [MaskSize]byte
		link [MaskSize]byte
	)
	// Check token
	copy(key[:], rootKeyset[:KeySize])
	copy(mask[:], rootKeyset[KeySize:])
	// 	fmt.Printf(`
	//     vkey: %s
	//    vmask: %s
	// `, base58.Encode(key[:]), base58.Encode(mask[:]))
	xorBytes(link[:], key[:], mask[:])
	var found bool
	for i := 0; i < MaxTokens; i++ {
		key = sha256.Sum256(key[:])
		mask = sha256.Sum256(mask[:])
		xorBytes(link[:], key[:], mask[:])
		if bytes.Equal(link[:], token) {
			if tokenSet.GetAt(i) {
				return nil, IncorrectTokenError
			}
			tokenSet = tokenSet.AddAt(i)
			found = true
			break
		}
	}
	// Remove used tokens from token set
	copy(key[:], rootKeyset[:KeySize])
	copy(mask[:], rootKeyset[KeySize:])
	xorBytes(link[:], key[:], mask[:])
	var i int
	for i = 0; tokenSet.GetAt(i); i++ {
		key = sha256.Sum256(key[:])
		mask = sha256.Sum256(mask[:])
		xorBytes(link[:], key[:], mask[:])
	}
	// Save new root keyset and tokenset
	tm.keys.Store(id, append(key[:], mask[:]...))
	tm.tokens.Store(id, tokenSet.DropUntil(i))
	// Return ID of the generator
	if found {
		return id, nil
	} else {
		return nil, IncorrectTokenError
	}
}

func xorBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

//------------------------------------------------------------------------------
// Simple client implementation

type simpleClient struct {
	keyset []byte
}

func NewPeernotifyClient(keyset []byte) (PeernotifyClient, error) {
	return &simpleClient{keyset: keyset}, nil
}

func (client *simpleClient) NewToken() ([]byte, error) {
	kSet := client.keyset
	keyEnd := IDSize + KeySize
	id, key, mask := kSet[:IDSize], kSet[IDSize:keyEnd], kSet[keyEnd:]
	// 	fmt.Printf(`
	//       id: %s
	//      key: %s
	//     mask: %s
	// `, base58.Encode(id), base58.Encode(key), base58.Encode(mask))
	newKey, newMask := sha256.Sum256(key), sha256.Sum256(mask)
	// 	fmt.Printf(`
	//  new key: %s
	// new mask: %s
	// `, base58.Encode(newKey[:]), base58.Encode(newMask[:]))
	copy(client.keyset[IDSize:keyEnd], newKey[:])
	copy(client.keyset[keyEnd:], newMask[:])
	token := make([]byte, TokenSize)
	copy(token[:IDSize], id)
	xorBytes(token[IDSize:], newKey[:], newMask[:])
	return token, nil
}
