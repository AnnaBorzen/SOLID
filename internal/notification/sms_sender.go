package notification

import (
	"SOLID/internal/repository"
	"fmt"
)

type SMSSender struct {
	apiKey string
	apiURL string
	sender string
}

func NewSMSSender(apiKey, apiURL, sender string) repository.Notifier {
	return &SMSSender{
		apiKey: apiKey,
		apiURL: apiURL,
		sender: sender,
	}
}

// Send Имитация отправки по sms
func (s *SMSSender) Send(customer string) {
	fmt.Printf("Уведомление отправлено клиенту %s по SMS\n", customer)
}
