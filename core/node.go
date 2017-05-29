package core

import (
	"github.com/golang/protobuf/proto"
	"github.com/mrwhythat/peernotify/core/tokens"
	"github.com/mrwhythat/peernotify/pb"
	"github.com/mrwhythat/peernotify/storage"
)

type PeernotifyNode struct {
	// TokenManager is an interface for generating/verifying
	// cryptographic tokens
	TokenManager tokens.TokenManager

	// Storage for unconfirmed contacts key-value pairs of
	// the form [<random key> -> <contact data>]. When contact
	// is verified, its data is removed from pending storage and
	// put into contacts storage.
	Pending storage.Storage

	// Permanent contacts storage. Contact data is moved here after
	// confirmation and all lookups are made in this storage.
	Contacts storage.Storage
}

func NewPeernotifyNode(storefile string) (*PeernotifyNode, error) {
	// keys := NewKeyPair()
	// TODO: move into one storage (?)
	tokenManager, err := tokens.NewTokenManager()
	if err != nil {
		return nil, err
	}
	pendingStore, err := storage.NewStorage("/tmp/peernotify.pending")
	if err != nil {
		return nil, err
	}
	contactStore, err := storage.NewStorage(storefile)
	if err != nil {
		return nil, err
	}
	// keyStore, err := storage.NewStorage("/tmp/peernotify.keystore")
	// if err != nil {
	// 	return nil, err
	// }
	return &PeernotifyNode{
		// KeyPair:  keys,
		TokenManager: tokenManager,
		Pending:      pendingStore,
		Contacts:     contactStore,
	}, nil
}

// Store contact data in temporary storage
func (n *PeernotifyNode) registerContact(key []byte, contact *pb.Contact) error {
	return storeContact(n.Pending, key, contact)
}

// Store contact data in permanent storage after verification
func (n *PeernotifyNode) saveContact(key []byte, contact *pb.Contact) error {
	return storeContact(n.Contacts, key, contact)
}

func storeContact(st storage.Storage, key []byte, contact *pb.Contact) error {
	contactBytes, err := proto.Marshal(contact)
	if err != nil {
		return err
	}
	if err := st.Store(key, contactBytes); err != nil {
		return err
	}
	return nil
}

// Get contact data from pending storage
func (n *PeernotifyNode) getPendingContact(key []byte) (*pb.Contact, error) {
	contactBytes, err := n.Pending.Get(key)
	var contact pb.Contact
	proto.Unmarshal(contactBytes, &contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (n *PeernotifyNode) deletePendingContact(key []byte) error {
	return n.Pending.Delete(key)
}

// Get contact data from permanent storage
func (n *PeernotifyNode) getContact(key []byte) (*pb.Contact, error) {
	contactBytes, err := n.Contacts.Get(key)
	if err != nil {
		return nil, err
	}
	var contact pb.Contact
	proto.Unmarshal(contactBytes, &contact)
	return &contact, nil
}

// Generate next public key in HD wallet
func (n *PeernotifyNode) getNextAddress() []byte {
	return []byte{}
}
