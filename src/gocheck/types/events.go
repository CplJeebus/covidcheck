package types

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Events struct {
	Event []Event `yaml:"events"`
}

type Event struct {
	Date  string `yaml:"date"`
	Name  string `yaml:"event"`
	GeoID string `yaml:"geoid"`
}

func (e *Events) LoadEvents() *Events {
	eventsFile, err := ioutil.ReadFile("events.yaml")

	if err != nil {
		log.Printf("Unable to open events file %v", err)
	}

	err = yaml.Unmarshal(eventsFile, e)

	if err != nil {
		log.Fatalf("Can't parse file %v", err)
	}

	return e
}
