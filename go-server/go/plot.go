package swagger

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func convertPointToPlotterXY(v Point) plotter.XYs {
	pts := make(plotter.XYs, 1)

	pts[0] = plotter.XY{
		X: v[0],
		Y: v[1]}

	return pts
}

func addScatters(plt *plot.Plot, index int, name string, xyers plotter.XYs, cent bool) error {
	s, err := plotter.NewScatter(xyers)
	if err != nil {
		return err
	}
	s.Color = plotutil.Color(index)
	s.Shape = plotutil.Shape(index)
	if cent {
		s.Radius = 3.5
	} else {
		s.Radius = 1
	}
	plt.Add(s)
	plt.Legend.Add(name)
	return nil
}

type xy struct{ X, Y float64 }

func sa() {
	data := make(plotter.XYs, 2)
	data[0].X = 18.803247233993062
	data[0].Y = -5.501522589654536e-06
	data[1].X = 7.041001168303179
	data[1].Y = -5.1005727797447055e-05

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	line, err := plotter.NewLine(data)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(line)
	p.Save(4*vg.Inch, 3*vg.Inch, "test.pdf")
}
