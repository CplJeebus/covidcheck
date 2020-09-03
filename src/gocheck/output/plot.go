package output

import (
	"gocheck/types"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreatePlot(r []types.CasesRecord, countries []string, title string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	xticks := plot.TimeTicks{Format: "2006-01-02"}

	p.Title.Text = title
	p.Legend.Top = true
	p.Legend.Left = false
	p.X.Label.Text = "Date"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Value"

	lines := make([]interface{}, 0)
	for c := range countries {
		lines = append(lines, countries[c])
		lines = append(lines, CreatePoints(r, strings.ToUpper(countries[c])))
	}

	_ = plotutil.AddLinePoints(p, lines...)

	if err := p.Save(12*vg.Inch, 12*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func CreatePoints(r []types.CasesRecord, s string) plotter.XYs {
	pts := make([]plotter.XY, 0)

	var pt plotter.XY

	for i := range r {
		if r[i].GeoID == s {
			layout := "02/01/2006"
			t, _ := time.Parse(layout, r[i].DateRep)
			c, err := strconv.ParseFloat(r[i].Cases, 64)

			if err != nil {
				panic(err)
			}

			pt.X = float64(t.Unix())
			pt.Y = c
			pts = append(pts, pt)
		}
	}

	return pts
}
