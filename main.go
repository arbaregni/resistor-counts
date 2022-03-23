package main

import "github.com/arbaregni/resistor-counts/rationals"
import "fmt"

func main() {
    r := rationals.MakeRational(1, 3)
    s := rationals.MakeRational(2, 4)

    fmt.Printf("%v || %v = %v\n", r, s, r.Harmonic(s))

    s = rationals.Inf()
    fmt.Printf("%v || %v = %v\n", r, s, r.Harmonic(s))
}
