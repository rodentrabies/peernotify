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

	tc, err := NewPeernotifyClient()
	if err != nil {
		t.Error("Create client")
	}
	idKey, keySet, err := tm.NewKeyset()
	if err != nil {
		t.Error("Generate new keyset")
	}
	token, err := tc.NewToken(keySet)
	if err != nil {
		t.Error("Generate new token")
	}
	generator, err := tm.Generator(token)
	if err != nil {
		t.Error("Get generator")
	}
	if !bytes.Equal(idKey, generator) {
		t.Error("Wrong IDKey")
	}
	_, err = tm.Generator(token)
	if err == nil {
		t.Error("Repeated usage")
	}
	keySet = token
	token, err = tc.NewToken(keySet)
	if err != nil {
		t.Error("Generating 2nd token")
	}
	generator, err = tm.Generator(token)
	if err != nil {
		t.Error("Getting 2nd generator")
	}
}
