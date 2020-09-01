package records

import (

	"gocheck/types"
	"strconv"
	"strings"
)

func GetRecords(number int, countries []string, theRecords types.Ecdcdata, stat string) []types.CasesRecord{
	var j int
	var data string
	ResultSet := make([]types.CasesRecord, 0)

	var result types.CasesRecord

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < number {
					if theRecords.Records[i].GeoID == strings.ToUpper(countries[p]) {
						switch {
						case stat == "deaths":
							data = strconv.Itoa(theRecords.Records[i].Deaths)
						case stat == "cases":
							data = strconv.Itoa(theRecords.Records[i].Cases)
						default:
							data = theRecords.Records[i].C14D100K
						}
						result.Cases = data
						result.GeoID = theRecords.Records[i].GeoID
						result.DateRep = theRecords.Records[i].DateRep
						ResultSet = append(ResultSet, result)
						j++
					}
				}
			}
		}
	}

	return ResultSet
}