package records

import (
	"fmt"
	"gocheck/types"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
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

	CreatePlot(Rs, countries)
}

func CreatePlot(r []types.C14D100K, countries []string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	xticks := plot.TimeTicks{Format: "2006-01-02"}

	p.Title.Text = "News Cases per 100K 14 day average"
	p.X.Label.Text = "Date"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Value"

	for c := range countries {
		line, points, err := plotter.NewLinePoints(CreatePoints(r, strings.ToUpper(countries[c])))
		if err != nil {
			panic(err)
		}

		p.Add(line, points)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func CreatePoints(r []types.C14D100K, s string) plotter.XYs {
	pts := make(plotter.XYs, len(r))
	for i := range pts {
		if r[i].GeoID == s {
			layout := "02/01/2006"
			t, _ := time.Parse(layout, r[i].DateRep)
			c, err := strconv.ParseFloat(r[i].Cases, 64)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s %f \n", t, c)
			pts[i].X = float64(t.Unix())
			pts[i].Y = c
		}
	}

	return pts
}
