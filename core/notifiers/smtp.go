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

func (n *smtpNotifier) Forward(msg []byte) {
	auth := smtp.PlainAuth("", n.Conf.Username, n.Conf.Password, n.Conf.Host)
	hostAddr := n.Conf.Host + ":" + n.Conf.Port
	err := smtp.SendMail(hostAddr, auth, n.Conf.Addr, []string{n.Email}, msg)
	if err != nil {
		panic(err.Error())
	}
}
