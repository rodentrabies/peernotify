package notifiers

import "github.com/yurizhykin/peernotify/pb"

type Notifier interface {
	Forward([]byte)
}

type notifierList struct {
	list []Notifier
}

func New(contact *pb.Contact) Notifier {
	smtpConf := smtpConfig{
		Host:     "localhost",
		Port:     "25",
		Addr:     "test@test.org",
		Username: "user",
		Password: "password",
	}
	return &notifierList{
		[]Notifier{
			// SMTP notifier
			&smtpNotifier{smtpConf, contact.Email.Address},
			// TODO: add more later
		},
	}
}

func (nlist *notifierList) Forward(msg []byte) {
	for _, notif := range nlist.list {
		notif.Forward(msg)
	}
}
