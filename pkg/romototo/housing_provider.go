package romototo

type HousingProvider interface {
	Refresh() error
	Query() (HousingResult, error)
}
