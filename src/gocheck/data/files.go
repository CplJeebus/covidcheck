package data

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Checkfiles() {
	f, err := os.Stat("./data/today-cdc.json")

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

	out, err := os.Create("./data/today-cdc.json")
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
}
