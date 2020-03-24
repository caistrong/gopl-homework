package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange .....+xyrange)
	angle   = math.Pi / 6 // angle of x,y axes (=30˚)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// go command:
// go run ./src/chapter3/work3_4/work3_4.go
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		width := r.FormValue("width")
		widthInt, err := strconv.ParseInt(width, 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><body>You need a param 'width' and that must be  a number, you can visit like<a href=\"http://localhost:8888/?width=2160&height=1280&color=red\">http://localhost:8888/?width=2160&height=1280&color=red</a></body></html>")
			return
		}
		height := r.FormValue("height")
		heightInt, err := strconv.ParseInt(height, 10, 64)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, "<html><body>You need a param 'height' and that must be  a number, you can visit like<a href=\"http://localhost:8888/?width=2160&height=1280&color=red\">http://localhost:8888/?width=2160&height=1280&color=red</a></body></html>")
			return
		}
		color := r.FormValue("color")
		surface(w, widthInt, heightInt, color)
	})
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func surface(out http.ResponseWriter, width int64, height int64, color string) {
	out.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg'"+
		" style='stroke: %s;fill:white;stroke-width:0.7'"+
		" width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)

			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				// 跳过无效的多边形
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)

		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, width int64, height int64) (float64, float64) {
	// Find Point (x, y) at corner of cell(i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	zscale := float64(height) * 0.4         // pixels per z unit
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit

	// compute surface height z
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := float64(width/2) + (x-y)*cos30*xyscale
	sy := float64(height/2) + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
