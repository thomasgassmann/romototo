package livingscience

import (
	"errors"
	"hash/fnv"

	"github.com/tebeka/selenium"
	"github.com/thomasgassmann/romototo/pkg/romototo"
	"github.com/thomasgassmann/romototo/pkg/romototo/web"
)

const (
	LivingScienceUrl = "http://reservation.livingscience.ch/wohnen"
	RowSelector      = ".row[class*=status]"
	NumberSelector   = "span.spalte7"
)

type SeleniumLivingScienceHousingProvider struct {
	driver *web.BrowserDriver
}

func InitSeleniumLivingScience() (*SeleniumLivingScienceHousingProvider, error) {
	driver := web.BrowserDriver{}
	if err := driver.Init(); err != nil {
		return nil, err
	}

	err := driver.Driver().Get(LivingScienceUrl)
	if err != nil {
		return nil, err
	}

	return &SeleniumLivingScienceHousingProvider{
		driver: &driver,
	}, nil
}

func (t *SeleniumLivingScienceHousingProvider) Refresh() error {
	return t.driver.Driver().Refresh()
}

func (t *SeleniumLivingScienceHousingProvider) Query() (romototo.HousingResult, error) {
	rows, err := t.driver.Driver().FindElements(selenium.ByCSSSelector, RowSelector)
	if err != nil {
		return romototo.HousingResult{}, err
	}

	if len(rows) == 0 {
		return romototo.HousingResult{}, errors.New("no housings found")
	}

	hash := fnv.New32a()

	var offers []romototo.Housing
	for _, row := range rows {
		roomNumber, err := row.FindElement(selenium.ByCSSSelector, NumberSelector)
		if err != nil {
			continue
		}

		numberText, err := roomNumber.Text()
		if err != nil {
			continue
		}

		hash.Write([]byte(numberText))
		offers = append(offers, romototo.Housing{
			RoomNumber: numberText,
		})
	}

	screenshot, _ := t.driver.Driver().Screenshot()
	return romototo.HousingResult{
		Results:    offers,
		Screenshot: screenshot,
		Id:         hash.Sum32(),
	}, nil
}
