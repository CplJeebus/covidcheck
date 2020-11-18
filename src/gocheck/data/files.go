package data

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"gocheck/types"
	"io"
	"net/http"
	"os"
	"strconv"
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
		covidRecord.DateRep = rec[0]
		covidRecord.Cases, _ = strconv.Atoi(rec[5])
		covidRecord.Deaths, _ = strconv.Atoi(rec[10])
		covidRecord.GeoID = "us-" + rec[1]
		covidRecord.PopData2019 = getStatePopulation(rec[1])
		covidRecords = append(covidRecords, covidRecord)
	}
	covidRecord = calculate14DayAvg(covidRecord)
}

func set14day100k (rs []types.CovidRecord(){
	var usStates types.States

	usStates.LoadStates()
	for _, s := range usStates.States {
		for i, r range rs {
			if r.GeoID = "us-"+s.Code {
				if d < 14 {

				}
			}
		}
	}

}

func calculate14Day100k(rs []types.CovidRecords, d int){
		if d < 14 {
			j=0
		} else {
			j=d-14
		}
		for j < d {

		}
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
