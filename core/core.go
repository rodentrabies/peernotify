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
	log.Printf("Registering contact \"%s\"\n", contact.Pubkey)
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
	sendVerificationRequest(&contact, idString, url)
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
	log.Printf("Storing contact \"%s\"\n", contact.Pubkey)
	// Create permanent key
	permKey, keyset, err := n.TokenManager.NewKeyset()
	if err != nil {
		return err
	}
	// Store contact data in permanent storage
	if err := n.saveContact(permKey, contact); err != nil {
		return err
	}
	if err := n.deletePendingContact(tmpKey); err != nil {
		return err
	}
	sendContactKeyset(contact, base58.Encode(keyset))
	return nil
}

// Lookup user by token and forward message via medium marked as desired
func (n *PeernotifyNode) Forward(message pb.Message) error {
	// Decode token
	token := base58.Decode(message.Token)
	permKey, err := n.TokenManager.Generator(token)
	if err != nil {
		return err
	}
	contact, err := n.getContact(permKey)
	if err != nil {
		return err
	}
	log.Printf("Forwarding to \"%s\"\n", message.Payload, contact.Pubkey)
	// Forward message
	notifiers.Forward(notifiers.New(contact), message.Payload)
	return nil
}

//------------------------------------------------------------------------------
// Utils

func sendVerificationRequest(contact *pb.Contact, vid, url string) {
	link := url + vid
	aref := fmt.Sprintf("<a href=\"%s\">%s</a>", link, link)
	note := "To confirm this notification method, please visit " + aref
	notifiers.New(contact).Notify("Peernotify Verification", note)
}

func sendContactKeyset(contact *pb.Contact, keyset string) {
	keyform := fmt.Sprintf("Your Peernotify root key is \"%s\".\n", keyset)
	warn := "It is used to generate all your contact tokens. Keep it safe."
	text := keyform + warn
	notifiers.New(contact).Notify("Peernotify Root Key", text)
}

func randBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}
