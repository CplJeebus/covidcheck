package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		C14D100K                string `json:"Cumulative_number_for_14_days_of_COVID-19_cases_per_100000"`
	} `json:"records"`
}

func main() {
	checkfile()

	number := flag.Int("n", 1, "number of records")
	refresh := flag.Bool("f", false, "get updated file")
	flag.Parse()

	var countries []string
	countries = flag.Args()

	if *refresh == true {
		getdata()
	}

	var theRecords ecdcdata
	fbytes, e := ioutil.ReadFile("./today-go.json")
	if e != nil {
		fmt.Printf("%s", e)
	}

	e = json.Unmarshal(fbytes, &theRecords)
	if e != nil {
		fmt.Printf("%s", e)
	}
	var j int
	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < *number {
					if theRecords.Records[i].GeoID == countries[p] {
						cases, e := strconv.ParseFloat(theRecords.Records[i].C14D100K, 32)
						if e != nil {
							fmt.Printf("%s", e)
						}
						fmt.Printf("%.2f\t%s\t%s\n", cases, theRecords.Records[i].GeoID, theRecords.Records[i].DateRep)
						j += 1
					}
				}
			}
		}
	}
}

func checkfile() {
	_, err := os.Stat("./today-go.json")
	if os.IsNotExist(err) {
		getdata()
	}
}

func getdata() {
	dataUrl := "https://opendata.ecdc.europa.eu/covid19/casedistribution/json/"
	resp, err := http.Get(dataUrl)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Getting latest file")
	defer resp.Body.Close()

	out, err := os.Create("./today-go.json")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

}
