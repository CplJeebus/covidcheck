package records

import (
	"fmt"
	"gocheck/types"
	"strings"
)

func GetDeaths(number int, countries []string, theRecords types.Ecdcdata) {
	var j int

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < number {
					if theRecords.Records[i].GeoID == strings.ToUpper(countries[p]) {
						var deaths = theRecords.Records[i].Deaths

						fmt.Printf("%d\t%s\t%s\n", deaths, theRecords.Records[i].GeoID, theRecords.Records[i].DateRep)
						j++
					}
				}
			}
		}
	}
}
