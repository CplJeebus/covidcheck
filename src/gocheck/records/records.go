package records

import (
	"gocheck/types"
	"strconv"
	"strings"
)

func GetRecords(number int, countries []string, theRecords types.CovidData, stat string) []types.CasesRecord {
	var j int

	var data string

	ResultSet := make([]types.CasesRecord, 0)

	var result types.CasesRecord

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.CovidRecords {
				if j < number {
					if theRecords.CovidRecords[i].GeoID == strings.ToUpper(countries[p]) {
						switch {
						case stat == "deaths":
							data = strconv.Itoa(theRecords.CovidRecords[i].Deaths)
						case stat == "cases":
							data = strconv.Itoa(theRecords.CovidRecords[i].Cases)
						case stat == "casespermillion":
							data = StatsPerMillion(theRecords.CovidRecords[i].Cases, theRecords.CovidRecords[i].PopData2019)
						case stat == "deathspermillion":
							data = StatsPerMillion(theRecords.CovidRecords[i].Deaths, theRecords.CovidRecords[i].PopData2019)
						default:
							data = theRecords.CovidRecords[i].C14D100K
						}

						if s, _ := strconv.ParseFloat(data, 64); s < 0 {
							data = "0"
						}

						result.Cases = data
						result.GeoID = theRecords.CovidRecords[i].GeoID
						result.DateRep = theRecords.CovidRecords[i].DateRep
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
