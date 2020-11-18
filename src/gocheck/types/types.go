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

type States struct {
	States []State `json:"states"`
}

type State struct {
	Rank            int     `json:"rank"`
	State           string  `json:"state"`
	Code            string  `json:"code"`
	Pop             int     `json:"pop"`
	Growth          string  `json:"growth"`
	Pop2018         int64   `json:"pop2018"`
	Pop2010         int64   `json:"pop2010"`
	GrowthSince2010 float64 `json:"growthSince2010"`
	Percent         string  `json:"percent"`
	Density         string  `json:"density"`
}

func (s *States) LoadStates() *States {
	states, err := ioutil.ReadFile("./data/us-states.json")

	if err != nil {
		log.Printf("Unable to open events file %v", err)
	}

	//	fmt.Println(string(states))
	err = yaml.Unmarshal(states, s)

	if err != nil {
		log.Fatalf("Can't parse file %v", err)
	}

	return s
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
