package domain

type MailInterface interface {
	Send(toEmailAddress []string, body string) error
}
