package romototo

type HousingStreamer struct {
	providers []HousingProvider
	notifiers []Notifier
}

func (t *HousingStreamer) AddNotifier(notifier Notifier) {
	t.notifiers = append(t.notifiers, notifier)
}

func (t *HousingStreamer) AddProvider(housingProvider HousingProvider) {
	t.providers = append(t.providers, housingProvider)
}

func (t *HousingStreamer) Run() {
	housingChannel := make(chan HousingResult)

	for _, provider := range t.providers {
		currentProvider := provider
		if err := currentProvider.Refresh(); err != nil {
			panic(err)
		}

		go func() {
			var lastHousingId uint32
			for {
				if err := currentProvider.Refresh(); err != nil {
					continue
				}

				housings, err := currentProvider.Query()

				if err == nil && lastHousingId != housings.Id {
					housingChannel <- housings
					lastHousingId = housings.Id
				}
			}
		}()
	}

	for {
		housing := <-housingChannel

		for _, notifier := range t.notifiers {
			notifier.Send(housing)
		}
	}
}
