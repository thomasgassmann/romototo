package web

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type BrowserDriver struct {
	service *selenium.Service
	webDriver *selenium.WebDriver
}

const (
	port = 8080
	seleniumPath = "/usr/share/selenium-server/selenium-server-standalone.jar"
	chromeDriver = "/usr/bin/chromedriver"
)

func (t *BrowserDriver) Driver() selenium.WebDriver {
	return *t.webDriver
}

func (t *BrowserDriver) Quit() error {
	if err := (*t.webDriver).Quit(); err != nil {
		return err
	}

	if err := (*t.service).Stop(); err != nil {
		return err
	}

	return nil
}

func (t *BrowserDriver) Init() error {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriver),
	}

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		return err
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return err
	}

	t.service = service
	t.webDriver = &wd

	return nil
}
