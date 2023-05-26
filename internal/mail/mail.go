package mail

import (
	"net/smtp"

	"github.com/andrew-pisotskyi/genesis-task/domain"
)

type Mail struct {
	from     string
	password string
	host     string
	port     string
}

func NewMail(host, port, from, password string) domain.MailInterface {
	return &Mail{from, password, host, port}
}

func (s *Mail) Send(toEmailAddress []string, body string) error {
	address := s.host + ":" + s.port
	message := []byte(body)

	auth := smtp.PlainAuth("", s.from, s.password, s.host)
	return smtp.SendMail(address, auth, s.from, toEmailAddress, message)
}
