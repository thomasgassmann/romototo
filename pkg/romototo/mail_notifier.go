package romototo

type MailNotifier struct {
}

func (t MailNotifier) Init() error {
	return nil
}

func (t MailNotifier) Send(housing HousingResult) {
	println(housing.Id)
	for _, result := range housing.Results {
		println(result.RoomNumber)
	}
}
