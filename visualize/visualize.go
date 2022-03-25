package visualize

import (
	"github.com/arbaregni/resistor-counts/rationals"

	"image"
	"image/color"
	"log"
	"math"
)

// images that we can change
type changeable interface {
	Set(x, y int, col color.Color)
}

func drawTick(img changeable, col color.Color, x0, y0, width, height int) {
	for y := y0; y < y0+height; y++ {
		for x := x0; x < x0+width; x++ {
			img.Set(x, y, col)
		}
	}
}

// LineDiagram constructs the diagram depicting the construction layers
func LineDiagram(layers [][]rationals.Rational) image.Image {
	n := float64(len(layers))

    width := 1024
	layerSize := 16
	tickSize := 1
    height := len(layers) * layerSize

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	tickCol := color.Black

	for c := range layers {
		for _, r := range layers[c] {
			// f is in the range [1/n n]
			f, ok := r.AsFloat()
			if !ok {
				log.Println("ignoring NaN in visualize")
				continue
			}
			f = math.Log(f)             // remap to [-log(n) log(n)]
			f += math.Log(n)            // remap to [0 2log(n)]
			f = f / (2.0 * math.Log(n)) // remap to [0 1]
			f = f * float64(width)      // remap to [0 width]

			x := int(f)
			y := layerSize * (c - 1)
			rem := height - y

			drawTick(img, tickCol, x, y, tickSize, rem)
		}
	}

	return img
}

// HeatDiagram constructs a grid showing us how hard each rational is to construct
func HeatDiagram(layers [][]rationals.Rational) image.Image {
    cellSize := 8

    // find the largest numerator and denominator
    maxP, maxQ := 0, 0
    for c := range layers {
        for _, r := range layers[c] {
            if r.N() > maxP {
                maxP = r.N()
            }
            if r.D() > maxQ {
                maxQ = r.D()
            }
        }
    }

    width, height := maxP * cellSize, maxQ * cellSize

	img := image.NewRGBA(image.Rect(0, 0, width, height))

    startCol := color.RGBA{92, 10, 10, 0xff} // dark red
    endCol   := color.RGBA{100, 167, 222, 0xff} // light blue

    n := float64(len(layers))

    for c := range layers {
        for _, r := range layers[c] {
            x := cellSize * (r.N() - 1)
            y := cellSize * (r.D() - 1)

            col := lerpCol(startCol, endCol, float64(c)/n)

            drawTick(img, col, x, y, cellSize, cellSize)
        }

    }

    return img
}

func lerpCol(col1, col2 color.Color, t float64) color.Color {
    r1, b1, g1, a1 := col1.RGBA()
    r2, b2, g2, a2 := col2.RGBA()
    r := lerp(r1, r2, t)
    b := lerp(b1, b2, t)
    g := lerp(g1, g2, t)
    a := lerp(a1, a2, t)
    return color.RGBA{r,b,g,a}
}
func lerp(a, b uint32, t float64) uint8 {
    return uint8(float64(a) + (float64(b)-float64(a)) * t)
}


