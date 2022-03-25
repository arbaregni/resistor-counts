package main

import (
	"fmt"
	"log"

	"github.com/arbaregni/resistor-counts/rationals"
)

const noisy = false

// Op enumerates the operations we can use in making RC-numbers.
type Op int

const (
	None Op = iota
	Series
	Parallel
)

// Action represents how we construct new RC-numbers.
type Action struct {
	op         Op
	arg1, arg2 rationals.Rational
}

// DP stores the solutions to the dynamic program.
type DP struct {
	// Map each rational to its resistor count
	count map[rationals.Rational]int
	// Map each rational to how we got here
	backtrack map[rationals.Rational]Action
	// Each layer[c] holds the set of c-resistor constructable numbers
	layers [][]rationals.Rational
}

// NewDP allocates a dynamic program.
// The parameter n is a hint on how much memory to allocate.
func NewDP(n int) *DP {
	// we expect something like n^3 rationals to be generated <--- this can be improved
	est := n * n * n

	dp := &DP{}
	dp.count = make(map[rationals.Rational]int, est)
	dp.backtrack = make(map[rationals.Rational]Action, est)

    cp := n+1
    if 2 > cp {
        cp = 2
    }
	dp.layers = make([][]rationals.Rational, 2, cp)

	return dp
}

// Call this function whenever we see a (potentially) new rational:
// c is the current resistor count (i.e. which layer we are on)
// r is the rational we have constructed
func (dp *DP) visit(c int, r rationals.Rational, a Action) {
	if _, ok := dp.count[r]; ok {
		// already seen, bail out early
		return
	}
	dp.count[r] = c
	dp.backtrack[r] = a
	dp.layers[c] = append(dp.layers[c], r)
}

// Generate returns a slice of layers, where layers[c] contains the set of c-resistor constructable numbers for c=0,1...n
func (dp *DP) Generate(n int) [][]rationals.Rational {
	// check if we have already computed the solution
	if n < len(dp.layers) {
		return dp.layers[:n+1]
	}

	// OK, so now we have to generate more solutions

	// We start with a 1 ohm resistor
	dp.visit(1, rationals.One(), Action{})

	// loop over each layer, until we have enough layers
	for len(dp.layers) <= n {
		c := len(dp.layers)
		{
			l := make([]rationals.Rational, 0, c*c) // estimate c^2 new numbers
			dp.layers = append(dp.layers, l)
		}

		// loop over all the ways to add up to c using positive integers
		foundOne := false
		for i := 1; i <= c/2; i++ {

			j := c - i

			for _, r := range dp.layers[i] {
				for _, s := range dp.layers[j] {

					p := r.Add(s)
					q := r.Harmonic(s)
					dp.visit(c, p, Action{Series, r, s})
					dp.visit(c, q, Action{Parallel, r, s})

					if noisy {
						if !foundOne && (p.Equals(rationals.One()) || q.Equals(rationals.One())) {
							foundOne = true
						}
					}

				}
			}

		}

		if noisy {
			if foundOne {
				fmt.Printf("Found a 1 on layer %v!\n", c)
			}
		}

	}

	return dp.layers
}

// N returns the resistor count of the deepest expanded layer.
func (dp *DP) N() int {
    return len(dp.layers) - 1
}

// Derive returns the formula we used to construct `r` 
func (dp *DP) Derive(r rationals.Rational) string {
    isNeg := false
    switch {
    case r.Equals(rationals.Zero()):
        return "0"
    case r.Equals(rationals.Inf()):
        return "inf"
    case r.IsNeg():
        isNeg = true
        r = rationals.MakeRational(-r.N(), r.D())
    }
    s, ok := dp.deriveHelper(r, None)
    for !ok {
        // try the next layer
        dp.Generate(dp.N() + 1)
        s, ok = dp.deriveHelper(r, None)
    }
    if isNeg {
        s = fmt.Sprintf("-%v", s)
    } else {
        // snip off the leading and trailing ( )
        s = s[1:len(s)-1]
    }
    return s
}
func (dp *DP) deriveHelper(r rationals.Rational, parentOp Op) (string, bool) {
	a, ok := dp.backtrack[r]
	if !ok {
		return "", false
	}

	// Found a terminal value
	if a.op == None {
		if r.D() == 1 {
			return fmt.Sprintf("%v", r.N()), true
		}
		return r.String(), true
	}

	left, ok := dp.deriveHelper(a.arg1, a.op)
	if !ok {
		log.Println("Missing backtrack info for ", a)
		return "", false
	}

	right, ok := dp.deriveHelper(a.arg2, a.op)
	if !ok {
		log.Println("Missing backtrack info for ", a)
		return "", false
	}

	var s string
	switch a.op {
	case Series:
		s = fmt.Sprintf("%v + %v", left, right)
	case Parallel:
		s = fmt.Sprintf("%v || %v", left, right)
	default:
		s = fmt.Sprintf("%v ? %v", left, right)
	}

	if a.op != parentOp {
		s = fmt.Sprintf("(%v)", s)
	}

	return s, true
}
