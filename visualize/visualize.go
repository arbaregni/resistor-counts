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


