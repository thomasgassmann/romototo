package romototo

import "github.com/thomasgassmann/romototo/pkg/romototo/web"

type HousingProvider interface {
	Init(driver *web.BrowserDriver) error
	Refresh() error
	Query() (HousingResult, error)
}
