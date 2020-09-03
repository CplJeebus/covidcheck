package records

import (
	"gocheck/types"
	"strconv"
	"strings"
)

func GetRecords(number int, countries []string, theRecords types.Ecdcdata, stat string) []types.CasesRecord {
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
						case stat == "casespermillion":
							data = StatsPerMillion(theRecords.Records[i].Cases, theRecords.Records[i].PopData2019)
						case stat == "deathspermillion":
							data = StatsPerMillion(theRecords.Records[i].Deaths, theRecords.Records[i].PopData2019)
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

func StatsPerMillion(stat, pop int) string {
	const million float64 = 1000000

	statf := float64(stat)
	popf := float64(pop)

	return strconv.FormatFloat((statf / (popf / million)), 'f', 4, 64)
}
