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

const longUsage = `Usage of ./check-ecdc:
	-c	get total number of new cases
	-d	get total number of new deaths
	-f	get updated file file the ECDC and US CDC
	-n	number of records to return (default 5)
	-dm get number of new deaths per million per day
	-cm get number of new cases per million per day
	-o  Output format (currently):
		raw - Prints a list of the results to the stdout
		csv - Prints csv formatted results to stdout
		plot - Creates a graph "points.png" in the current directory
		default - Creates a graph

	default get cumulative number of new cases per 100K of population for the last 14days.
	A list of country codes must be supplied e.g IE DE ...`

func main() {
	data.Checkfiles()

	number := flag.Int("n", 5, "number of records to return")
	refresh := flag.Bool("f", false, "get updated data file from the ECDC and US-CDC")
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
		fmt.Printf("%s", longUsage)
	}

	if *refresh {
		data.GetData()
	}

	var theRecords types.CovidData

	fbytes, e := ioutil.ReadFile("./data/today.json")
	if e != nil {
		fmt.Printf("%s", e)
	}

	e = json.Unmarshal(fbytes, &theRecords)
	if e != nil {
		fmt.Printf("%s", e)
	}

	var ResultSet []types.CasesRecord
	var title string
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
		title = "New Cases cumulative 14 day per 100K"
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
