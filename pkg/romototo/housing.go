package romototo

type Housing struct {
	RoomNumber string
}

type HousingResult struct {
	Results []Housing
	Screenshot []byte
	Id uint32
}
