package cmd

import (
	"github.com/thomasgassmann/romototo/pkg/romototo"
	"github.com/thomasgassmann/romototo/pkg/romototo/config"
	"github.com/thomasgassmann/romototo/pkg/romototo/livingscience"
	"github.com/thomasgassmann/romototo/pkg/romototo/web"
)

func Execute(config *config.Config) {
	driver := web.BrowserDriver{}
	notifier := romototo.MailNotifier{}
	if err := driver.Init(); err != nil {
		panic(err)
	}

	if err := notifier.Init(config.Mail); err != nil {
		panic(err)
	}

	streamer := romototo.HousingStreamer{}
	streamer.Init(driver)
	streamer.AddNotifier(notifier)

	streamer.AddProvider(new(livingscience.LivingScienceHousingProvider))

	streamer.Run()
}
