package batch

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	repo "github.com/flexphere/slurple-solar/repository"

	_ "github.com/mattn/go-sqlite3"
)

const URL_HISTORY_DAILY = "http://%s/asyncquery.cgi?type=Record&timeType=1&year=%d&month=%d&day=%d&timeStamp=%d"

type SolarService struct {
	repository repo.SolarRepository
}

func New(repo repo.SolarRepository) *SolarService {
	return &SolarService{
		repository: repo,
	}
}

func (s *SolarService) Today() {
	t := time.Now()
	result := s.FetchResults(t)
	s.repository.SaveRecords(s.formatRecords(result))
}

func (s *SolarService) Duration(days int) {
	t := time.Now()
	from := t.Add(-(time.Hour * time.Duration(24*days)))
	results := s.FetchAllResults(from, t)
	for _, result := range results {
		s.repository.SaveRecords(s.formatRecords(result))
	}
}

func (s *SolarService) FetchAllResults(from time.Time, to time.Time) []Results {
	results := []Results{}
	for t := from; t.Before(to); t = t.AddDate(0, 0, 1) {
		results = append(results, s.FetchResults(t))
	}
	return results
}

func (s *SolarService) FetchResults(t time.Time) Results {
	url := fmt.Sprintf(URL_HISTORY_DAILY, os.Getenv("DATASOURCE_HOST"), t.Year(), t.Month(), t.Day(), t.UnixNano())
	body, err := Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	results := Results{}
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return results
}

func (s *SolarService) formatRecords(result Results) []repo.SolarRecord {
	records := []repo.SolarRecord{}
	for i, record := range result.RecordList {
		records = append(records, repo.SolarRecord{
			TS:          time.Date(result.Year, time.Month(result.Month), result.Day, i, 0, 0, 0, time.Local).UnixNano(),
			Year:        result.Year,
			Month:       result.Month,
			Day:         result.Day,
			Generation:  record.Generation,
			Consumption: record.Consumption,
			Selling:     record.Selling,
			Buying:      record.Buying,
		})
	}
	return records
}
