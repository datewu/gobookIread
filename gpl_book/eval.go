package main

import (
	"fmt"
	"io"
	"math"
	"net/http"

	"./eval"
)

func main() {
	http.HandleFunc("/plot", plot)
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := eval.ParseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expression:"+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	fn := func(m, n float64) float64 {
		distance := math.Hypot(m, n) // distance from (0,0)
		return expr.Eval(eval.Env{"x": m, "y": n, "r": distance})
	}

	surface(w, fn)
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Hello World! <hr> %s", form)
}

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0) // sin(30°), cos(30°)
const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	form          = `
<form action="/plot" method="post">
    <label for="expr">Expression:</label>
    <input type="text" name="expr">
    <input type="submit" value="submit">
</form>
	`
)

func surface(w io.Writer, f func(float64, float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(f func(float64, float64) float64, i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
