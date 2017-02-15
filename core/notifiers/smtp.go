package notifiers

import "net/smtp"

type smtpConfig struct {
	Host     string
	Port     string
	Addr     string
	Username string
	Password string
}

type smtpNotifier struct {
	Conf  smtpConfig
	Email string
}

func (n *smtpNotifier) Notify(subj, msg string) {
	auth := smtp.PlainAuth("", n.Conf.Username, n.Conf.Password, n.Conf.Host)
	addr := n.Conf.Host + ":" + n.Conf.Port
	head := "From: " + n.Conf.Addr + "\r\n" +
		"To: " + n.Email + "\r\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"Subject: " + subj + "\r\n\r\n" +
		msg + "\r\n"
	err := smtp.SendMail(addr, auth, n.Conf.Addr, []string{n.Email}, []byte(head))
	if err != nil {
		panic(err.Error())
	}
}
