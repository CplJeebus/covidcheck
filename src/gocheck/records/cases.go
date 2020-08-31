package records

import (
	"fmt"
	"gocheck/types"
	"strconv"
)

func GetCases(number int, countries []string, theRecords types.Ecdcdata) {
	var j int

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < number {
					if theRecords.Records[i].GeoID == countries[p] {
						cases, e := strconv.ParseFloat(theRecords.Records[i].C14D100K, 32)
						if e != nil {
							fmt.Printf("%s", e)
						}

						fmt.Printf("%.2f\t%s\t%s\n", cases, theRecords.Records[i].GeoID, theRecords.Records[i].DateRep)
						j++
					}
				}
			}
		}
	}
}
