package records

import (
	"fmt"
	"gocheck/types"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func Get14dayaverage(number int, countries []string, theRecords types.Ecdcdata) {
	var j int

	Rs := make([]types.C14D100K, 0)

	var r types.C14D100K

	if len(countries) != 0 {
		for p := range countries {
			j = 0
			for i := range theRecords.Records {
				if j < number {
					if theRecords.Records[i].GeoID == strings.ToUpper(countries[p]) {
						cases, e := strconv.ParseFloat(theRecords.Records[i].C14D100K, 32)
						if e != nil {
							fmt.Printf("%s", e)
						}

						fmt.Printf("%.2f\t%s\t%s\n", cases, theRecords.Records[i].GeoID, theRecords.Records[i].DateRep)
						r.Cases = theRecords.Records[i].C14D100K
						r.GeoID = theRecords.Records[i].GeoID
						r.DateRep = theRecords.Records[i].DateRep
						Rs = append(Rs, r)
						j++
					}
				}
			}
		}
	}

	CreatePlot(Rs)
}

func CreatePlot(r []types.C14D100K) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p, "first", CreatePoints())
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func CreatePoints(r []types.C14D100K) plotter.XY {
	pts := make(plotter.XY, len(r))
	for i := range pts {
		c, err := strconv.ParseFloat(r[i].Cases, 64)
		pts[i].X = r[i].DateRep
		pts[i].Y = c
	}
	return pts
}
