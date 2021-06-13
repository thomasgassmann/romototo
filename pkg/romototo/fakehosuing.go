package romototo

import (
	"github.com/thomasgassmann/robomoto/pkg/romototo/web"
)

type FakeHousingProvider struct {
}

func (t *FakeHousingProvider) Init(driver *web.BrowserDriver) error {
	return nil
}

func (t *FakeHousingProvider) Refresh() error {
	return nil
}

func (t *FakeHousingProvider) Query() (HousingResult, error) {
	return HousingResult{
		Results:    []Housing{
			{
				RoomNumber: "abc",
			},
		},
		Id: 1,
		Screenshot: []byte{},
	}, nil
}
