package romototo

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type RomototoDriver struct {
	service *selenium.Service
	webDriver *selenium.WebDriver
}

const (
	port = 8080
	seleniumPath = ""
)

func (t RomototoDriver) Init() error {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(""),
	}

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		return err
	}

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return err
	}

	t.service = service
	t.webDriver = &wd
	return nil
}

func (t RomototoDriver) FindHousing() []Housing {
	return []Housing{}
}
