package romototo

import (
	"github.com/thomasgassmann/romototo/pkg/romototo/config"
)

type MailNotifier struct {
}

func (t MailNotifier) Init(mail config.MailConfig) error {
	return nil
}

func (t MailNotifier) Send(housing HousingResult) {
	println(housing.Id)
	for _, result := range housing.Results {
		println(result.RoomNumber)
	}
}
