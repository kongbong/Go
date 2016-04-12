// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
    "fmt"
    "math"
)

const (
    width, height   = 600, 320              // canvas size in pixels
    cells           = 100                   // number of grid cells
    xyrange         = 30.0                  // axis ranges (-xyrange..+xyrange)
    xyscale         = width / 2 / xyrange   // pixels per x or y unit
    zscale          = height * 0.4          // pixels per z unit
    angle           = math.Pi / 6           // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)     // sin(30), cos(30)

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' " +
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            ax, ay, az, err := corner(i+1, j)
            if err { continue }
            bx, by, bz, err := corner(i, j)
            if err { continue }
            cx, cy, cz, err := corner(i, j+1)
            if err { continue }
            dx, dy, dz, err := corner(i+1, j+1)
            if err { continue }
            
            z := (az + bz + cz + dz) / 4
            c := uint32(0x0000ff + (0xff0000 - 0x0000ff) / 2 * (z + 1))
            r := uint8(c & 0xff)
            g := uint8((c >> 8) & 0xff)
            b := uint8((c >> 16) & 0xff)             
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:rgb(%d, %d, %d)'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
        }
    }
    fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
    // Find point (x,y) at corner of cell (i, j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)
    
    // Compute surface height z.
    z := f(x,y)
    if math.IsNaN(z) {
        return 0, 0, 0, true
    }
    
    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z, false
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) // distance from (0,0)
    return math.Sin(r)/r
}