package types

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

const DateLayout = "02/01/2006"
const D14s = 14
const K100s = 100000

type States struct {
	States []State `json:"states"`
}

type State struct {
	Rank            int     `json:"rank"`
	State           string  `json:"state"`
	Code            string  `json:"code"`
	Pop             int     `json:"pop"`
	Growth          string  `json:"growth"`
	Pop2018         int64   `json:"pop2018"`
	Pop2010         int64   `json:"pop2010"`
	GrowthSince2010 float64 `json:"growthSince2010"`
	Percent         string  `json:"percent"`
	Density         string  `json:"density"`
}

func (s *States) LoadStates() *States {
	states, err := ioutil.ReadFile("./data/us-states.json")

	if err != nil {
		log.Printf("Unable to open events file %v", err)
	}

	//	fmt.Println(string(states))
	err = yaml.Unmarshal(states, s)

	if err != nil {
		log.Fatalf("Can't parse file %v", err)
	}

	return s
}
