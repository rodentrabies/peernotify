package core

import (
	"github.com/golang/protobuf/proto"
	"github.com/yurizhykin/peernotify/pb"
	"github.com/yurizhykin/peernotify/storage"
)

type PeernotifyNode struct {
	KeyPair KeyPair
	Storage storage.Storage
}

func NewPeernotifyNode(storefile string) (*PeernotifyNode, error) {
	keys := NewKeyPair()
	store, err := storage.NewStorage(storefile)
	if err != nil {
		return nil, err
	}
	return &PeernotifyNode{KeyPair: keys, Storage: store}, nil
}

func (n *PeernotifyNode) storeContact(contact *pb.Contact) ([]byte, error) {
	contactBytes, err := proto.Marshal(contact)
	contactKey := ContactKey(contactBytes)
	if err := n.Storage.Store(contactKey, contactBytes); err != nil {
		return nil, err
	}
	return contactKey, nil
}

func (n *PeernotifyNode) getContact(key []byte) (*pb.Contact, error) {
	contactBytes, err := n.Storage.Get(key)
	var contact pb.Contact
	proto.Unmarshal(contactBytes, &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil

}
