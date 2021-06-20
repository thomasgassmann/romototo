package cmd

import (
	"github.com/thomasgassmann/robomoto/pkg/romototo"
	"github.com/thomasgassmann/robomoto/pkg/romototo/web"
)


func Execute() {
	driver := web.BrowserDriver{}
	notifier := romototo.Notifier{}
	if err := driver.Init(); err != nil {
		panic(err)
	}

	if err := notifier.Init(); err != nil {
		panic(err)
	}

	streamer := romototo.HousingStreamer{}
	streamer.Init(driver)
	streamer.AddNotifier(notifier)

	streamer.AddProvider(new(romototo.LivingScienceHousingProvider))

	streamer.Run()
}
