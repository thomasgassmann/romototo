package romototo

import (
	"github.com/tebeka/selenium"
	"github.com/thomasgassmann/robomoto/pkg/romototo/web"
)

const (
	LivingScienceUrl = "http://reservation.livingscience.ch/wohnen"
	RowSelector = ".row[class*=status]"
	NumberSelector = "span.spalte7"
)

type LivingScienceHousingProvider struct  {
	driver *web.BrowserDriver
}

func (t *LivingScienceHousingProvider) Init(driver *web.BrowserDriver) {
	t.driver = driver
	t.driver.Driver().Get(LivingScienceUrl)
}

func (t *LivingScienceHousingProvider) Refresh() error {
	t.driver.Driver().Refresh()
	return nil
}

func (t *LivingScienceHousingProvider) Query() (HousingResult, error) {
	rows, err := t.driver.Driver().FindElements(selenium.ByCSSSelector, RowSelector)
	if err != nil {
		return HousingResult{}, err
	}

	var offers []Housing
	for _, row := range rows {
		roomNumber, err := row.FindElement(selenium.ByCSSSelector, NumberSelector)
		if err != nil {
			continue
		}

		numberText, err := roomNumber.Text()
		if err != nil {
			continue
		}

		offers = append(offers, Housing{
			RoomNumber: numberText,
		})
	}

	return HousingResult{
		Results:    offers,
		Screenshot: nil,
	}, nil
}
