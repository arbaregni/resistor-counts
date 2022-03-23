package main

import (
    "fmt"

    "github.com/arbaregni/resistor-counts/rationals"
)

const noisy = false

// Generate returns a slice of layers, where layers[c] contains the set of c-resistor constructable numbers for c=0,1...n
func Generate(n int) [][]rationals.Rational {

    // we expect something like n^3 rationals to be generated <--- this can be improved
    est := n*n*n

    // Map each rational to its resistor count
    count := make(map[rationals.Rational]int, est)

    layers := make([][]rationals.Rational, n + 1)

    // Call this function whenever we see a new rational:
    // c is the current resistor count (i.e. which layer we are on)
    // r is the rational we have constructed
    visit := func(c int, r rationals.Rational) {
        if _, ok := count[r]; ok {
            // already seen, bail out early
            return
        }
        count[r] = c
        layers[c] = append(layers[c], r)
    }

    // We start with a 1 ohm resistor
    visit(1, rationals.One())

    // loop over each layer
    for c := 2; c <= n; c++ {

        // loop over all the ways to add up to c using positive integers
        foundOne := false
        for i := 1; i <= c/2; i++ {

            j := c - i

            for _, r := range layers[i] {
                for _, s := range layers[j] {

                    p := r.Add(s)
                    q := r.Harmonic(s)
                    visit(c, p)
                    visit(c, q)

                    if !foundOne && (p.Equals(rationals.One()) || q.Equals(rationals.One())) {
                        foundOne = true
                    }

                }
            }

        }


        if noisy {
            if foundOne {
                fmt.Printf("Found a 1 on layer %v!\n", c)
            }
            fmt.Printf("Layer %v (%v elements)\n", c, len(layers[c]))
            fmt.Printf("[press enter to show]", c+1)
            fmt.Scanln()
            for i, r := range layers[c] {
                if i != 0 && i % 10 == 0 {
                    fmt.Println()
                }
                fmt.Printf("    %v", r)
            }
            fmt.Println()
        }

    }

    return layers
}







