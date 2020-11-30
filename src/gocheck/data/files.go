package data

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"gocheck/types"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Checkfiles() {
	// Change this so that it only checking for a single file
	// This will be combined CDC and ECDC file
	f, err := os.Stat("./data/today-ecdc.json")

	if os.IsNotExist(err) {
		GetData()
	}

	created := f.ModTime()
	if time.Since(created) > 8*time.Hour {
		fmt.Println("Stale file")
		GetData()
	}
}

func GetData() {
	getdataCdc()
	getdataEcdc()
}

// All this really does is get the raw file from the ECDC
// Will change this to return an array.
func getdataEcdc() {
	dataURL := "https://opendata.ecdc.europa.eu/covid19/casedistribution/json/"
	resp, err := http.Get(dataURL)

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Getting latest file - ECDC")

	defer resp.Body.Close()

	out, err := os.Create("./data/today-ecdc.json")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func getdataCdc() {
	dataURL := "https://data.cdc.gov/api/views/9mfq-cb36/rows.csv?accessType=DOWNLOAD"
	resp, err := http.Get(dataURL)

	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Println("Getting latest file - CDC")

	defer resp.Body.Close()

	out, err := os.Create("./data/cdc-raw.csv")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	createBaseUSjsonFile()
}

func createBaseUSjsonFile() {
	f, err := os.Open("./data/cdc-raw.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	covidRecords := make(types.CovidRecords, 0)

	var covidRecord types.CovidRecord
	var cd types.CovidData

	r := csv.NewReader(bufio.NewReader(f))
	r.Read()
	records, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, rec := range records {
		covidRecord.DateRep = dumbUSdates(rec[0])
		covidRecord.Cases, _ = strconv.Atoi(rec[5])
		covidRecord.Deaths, _ = strconv.Atoi(rec[10])
		covidRecord.GeoID = "US-" + rec[1]
		covidRecord.PopData2019 = getStatePopulation(rec[1])
		covidRecords = append(covidRecords, covidRecord)
	}
	sort.Sort(types.CovidRecords(covidRecords))
	covidRecords.Set14day100k()

	cd.CovidRecords = covidRecords
	file, err := json.Marshal(cd)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./data/today-us.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
	// At this stage we have two files in similar format
}

func dumbUSdates(d string) string {
	s := strings.Split(d, "/")

	return s[1] + "/" + s[0] + "/" + s[2]
}

func getStatePopulation(c string) int {
	var usStates types.States

	usStates.LoadStates()

	for _, state := range usStates.States {
		if state.Code == c {
			return state.Pop
		}
	}

	return 1
}
