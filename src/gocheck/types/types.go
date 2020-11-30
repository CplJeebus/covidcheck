package types

import (
	"io/ioutil"
	"log"
	"strconv"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type CovidData struct {
	CovidRecords CovidRecords `json:"records"`
}

type CovidRecords []CovidRecord
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

func (rs CovidRecords) Len() int      { return len(rs) }
func (rs CovidRecords) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs CovidRecords) Less(i, j int) bool {
	layout := "02/01/2006"
	a, _ := time.Parse(layout, rs[i].DateRep)
	b, _ := time.Parse(layout, rs[j].DateRep)

	return b.Before(a)
}

func (rs CovidRecords) Set14day100k() CovidRecords {
	var usStates States

	usStates.LoadStates()

	d := int(0)

	for _, s := range usStates.States {
		d = 0

		for i, r := range rs {
			if r.GeoID == "US-"+s.Code {
				rs[i].C14D100K = strconv.FormatFloat(calculateRange(rs, i, d), 'f', 6, 64)
				d++
			}
		}
	}

	return rs
}

func calculateRange(rs []CovidRecord, index, d int) float64 {
	var s float64

	var si int

	if d < 14 {
		si = (index - d)

		for _, p := range rs[si:index] {
			s = (s + float64(p.Cases))
		}

		s = (s / float64(rs[index].PopData2019)) * 100000
	} else {
		si = (index - 14)
		for _, p := range rs[si:index] {
			s = (s + float64(p.Cases))
		}
		s = (s / float64(rs[index].PopData2019)) * 100000
	}

	return s
}
