package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gocheck/output"
	"gocheck/records"
	"gocheck/types"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	checkfile()

	number := flag.Int("n", 5, "number of records to return")
	refresh := flag.Bool("f", false, "get updated data file from the ECDC")
	deaths := flag.Bool("d", false, "get total number of new deaths")
	cases := flag.Bool("c", false, "get total number of new cases")
	deathspermillion := flag.Bool("dm", false, "get total number of new deaths")
	casespermillion := flag.Bool("cm", false, "get total number of new cases")
	out := flag.String("o", "print", "Output format (currently) plot or screen, default is screen")
	flag.Parse()

	var countries = flag.Args()
	if len(countries) == 0 && !*refresh {
		fmt.Println(`Usage of ./check-ecdc:
	-c	get total number of new cases
	-d	get total number of new deaths
	-f	get updated file file the ECDC
	-n	number of records to return (default 5)

	default get average number of new cases per 100K of population for the last 14days.
	A list of country codes must be supplied e.g IE DE ...`)
	}

	if *refresh {
		getdata()
	}

	var theRecords types.Ecdcdata

	var ResultSet []types.CasesRecord

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
		ResultSet = records.GetRecords(*number, countries, theRecords, "deaths")
	case *cases:
		ResultSet = records.GetRecords(*number, countries, theRecords, "cases")
	case *casespermillion:
		ResultSet = records.GetRecords(*number, countries, theRecords, "casespermillion")
	case *deathspermillion:
		ResultSet = records.GetRecords(*number, countries, theRecords, "deathspermillion")
	default:
		ResultSet = records.GetRecords(*number, countries, theRecords, "c14d100k")
	}

	switch {
	case *out == "plot":
		output.CreatePlot(ResultSet, countries)
	default:
		output.PrintCases(ResultSet)
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
	if err != nil {
		fmt.Printf("%s", err)
	}
}
