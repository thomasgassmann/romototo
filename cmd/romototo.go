package cmd

import (
	"github.com/thomasgassmann/romototo/pkg/romototo"
	"github.com/thomasgassmann/romototo/pkg/romototo/config"
	"github.com/thomasgassmann/romototo/pkg/romototo/livingscience"
)

func Execute(config *config.Config) {
	notifier := romototo.MailNotifier{}

	if err := notifier.Init(config.Mail); err != nil {
		panic(err)
	}

	streamer := romototo.HousingStreamer{}

	streamer.AddNotifier(notifier)

	if livingScience, err := livingscience.InitSeleniumLivingScience(); err == nil {
		streamer.AddProvider(livingScience)
	} else {
		panic(err)
	}

	streamer.Run()
}
