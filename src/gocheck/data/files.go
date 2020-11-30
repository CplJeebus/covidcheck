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
	"strconv"
	"strings"
	"time"
)

func Checkfiles() {
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

	covidRecords := make([]types.CovidRecord, 0)

	var covidRecord types.CovidRecord

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

	var cd types.CovidData
	cd.CovidRecords = set14day100k(covidRecords)
	file, err := json.Marshal(cd)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./data/today-us.json", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func set14day100k(rs []types.CovidRecord) []types.CovidRecord {
	var usStates types.States

	d := int(0)
	// We are going to make some assumtions here.
	// Mainly that all of the records are in date order.
	// And that all states are contiguous
	usStates.LoadStates()

	for _, s := range usStates.States {
		d = 0

		for i, r := range rs {
			if r.GeoID == "US-"+s.Code {
				rs[i].C14D100K = strconv.FormatFloat(calculateRange(rs, i, d), 'f', 6, 64)
				d++
			}
		}
	}

	return rs
}

func dumbUSdates(d string) string {
	s := strings.Split(d, "/")
	return s[1] + "/" + s[0] + "/" + s[2]
}

func calculateRange(rs []types.CovidRecord, index, d int) float64 {
	var s float64
	var si int

	if d < 14 {
		si = (index - d)

		for _, p := range rs[si:index] {
			s = (s + float64(p.Cases))
		}
		s = s / float64(d)
	} else {
		si = (index - 14)
		for _, p := range rs[si:index] {
			s = (s + float64(p.Cases))
		}
		s = s / 14
	}

	return s
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
