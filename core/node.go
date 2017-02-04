package core

import (
	"github.com/yurizhykin/peernotify/crypto"
	"github.com/yurizhykin/peernotify/storage"
)

type PeernotifyNode struct {
	KeyPair crypto.KeyPair
	Storage storage.Storage
}

func NewPeernotifyNode(storefile string) (*PeernotifyNode, error) {
	keys := crypto.NewKeyPair()
	store, err := storage.NewStorage(storefile)
	if err != nil {
		return nil, err
	}
	return &PeernotifyNode{KeyPair: keys, Storage: store}, nil
}
