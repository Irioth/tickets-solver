package main

import "math"

var (
	fact       = factOp{}
	unaryMinus = unaryMinusOp{}
	sqrt       = sqrtOp{}

	plus  = plusOp{}
	minus = minusOp{}
	times = timesOp{}
	div   = divOp{}
	power = powerOp{}
)

// ====================================================================
type plusOp struct{}

func (plusOp) Apply(a, b int) (int, bool)    { return a + b, true }
func (plusOp) ReverseA(c, b int) (int, bool) { return c - b, true }
func (plusOp) ReverseB(c, a int) (int, bool) { return c - a, true }
func (plusOp) Node(a, b Node) Node           { return Node("(" + a + ")+(" + b + ")") }

// ====================================================================
type minusOp struct{}

func (minusOp) Apply(a, b int) (int, bool)    { return a - b, true }
func (minusOp) ReverseA(c, b int) (int, bool) { return c + b, true }
func (minusOp) ReverseB(c, a int) (int, bool) { return a - c, true }
func (minusOp) Node(a, b Node) Node           { return Node("(" + a + ")-(" + b + ")") }

// ====================================================================
type timesOp struct{}

func (timesOp) Apply(a, b int) (int, bool) { return a * b, true }
func (timesOp) ReverseA(c, b int) (int, bool) {
	if b == 0 || c%b != 0 {
		return 0, false
	}
	return c / b, true
}
func (timesOp) ReverseB(c, a int) (int, bool) { return timesOp{}.ReverseA(c, a) }
func (timesOp) Node(a, b Node) Node           { return Node("(" + a + ")*(" + b + ")") }

// ====================================================================
type divOp struct{}

func (divOp) Apply(a, b int) (int, bool) {
	if b == 0 || a%b != 0 {
		return 0, false
	}
	return a / b, true
}
func (divOp) ReverseA(c, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return c * b, true
}
func (divOp) ReverseB(c, a int) (int, bool) {
	if a == 0 {
		return 0, false
	}
	return divOp{}.Apply(a, c)
}
func (divOp) Node(a, b Node) Node { return Node("(" + a + ")/(" + b + ")") }

// ====================================================================
type powerOp struct{}

func (powerOp) Apply(a, b int) (int, bool) {
	if b == 0 {
		return 1, true
	}
	if b < 0 || a < 0 {
		return 0, false
	}
	if a == 1 {
		return 1, true
	}
	if a == 0 {
		return 0, true
	}
	p := 1
	for i := 0; i < b; i++ {
		p *= a
		if p > 10000 {
			return 0, false
		}
	}
	return p, true
}

func (o powerOp) ReverseA(c, b int) (int, bool) {
	if c <= 0 || b <= 0 {
		return 0, false
	}
	q := math.Pow(2, math.Log2(float64(c))/float64(b))
	a := int(q)
	if r, ok := o.Apply(a, b); ok && r == c {
		return a, true
	}
	if r, ok := o.Apply(a+1, b); ok && r == c {
		return a + 1, true
	}
	return 0, false
}
func (powerOp) ReverseB(c, a int) (int, bool) {
	p := 1
	for i := 0; i < 20; i++ {
		if p == c {
			return i, true
		}
		p *= a
		if p > c {
			return 0, false
		}
	}
	return 0, false
}
func (powerOp) Node(a, b Node) Node { return Node("(" + a + ")^(" + b + ")") }

// ====================================================================
type factOp struct{}

var factVal = map[int]int{0: 1, 1: 1, 2: 2, 3: 6, 4: 24, 5: 120, 6: 720, 7: 5040, 8: 40320, 9: 362880, 10: 3628800, 11: 39916800, 12: 479001600, 13: 6227020800, 14: 87178291200, 15: 1307674368000, 16: 20922789888000, 17: 355687428096000, 18: 6402373705728000, 19: 121645100408832000, 20: 2432902008176640000}

func (factOp) Apply(a int) (int, bool) {
	r, ok := factVal[a]
	return r, ok
}
func (factOp) Reverse(c int) (int, bool) { return 0, false }
func (factOp) Node(a Node) Node          { return Node("(" + a + ")!") }

// ====================================================================
type unaryMinusOp struct{}

func (unaryMinusOp) Apply(a int) (int, bool)   { return -a, true }
func (unaryMinusOp) Reverse(c int) (int, bool) { return -c, true }
func (unaryMinusOp) Node(a Node) Node          { return Node("-(" + a + ")") }

// ====================================================================
type sqrtOp struct{}

var sqrtVal = map[int]int{4: 2, 9: 3, 16: 4, 25: 5, 36: 6, 49: 7}

func (sqrtOp) Apply(a int) (int, bool) {
	if a <= 0 {
		return 0, false
	}
	r := int(math.Sqrt(float64(a)))
	if r*r != a {
		return 0, false
	}
	return r, true
}
func (sqrtOp) Reverse(c int) (int, bool) {
	if c < 0 || c > 1000000 {
		return 0, false
	}
	return c * c, true
}
func (sqrtOp) Node(a Node) Node { return Node("sqrt(" + a + ")") }
