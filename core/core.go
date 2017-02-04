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

// Put user data into temporary storage and send verification link
// to the email supplied with data
func (n *PeerNotifyNode) Register(user User) error {

}

// Move user data at 'vid' key in temporary data storage to main storage
func (n *PeerNotifyNode) Verify(vid string) error {

}

// Lookup user by token and forward message via medium marked as desired
func (n *PeerNotifyNode) Forward(token Token, msg []byte) error {

}
