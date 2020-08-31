package types

type Ecdcdata struct {
	Records []struct {
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
	} `json:"records"`
}
