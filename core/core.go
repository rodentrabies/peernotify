package core

import (
	"log"

	"github.com/yurizhykin/peernotify/pb"
)

// Put user data into temporary storage and send verification link
// to the email supplied with data
func (n *PeernotifyNode) Register(contact pb.Contact) error {
	log.Printf("Registering %+v", contact)
	_, err := n.storeContact(&contact)
	return err
}

// Move user data at 'vid' key in temporary data storage to main storage
func (n *PeernotifyNode) Verify(vid string) error {
	return nil
}

// Lookup user by token and forward message via medium marked as desired
func (n *PeernotifyNode) Forward(token Token, msg []byte) error {
	contact, err := n.getContact(token.Key)
	if err != nil {
		return err
	}
	notif := notifiers.New(contact)
	return notif.Forward(msg)
}
