package repository

type SolarRecord struct {
	TS          int64 `json:"ts"`
	Year        int   `json:"year"`
	Month       int   `json:"month"`
	Day         int   `json:"day"`
	Hour        int   `json:"hour"`
	Generation  int   `json:"generation"`
	Consumption int   `json:"consumption"`
	Selling     int   `json:"selling"`
	Buying      int   `json:"buying"`
}

type SolarRepository interface {
	Connect()
	Disconnect()
	SaveRecords(records []SolarRecord)
}
