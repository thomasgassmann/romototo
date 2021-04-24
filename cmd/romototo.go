package cmd

import "github.com/thomasgassmann/robomoto/pkg/romototo"

func Execute() {
	driver := romototo.RomototoDriver{}
	driver.Init()

	housingOffers := driver.FindHousing()
	for _, housingOffer := range housingOffers {
		println(housingOffer.Name)
	}
}
