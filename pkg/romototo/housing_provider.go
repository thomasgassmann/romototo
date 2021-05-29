package romototo

import "github.com/thomasgassmann/robomoto/pkg/romototo/web"

type HousingProvider interface {
	Init(driver *web.BrowserDriver) error
	Refresh() error
	Query() (HousingResult, error)
}
