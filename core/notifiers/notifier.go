package notifiers

import "github.com/yurizhykin/peernotify/pb"

type Notifier interface {
	Forward([]byte)
}

type notifierList struct {
	list []Notifier
}

func New(contact pb.Contact) *Notifier {
	return []Notifier{contact.Email}
}

func (nlist *notifierList) Forward(msg []byte) {
	for _, notif := range nlist.list {
		notif.Forward(msg)
	}
}
