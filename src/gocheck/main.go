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
	if *deaths {
		if len(countries) != 0 {
			records.GetDeaths(*number, countries, theRecords)
		}
	} else {
		if len(countries) != 0 {
			records.GetCases(*number, countries, theRecords)

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
