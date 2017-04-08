package core

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
	"github.com/mrwhythat/peernotify/core/notifiers"
	"github.com/mrwhythat/peernotify/pb"
)

//------------------------------------------------------------------------------
// Main operations

// Put user data into temporary storage and send verification link
// to the email supplied with data
func (n *PeernotifyNode) Register(contact pb.Contact, url string) error {
	log.Printf(
		"Registering contact \"%s\" (mail: %s)",
		contact.Pubkey,
		contact.Email.Address,
	)
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
	idString := base58.Encode(tmpKey)
	// Send verification request
	if err := sendVerificationRequest(contact, idString, url); err != nil {
		return err
	}
	return nil
}

// Move user data at 'vid' key in temporary data storage to main storage
func (n *PeernotifyNode) Verify(vid string) error {
	log.Printf("Verifying contact randID \"%s\"\n", vid)
	// Decode key string
	tmpKey := base58.Decode(vid)
	// Lookup contact by ID key
	contact, err := n.getPendingContact(tmpKey)
	if err != nil {
		return err
	}
	log.Printf(
		"Storing contact \"%s\" (mail: %s)",
		contact.Pubkey,
		contact.Email.Address,
	)
	// Create permanent key
	// NOTE: for testing purposes, currently use Contact.Pubkey
	permKey := []byte(contact.Pubkey)
	// Store contact data in permanent storage
	if err := n.saveContact(permKey, contact); err != nil {
		return err
	}
	return nil
}

// Lookup user by token and forward message via medium marked as desired
func (n *PeernotifyNode) Forward(message pb.Message) error {
	// Decode token
	token := []byte(message.Token)
	// if err != nil {
	// 	return err
	// }
	// Get contact from storage
	contact, err := n.getContact(token)
	if err != nil {
		return err
	}
	log.Printf(
		"Forwarding message \"%s\" to contact \"%s\"\n",
		message.Payload,
		contact.Pubkey,
	)
	// Create notifiers
	notif := notifiers.New(contact)
	// Forward message
	notifiers.Forward(notif, message.Payload)
	return nil
}

//------------------------------------------------------------------------------
// Utils

func sendVerificationRequest(contact pb.Contact, vid, url string) error {
	notifier := notifiers.New(&contact)
	link := url + vid
	aref := fmt.Sprintf("<a href=\"%s\">%s</a>", link, link)
	note := "To confirm this notification method, please visit " + aref
	notifier.Notify("Peernotify Verification", note)
	return nil
}

func randBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
