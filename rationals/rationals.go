package rationals

import "fmt"
import "math"

// Rational represents a rational number n/d, or the point at infinity.
type Rational struct {
    // Invariants:
    //  gcd(n,d) is 1
    //  d > 0
    n, d int
}

// MakeRational constructs the rational representation of n / d.
func MakeRational(n, d int) Rational {
    if d == 0 {
        return Rational{1, 0}
    }
    if d < 0 {
        n = -n
    }
    // gcd computes the greatest common divisor of two integers.
    gcd := func(a, b int) int {
        if b == 0 || a == 0 {
            return 1
        }
        // Euclid's algorithm
        for b != 0 {
            a, b = b, a%b
        }
        return a
    }
    k := gcd(n, d)
    return Rational{n / k, d / k}
}

// Add computes the sum r + s of two rationals (inf plus anything is inf).
func (r Rational) Add(s Rational) Rational {
    // calculate the common denominator
    d := r.d * s.d
    // Normalize the fractions
    n1 := r.n * s.d
    n2 := s.n * r.d
    return MakeRational(n1+n2, d)
}

// Mul computes the product r * s of two rationals (inf times anything is inf)
func (r Rational) Mul(s Rational) Rational {
    return MakeRational(r.n*s.n, r.d*s.d)
}

// Reciprocal computes the reciprocal 1/r of a rational number (1/0 = inf and 1/inf = 0).
func (r Rational) Reciprocal() Rational {
    return MakeRational(r.d, r.n)
}

// Harmonic computes the harmonic sum r || s of two rationals (inf || anything is that input).
func (r Rational) Harmonic(s Rational) Rational {
    rinv, sinv := r.Reciprocal(), s.Reciprocal()
    inv := rinv.Add(sinv)
    return inv.Reciprocal()
}

// Equals compares two rationals.
func (r Rational) Equals(s Rational) bool {
    return r.n * s.d == s.n * r.d
}

// AsFloat divides out the numerator and denominator,
// returning NaN, false if the denominator was zero.
func (r Rational) AsFloat() (float64, bool) {
    if r.d == 0 {
        return math.NaN(), false
    }
    return float64(r.n) / float64(r.d), true
}

// String returns the formatted fraction.
func (r Rational) String() string {
    if r.d == 0 {
        return "inf"
    }
    return fmt.Sprintf("%v//%v", r.n, r.d)
}

// IsNeg tests if the rational is negative.
func (r Rational) IsNeg() bool { return r.n < 0 }

// Numerator returns the numerator n of the rational n/d, in its most reduced form.
func (r Rational) Numerator() int { return r.n }

// Denominator returns the denominator d of the rational n/d, in its most reduced form.
func (r Rational) Denominator() int { return r.d }

// Inf returns the rational representing the point at infinity.
func Inf() Rational { return Rational{1, 0} }

// Zero returns the rational representing 0.
func Zero() Rational { return Rational{0,1} }

// One returns the rational representing 1.
func One() Rational { return Rational{1,1} }
