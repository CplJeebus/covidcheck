package records

import (
	"fmt"
	"gocheck/types"
	"strings"
)

func GetCases(number int, countries []string, theRecords types.Ecdcdata) {
	var j int

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < number {
					if theRecords.Records[i].GeoID == strings.ToUpper(countries[p]) {
						var cases = theRecords.Records[i].Cases

						fmt.Printf("%d\t%s\t%s\n", cases, theRecords.Records[i].GeoID, theRecords.Records[i].DateRep)
						j++
					}
				}
			}
		}
	}
}
