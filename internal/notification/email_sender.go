package notification

import (
	"SOLID/internal/repository"
	"fmt"
)

type EmailSender struct {
	smtpHost  string
	smtpPort  int
	fromEmail string
}

func NewEmailSender(host string, port int, from string) repository.Notifier {
	return &EmailSender{
		smtpHost:  host,
		smtpPort:  port,
		fromEmail: from,
	}
}

// Send Имитация отправки по email
func (s *EmailSender) Send(customer string) {
	fmt.Printf("Уведомление отправлено клиенту %s по email \n", customer)
}
