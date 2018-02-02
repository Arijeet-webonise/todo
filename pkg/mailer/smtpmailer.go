package mailer

import (
	"net/smtp"
)

type mailer struct {
	Client   string
	username string
	password string
}

type mailer_interface interface {
	SendMail()
	testSMTP()
	Auth()
}

func (m *mailer) Auth() (auth smtp.Auth) {
	return smtp.PlainAuth("", m.username, m.password, m.Client)
}

func (m *mailer) testSMTP(to []string) (err error) {
	return m.SendMail("test@local.com", to, "Test is a test")
}

func (m *mailer) SendMail(from string, to []string, msg string) (err error) {
	auth := m.Auth()

	return smtp.SendMail(m.Client, auth, from, to, []byte(msg))

}
