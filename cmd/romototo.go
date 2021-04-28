package cmd

import (
	"github.com/thomasgassmann/robomoto/pkg/romototo"
	"github.com/thomasgassmann/robomoto/pkg/romototo/web"
)


func Execute() {
	driver := web.BrowserDriver{}
	err := driver.Init()
	if err != nil {
		panic(err)
	}

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
			for _, housing := range housings.Results {
				println(housing.RoomNumber)
			}
		}
	}
}
