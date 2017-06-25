package tokens

import (
	"bytes"
	"testing"
)

func TestSimpleTokenManager(t *testing.T) {
	tm, err := NewTokenManager("/home/whythat/.peernotify")
	if err != nil {
		t.Error("Create token manager: %v", err)
	}

	idKey, keySet, err := tm.NewKeyset()
	if err != nil {
		t.Error("Generate new keyset")
	}
	tc, err := NewPeernotifyClient(keySet)
	if err != nil {
		t.Error("Create client")
	}
	token, err := tc.NewToken()
	if err != nil {
		t.Error("Generate new token")
	}
	generator, err := tm.Generator(token)
	if err != nil {
		t.Errorf("Get generator: %v", err)
	}
	if !bytes.Equal(idKey, generator) {
		t.Error("Wrong IDKey")
	}
	_, err = tm.Generator(token)
	if err == nil {
		t.Error("Repeated usage")
	}
	keySet = token
	token, err = tc.NewToken()
	if err != nil {
		t.Error("Generating 2nd token")
	}
	generator, err = tm.Generator(token)
	if err != nil {
		t.Errorf("Getting 2nd generator: %v", err)
	}
}
