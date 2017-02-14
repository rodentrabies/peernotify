package core

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/yurizhykin/peernotify/core/notifiers"
	"github.com/yurizhykin/peernotify/pb"
)

//------------------------------------------------------------------------------
// Main operations

// Put user data into temporary storage and send verification link
// to the email supplied with data
func (n *PeernotifyNode) Register(contact pb.Contact) error {
	log.Printf("Registering %+v", contact)
	// Generate random secure key
	tmpKey, err := randBytes(32)
	if err != nil {
		return err
	}
	// Save to temporary storage
	if err := n.registerContact(tmpKey, &contact); err != nil {
		return err
	}
	// Encode random key to URL Base 64 string
	idString := base64.URLEncoding.EncodeToString(tmpKey)
	// Send verification request
	if err := sendVerificationRequest(contact, idString); err != nil {
		return err
	}
	return nil
}

// Move user data at 'vid' key in temporary data storage to main storage
func (n *PeernotifyNode) Verify(vid string) error {
	log.Printf("Verified contact %+v", vid)
	// Decode key string
	tmpKey, err := base64.URLEncoding.DecodeString(vid)
	if err != nil {
		return err
	}
	// Lookup contact by ID key
	contact, err := n.getPendingContact(tmpKey)
	if err != nil {
		return err
	}
	// Create permanent key
	permKey := make([]byte, 32)
	// Store contact data in permanent storage
	if err := n.saveContact(permKey, contact); err != nil {
		return err
	}
	return nil
}

// Lookup user by token and forward message via medium marked as desired
func (n *PeernotifyNode) Forward(message pb.Message) error {
	// Decode token
	token, err := base64.StdEncoding.DecodeString(message.Token)
	if err != nil {
		return err
	}
	// Get contact from storage
	contact, err := n.getContact(token)
	if err != nil {
		return err
	}
	log.Printf("Forwarding message %+v to contact %+v", message, contact)
	// Create notifiers
	notif := notifiers.New(contact)
	// Decode message payload
	payload, err := base64.StdEncoding.DecodeString(message.Payload)
	if err != nil {
		return err
	}
	// Forward message
	notif.Forward(payload)
	return nil
}

//------------------------------------------------------------------------------
// Utils

func sendVerificationRequest(contact pb.Contact, vid string) error {
	return nil
}

func randBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
