package output

import (
	"fmt"
	"gocheck/types"
	"strings"
)

func PrintCases(resultset []types.CasesRecord) {
	for i := range resultset {
		fmt.Printf("%v\t%s\t%s\n", resultset[i].Cases, resultset[i].GeoID, resultset[i].DateRep)
	}
}

func PrintCasesTabs(resultset []types.CasesRecord, countries []string) {
	m := make(map[string][]string)
	rsc := ResultsetByCountry(resultset, countries)

	for i := range rsc {
		for j := range rsc[i] {
			m[rsc[i][j].DateRep] = append(m[rsc[i][j].DateRep], rsc[i][j].Cases)
		}
	}

	fmt.Printf("%v \n", m)
}

func ResultsetByCountry(resultset []types.CasesRecord, countries []string) [][]types.CasesRecord {
	var rs []types.CasesRecord

	var rsc [][]types.CasesRecord

	for i := range countries {
		for j := range resultset {
			if resultset[j].GeoID == strings.ToUpper(countries[i]) {
				rs = append(rs, resultset[j])
			}
		}

		rsc = append(rsc, rs)
		rs = nil
	}
	fmt.Printf("%v \n", rsc)
	return rsc
}
