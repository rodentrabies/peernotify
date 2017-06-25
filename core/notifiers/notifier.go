package notifiers

import "github.com/mrwhythat/peernotify/pb"

type Notifier interface {
	Notify(subj, msg string)
}

type notifierList struct {
	list []Notifier
}

func New(contact *pb.Contact) Notifier {
	// Configure SMTP
	smtpConf := smtpConfig{
		Host:     "0.0.0.0",
		Port:     "1025",
		Addr:     "test@test.org",
		Username: "user",
		Password: "password",
	}
	notifiers := []Notifier{}
	for _, method := range contact.Methods {
		var notifier Notifier
		switch method.Protocol {
		case "SMTP":
			notifier = &smtpNotifier{smtpConf, method.Address}
		}
		notifiers = append(notifiers, notifier)
	}
	return &notifierList{notifiers}
}

func (nlist *notifierList) Notify(subj, msg string) {
	for _, notif := range nlist.list {
		notif.Notify(subj, msg)
	}
}

func Forward(notif Notifier, msg string) {
	note := "Forwarding message:\n\n"
	notif.Notify("Peernotify Forward", note+msg)
}
