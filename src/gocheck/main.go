package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gocheck/records"
	"gocheck/types"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	checkfile()

	number := flag.Int("n", 1, "number of records to return")
	refresh := flag.Bool("f", false, "get updated file")
	deaths := flag.Bool("d", false, "get number of deaths")
	fourteendayavg := flag.Bool("c", false, "get average number of new cases per 100K of population")
	flag.Parse()

	var countries = flag.Args()

	if *refresh {
		getdata()
	}

	var theRecords types.Ecdcdata

	fbytes, e := ioutil.ReadFile("./today-go.json")

	if e != nil {
		fmt.Printf("%s", e)
	}

	e = json.Unmarshal(fbytes, &theRecords)
	if e != nil {
		fmt.Printf("%s", e)
	}

	switch {
	case *deaths:
		records.GetDeaths(*number, countries, theRecords)
	case *fourteendayavg:
		records.Get14dayaverage(*number, countries, theRecords)
	default:
		records.GetCases(*number, countries, theRecords)
	}
}

func checkfile() {
	_, err := os.Stat("./today-go.json")
	if os.IsNotExist(err) {
		getdata()
	}
}

func getdata() {
	dataURL := "https://opendata.ecdc.europa.eu/covid19/casedistribution/json/"
	resp, err := http.Get(dataURL)

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
