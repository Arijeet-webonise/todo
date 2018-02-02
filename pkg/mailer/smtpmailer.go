package mailer

import (
	"net/smtp"
	"strconv"
)

type SMTPMailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

type ISMTPMailer interface {
	SendMail()
	TestSMTP()
	Auth()
}

func (m *SMTPMailer) Auth() (auth smtp.Auth) {
	return smtp.PlainAuth("", m.Username, m.Password, m.Host)
}

func (m *SMTPMailer) TestSMTP(to []string) (err error) {
	return m.SendMail("test@local.com", to, "Test is a test")
}

func (m *SMTPMailer) SendMail(from string, to []string, msg string) (err error) {
	auth := m.Auth()
	return smtp.SendMail(m.Host+":"+strconv.Itoa(m.Port), auth, from, to, []byte(msg))

}
