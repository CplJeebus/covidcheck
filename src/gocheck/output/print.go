package output

import (
	"fmt"
	"gocheck/types"
)

func PrintCases(resultset []types.CasesRecord) {
	for i := range resultset {
		fmt.Printf("%v\t%s\t%s\n", resultset[i].Cases, resultset[i].GeoID, resultset[i].DateRep)
	}
}
