package core

import (
	"github.com/yurizhykin/peernotify/crypto"
	"github.com/yurizhykin/peernotify/storage"
)

type PeerNotifyNode struct {
	KeyPair crypto.KeyPair
	Storage storage.Storage
}

type User struct{}

func (n *PeerNotifyNode) Register(user User) error {

}

func (n *PeerNotifyNode) LookUp(token Token) (*User, error) {

}
