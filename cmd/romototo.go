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

	livingScience := new(romototo.LivingScienceHousingProvider)
	// studentVillage := new(romototo.StudentVillageHousingProvider)
	providers := []romototo.HousingProvider{livingScience}

	for _, provider := range providers {
		provider.Init(&driver)
		if err := provider.Refresh(); err != nil {
			panic(err)
		}

		housings, err := provider.Query()
		if err != nil {
			notifier.Send(housings)
			for _, housing := range housings.Results {
				println(housing.RoomNumber)
			}
		}
	}
}
