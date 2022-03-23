package main

import (
    "fmt"

    "github.com/arbaregni/resistor-counts/rationals"
)

// Generate prints out the n-resistor constructable numbers.
func Generate(n int) {

    // we expect something like n^3 rationals to be generated
    est := n*n*n

    // Map each rational to its resistor count
    count := make(map[rationals.Rational]int, est)

    visit := func(layer []rationals.Rational, c int, r rationals.Rational) []rationals.Rational {
        if _, ok := count[r]; ok {
            // already seen, bail out early
            return layer
        }
        count[r] = c
        return append(layer, r)
    }

    // For easy iteration, keep a slice of the rationals we have seen before
    seen := make([]rationals.Rational, 0, est)

//    visit(seen, 0, rationals.MakeRational(0,1))
    seen = visit(seen, 1, rationals.MakeRational(1,1))


    fmt.Println(seen)
    fmt.Println(count)

    // The current level we are on
    c := 2
    for c <= n {

        layer := make([]rationals.Rational, 0, c*c)

        for i, r := range seen {
            for j, s := range seen {
                if j < i {
                    continue
                }

                layer = visit(layer, c, r.Add(s))
                layer = visit(layer, c, r.Harmonic(s))
            }
        }

        fmt.Printf("Layer %v (%v elements)", c, len(layer))

        seen = append(seen, layer...)
        c += 1

        fmt.Scanln()
    }

}







