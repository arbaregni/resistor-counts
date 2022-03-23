package main

import (
    "github.com/arbaregni/resistor-counts/rationals"

    "image"
    "image/color"
)

// images that we can change
type changeable interface {
    Set(x,y int, col color.Color)
}

func drawTick(img changeable, col color.Color, x0, y0, width, height int) {
    for y := y0; y < y0 + height; y++ {
        for x := x0; x < x0 + width; x++ {
            img.Set(x,y, col)
        }
    }
}

// Visualize constructs the diagram depicting the construction layers (ignoring width and height for now).
func Visualize(layers [][]rationals.Rational, width, height int) image.Image {
    n := len(layers)

    width = 1024
    layerSize := 16
    tickSize := 1
    height = n * layerSize

    img := image.NewRGBA(image.Rect(0, 0, width, height))

    tickCol := color.Black

    for c := range layers {
        for _, r := range layers[c] {
            // r is in the range [0 n]
            f, ok := r.AsFloat()
            if !ok {
                continue
            }
            pct := f / float64(n)  // pct is in the range [0 1]
            frac := pct * float64(width) // frac is in the range [0 width]
            x := int(frac)

            y := layerSize * (c - 1)

            rem := height - y
            drawTick(img, tickCol, x, y, tickSize, rem)
        }
    }

    return img
}
