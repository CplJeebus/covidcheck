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
	deaths := flag.Bool("d", false, "get total number of new deaths per day")
	cases := flag.Bool("c", false, "get total number of new cases per day")
	deathspermillion := flag.Bool("dm", false, "get number of new deaths per million per day")
	casespermillion := flag.Bool("cm", false, "get number of new cases per million per day")
	out := flag.String("o", "", `Output format (currently):
default - Prints a list of the results to the stdout
csv - Prints csv formatted results to stdout
plot - Creates a graph "points.png"`)

	flag.Parse()

	var countries = flag.Args()
	if len(countries) == 0 && !*refresh {
		fmt.Println(`Usage of ./check-ecdc:
	-c	get total number of new cases
	-d	get total number of new deaths
	-f	get updated file file the ECDC
	-n	number of records to return (default 5)

	-dm get number of new deaths per million per day
	-cm get number of new cases per million per day
	-o  Output format (currently):
		default - Prints a list of the results to the stdout
		csv - Prints csv formatted results to stdout
		plot - Creates a graph "points.png" in the current directory

	default get average number of new cases per 100K of population for the last 14days.
	A list of country codes must be supplied e.g IE DE ...`)
	}

	if *refresh {
		getdata()
	}

	var theRecords types.Ecdcdata

	var ResultSet []types.CasesRecord

	var title string

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
		title = "New Deaths per day"
	case *cases:
		ResultSet = records.GetRecords(*number, countries, theRecords, "cases")
		title = "New Cases per day"
	case *casespermillion:
		ResultSet = records.GetRecords(*number, countries, theRecords, "casespermillion")
		title = "New Cases per million of pop per day"
	case *deathspermillion:
		ResultSet = records.GetRecords(*number, countries, theRecords, "deathspermillion")
		title = "New Deaths per million of pop per day"
	default:
		ResultSet = records.GetRecords(*number, countries, theRecords, "c14d100k")
		title = "New Cases per 100K 14 day average"
	}

	switch {
	case *out == "plot":
		output.CreatePlot(ResultSet, countries, title)
	case *out == "csv":
		output.PrintCasesTabs(ResultSet, countries)
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
