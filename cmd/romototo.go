package cmd

import (
	"github.com/thomasgassmann/robomoto/pkg/romototo"
	"github.com/thomasgassmann/robomoto/pkg/romototo/notifications"
	"github.com/thomasgassmann/robomoto/pkg/romototo/web"
)


func Execute() {
	driver := web.BrowserDriver{}
	notifier := notifications.Notifier{}
	if err := driver.Init(); err != nil {
		panic(err)
	}

	if err := notifier.Init(); err != nil {
		panic(err)
	}

	housingChannel := make(chan romototo.HousingResult)

	// livingScience := new(romototo.LivingScienceHousingProvider)
	// studentVillage := new(romototo.StudentVillageHousingProvider)
	fake := new(romototo.FakeHousingProvider)
	providers := []romototo.HousingProvider{fake}

	for _, provider := range providers {
		currentProvider := provider
		if err := currentProvider.Init(&driver); err != nil {
			panic(err)
		}
		
		
		if err := currentProvider.Refresh(); err != nil {
			panic(err)
		}

		go func() {
			for {
				housings, err := currentProvider.Query()
				if err != nil {
					housingChannel <- housings
				}
			}
		}()
	}

	for {
		housing := <-housingChannel

		for _, item := range housing.Results {
			println(item.RoomNumber)
		}
	}
}
