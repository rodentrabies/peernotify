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

func (n *smtpNotifier) Notify(msg string) {
	auth := smtp.PlainAuth("", n.Conf.Username, n.Conf.Password, n.Conf.Host)
	addr := n.Conf.Host + ":" + n.Conf.Port
	err := smtp.SendMail(addr, auth, n.Conf.Addr, []string{n.Email}, []byte(msg))
	if err != nil {
		panic(err.Error())
	}
}
