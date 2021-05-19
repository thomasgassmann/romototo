package notifications

import "github.com/thomasgassmann/robomoto/pkg/romototo"

type Notifier struct {
}

func (t *Notifier) Init() error {
	return nil
}

func (t *Notifier) Send(housing romototo.HousingResult) {
	
}
