package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gocheck/data"
	"gocheck/output"
	"gocheck/records"
	"gocheck/types"
	"io/ioutil"
)

func main() {
	data.Checkfiles()

	number := flag.Int("n", 5, "number of records to return")
	refresh := flag.Bool("f", false, "get updated data file from the ECDC")
	deaths := flag.Bool("d", false, "get total number of new deaths per day")
	cases := flag.Bool("c", false, "get total number of new cases per day")
	deathspermillion := flag.Bool("dm", false, "get number of new deaths per million per day")
	casespermillion := flag.Bool("cm", false, "get number of new cases per million per day")
	events := flag.Bool("xe", false, "disable the display of events in the output graphs")
	out := flag.String("o", "", `Output format (currently):
raw - Prints a list of the results to the stdout
csv - Prints csv formatted results to stdout
plot - Creates a graph "points.png" in the current directory
default - Creates a graph"`)

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
		raw - Prints a list of the results to the stdout
		csv - Prints csv formatted results to stdout
		plot - Creates a graph "points.png" in the current directory
		default - Creates a graph

	default get average number of new cases per 100K of population for the last 14days.
	A list of country codes must be supplied e.g IE DE ...`)
	}

	if *refresh {
		data.GetData()
	}

	var Records types.CovidData
	var theRecords types.CovidData
	CovidRS := make([]types.CovidRecord, 0)
	var ResultSet []types.CasesRecord
	var title string

	// I hate this two file solution!!
	// Once I get the MVP need a big refactor.
	fbytesA, e := ioutil.ReadFile("./data/today-ecdc.json")
	if e != nil {
		fmt.Printf("%s", e)
	}

	e = json.Unmarshal(fbytesA, &Records)
	if e != nil {
		fmt.Printf("%s", e)
	}

	for _, c := range Records.CovidRecords {
		CovidRS = append(CovidRS, c)
	}

	fbytesB, e := ioutil.ReadFile("./data/today-us.json")
	if e != nil {
		fmt.Printf("%s", e)
	}
	e = json.Unmarshal(fbytesB, &Records)
	if e != nil {
		fmt.Printf("%s", e)
	}

	for _, c := range Records.CovidRecords {
		CovidRS = append(CovidRS, c)
	}

	theRecords.CovidRecords = CovidRS

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
		output.CreatePlot(ResultSet, countries, title, *events)
	case *out == "raw":
		output.PrintCases(ResultSet)
	case *out == "csv":
		output.PrintCasesTabs(ResultSet, countries)
	default:
		output.CreatePlot(ResultSet, countries, title, *events)
	}
}
