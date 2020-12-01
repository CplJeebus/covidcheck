package types

import (
	"strconv"
	"time"
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

func (rs CovidRecords) Len() int      { return len(rs) }
func (rs CovidRecords) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }
func (rs CovidRecords) Less(i, j int) bool {
	a, _ := time.Parse(DateLayout, rs[i].DateRep)
	b, _ := time.Parse(DateLayout, rs[j].DateRep)

	return b.Before(a)
}

func (rs CovidRecords) Set14day100k() CovidRecords {
	var usStates States

	usStates.LoadStates()

	for _, s := range usStates.States {
		for i, r := range rs {
			if r.GeoID == "US-"+s.Code {
				rs[i].C14D100K = strconv.FormatFloat(calculateRange(rs, i), 'f', 6, 64)
			}
		}
	}

	return rs
}

func calculateRange(rs []CovidRecord, index int) float64 {
	s := float64(0)

	se := (index + D14s)
	for _, p := range rs[index:se] {
		s = (s + float64(p.Cases))
	}

	s = (s / float64(rs[index].PopData2019)) * K100s

	return s
}
