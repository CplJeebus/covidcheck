package output

import (
	"gocheck/types"
	. "gocheck/types"
	"math"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreatePlot(r []types.CasesRecord, countries []string, title string, plotEvents bool) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	xticks := plot.TimeTicks{Format: "2006-01-02"}

	p.Title.Text = title
	p.Legend.Top = true
	p.Legend.Left = true
	p.X.Label.Text = "Date"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Value"

	lines := make([]interface{}, 0)
	for c := range countries {
		lines = append(lines, countries[c])
		lines = append(lines, CreatePoints(r, strings.ToUpper(countries[c])))
	}

	if !plotEvents {
		var e Events

		e.LoadEvents()

		for i := range e.Event {
			if isValidEvent(e.Event[i], r) {
				lines = append(lines, e.Event[i].Name+" "+e.Event[i].GeoID)
				lines = append(lines, EventPoints(e.Event[i].Date, GetMaxPoint(r)))
			}
		}
	}
	plotutil.AddLinePoints(p, lines...)

	if err := p.Save(24*vg.Inch, 12*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func isValidEvent(e Event, r []types.CasesRecord) bool {
	layout := "02/01/2006"

	d, _ := time.Parse(layout, e.Date)

	for i := range r {
		t, _ := time.Parse(layout, r[i].DateRep)
		if d.After(t) && (r[i].GeoID == strings.ToUpper(e.GeoID)) {
			return true
		}
	}

	return false
}

func GetMaxPoint(r []types.CasesRecord) float64 {
	var p float64
	p = 0

	for i := range r {
		q, err := strconv.ParseFloat(r[i].Cases, 64)

		if err != nil {
			panic(err)
		}

		if q > p {
			p = q
		}
	}

	return p
}

func EventPoints(d string, mx float64) plotter.XYs {
	pts := make([]plotter.XY, 0)

	var pt plotter.XY

	layout := "02/01/2006"
	t, _ := time.Parse(layout, d)
	pt.X = float64(t.Unix())
	pt.Y = 0
	pts = append(pts, pt)
	pt.Y = mx
	pts = append(pts, pt)

	return pts
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

			if math.IsNaN(c) {
				c = 0
			}

			pt.X = float64(t.Unix())
			pt.Y = c
			pts = append(pts, pt)
		}
	}

	return pts
}
