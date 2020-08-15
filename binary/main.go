package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
)

type ecdcdata struct {
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
		C14D100K                string `json:"c14d100k"`
	} `json:"records"`
}

func main() {
	country := flag.String("c", "", "Dumb give me data for a country")
	number := flag.Int("n", 1, "number of records")

	flag.Parse()
	var theRecords ecdcdata
	fbytes, e := ioutil.ReadFile("../today-fixed.json")

	fmt.Println("Thing " + strconv.Itoa(*number))

	if e != nil {
		fmt.Printf("%s", e)
	}

	e = json.Unmarshal(fbytes, &theRecords)
	if e != nil {
		fmt.Printf("%s", e)
	}
	var j int

	if len(*country) != 0 {
		for i := range theRecords.Records {
			if j < *number {
				if theRecords.Records[i].GeoID == *country {
					fmt.Println(theRecords.Records[i].C14D100K + " " + theRecords.Records[i].GeoID + " " + theRecords.Records[i].DateRep)
					j += 1
				}
			}
		}
	}
}
