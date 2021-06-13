package romototo

type Notifier struct {
}

func (t *Notifier) Init() error {
	return nil
}

func (t *Notifier) Send(housing HousingResult) {
	for _, result := range housing.Results {
		println(result.RoomNumber)
	}
}
