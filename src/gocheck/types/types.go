package types

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type CovidData struct {
	CovidRecords []CovidRecord `json:"records"`
}

type CovidRecord struct {
	DateRep                 string `json:"dateRep"`
	Day                     string `json:"day"`
	Month                   string `json:"month"`
	Year                    string `json:"year"`
	Cases                   int    `json:"cases"`
	Deaths                  int    `json:"deaths"`
	CountriesAndTerritories string `json:"countriesAndTerritories"`
	GeoID                   string `json:"geoId"`
	CountryterritoryCode    string `json:"countryterritoryCode"`
	PopData2019             int    `json:"popData2019"`
	ContinentExp            string `json:"continentExp"`
	C14D100K                string `json:"Cumulative_number_for_14_days_of_COVID-19_cases_per_100000"`
}

type CasesRecord struct {
	Cases   string
	GeoID   string
	DateRep string
}

type Events struct {
	Event []Event `yaml:"events"`
}

type Event struct {
	Date  string `yaml:"date"`
	Name  string `yaml:"event"`
	GeoID string `yaml:"geoid"`
}

func (e *Events) LoadEvents() *Events {
	eventsFile, err := ioutil.ReadFile("events.yaml")

	if err != nil {
		log.Printf("Unable to open events file %v", err)
	}

	err = yaml.Unmarshal(eventsFile, e)

	if err != nil {
		log.Fatalf("Can't parse file %v", err)
	}

	return e
}
