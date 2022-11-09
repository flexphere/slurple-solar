package batch

type DateTime struct {
	Year   int `json:"year"`
	Month  int `json:"month"`
	Day    int `json:"day"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

type MsrInfo struct {
	PconMode         int `json:"pconMode"`
	SupVol           int `json:"supVol"`
	SupTemp          int `json:"supTemp"`
	SupVolDisp       int `json:"supVolDisp"`
	SupTempDisp      int `json:"supTempDisp"`
	PcsActiveCount   int `json:"pcsActiveCount"`
	PConCommErr      int `json:"pConCommErr"`
	SysInitState     int `json:"sysInitState"`
	SupOutputControl int `json:"supOutputControl"`
}

type ErrInfo struct {
	ErrorLevel int    `json:"errorLevel"`
	MachineNo  int    `json:"machineNo"`
	CodeStr    string `json:"codeStr"`
	TimeValue  string `json:"timeValue"`
	ErrorValue string `json:"errorValue"`
}

type Record struct {
	Generation  int `json:"generation"`
	Consumption int `json:"consumption"`
	Selling     int `json:"selling"`
	Buying      int `json:"buying"`
}

type Results struct {
	DateTime                  DateTime `json:"dateTime"`
	MsrInfo                   MsrInfo  `json:"msrInfo"`
	Attainment                int      `json:"attainment"`
	IsStartDateEventDisp      int      `json:"isStartDateEventDisp"`
	EventAnniversary          int      `json:"eventAnniversary"`
	IsEnergyEventDisp         int      `json:"isEnergyEventDisp"`
	EventAttainment           int      `json:"eventAttainment"`
	IsAutoFirmwareUpdatedDisp int      `json:"isAutoFirmwareUpdatedDisp"`
	EventAutoFirmwareUpdated  int      `json:"eventAutoFirmwareUpdated"`
	ErrInfo                   ErrInfo  `json:"errInfo"`
	ReturnCode                int      `json:"returnCode"`
	Year                      int      `json:"year"`
	Month                     int      `json:"month"`
	Day                       int      `json:"day"`
	RecordList                []Record `json:"recordList"`
	TotalGeneration           int      `json:"totalGeneration"`
	TotalConsumption          int      `json:"totalConsumption"`
	TotalSelling              int      `json:"totalSelling"`
	TotalBuying               int      `json:"totalBuying"`
	GraphScale                int      `json:"graphScale"`
}
