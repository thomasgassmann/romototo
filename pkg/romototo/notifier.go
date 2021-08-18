package romototo

type Notifier interface {
	Send(housing HousingResult)
}
