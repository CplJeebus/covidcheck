package output

import (
	"fmt"
	"gocheck/types"
)

func PrintCases(ResultSet []types.CasesRecord){

	for i := range ResultSet {
				fmt.Printf("%v\t%s\t%s\n", ResultSet[i].Cases, ResultSet[i].GeoID, ResultSet[i].DateRep)
			}
}
