package main

import (
	"fmt"
	"math"
	"strconv"
)

// --------------------------------------------------------------
type Ticket struct {
	n int
	l int
}

func (t Ticket) Length() int { return t.l }
func (t Ticket) Value() int  { return t.n }

func assert(cond bool) {
	if !cond {
		panic("ups")
	}
}

func (t Ticket) Split(i int) (Ticket, Ticket) {
	assert(0 < i && i < t.Length())
	z := int(math.Pow10(i))
	return Ticket{t.Value() / z, t.Length() - i}, Ticket{t.Value() % z, i}
}

// --------------------------------------------------------------
type Node string

type BinaryOperation interface {
	Apply(a, b int) (int, bool)

	ReverseA(c, b int) (int, bool)
	ReverseB(c, a int) (int, bool)

	Node(a, b Node) Node
}

type UnaryOperation interface {
	Apply(a int) (int, bool)

	Reverse(c int) (int, bool)

	Node(a Node) Node
}

type Solver struct {
	cache     map[Ticket]map[int]Node
	binaryOps []BinaryOperation
	unaryOps  []UnaryOperation
}

func NewSolver() *Solver {
	c := make(map[Ticket]map[int]Node)
	return &Solver{c, nil, nil}
}

func (s *Solver) WithBinaryOps(ops ...BinaryOperation) {
	s.binaryOps = ops
}

func (s *Solver) WithUnaryOps(ops ...UnaryOperation) {
	s.unaryOps = ops
}

func (s *Solver) solutions(t Ticket) map[int]Node {
	if r, ok := s.cache[t]; ok {
		return r
	}
	r := make(map[int]Node)
	// self
	r[t.Value()] = Node(strconv.Itoa(t.Value()))

	// binary ops
	if t.Length() > 1 {
		for i := 1; i < t.Length(); i++ {
			prefix, suffix := t.Split(i)
			for pv, ps := range s.solutions(prefix) {
				for sv, ss := range s.solutions(suffix) {
					for _, operation := range s.binaryOps {
						if rv, ok := operation.Apply(pv, sv); ok {
							if _, ok := r[rv]; !ok { // add if not exists
								r[rv] = operation.Node(ps, ss)
							}
						}
					}
				}
			}
		}
	}

	// unary ops (apply to all results while generates new values)
	for {
		nv := make(map[int]Node)
		for v, n := range r {
			for _, operation := range s.unaryOps {
				if rv, ok := operation.Apply(v); ok {
					if _, ok := r[rv]; !ok { // add if not exists
						nv[rv] = operation.Node(n)
					}
				}
			}
		}
		if len(nv) == 0 {
			break
		}
		for v, n := range nv {
			r[v] = n
		}
	}
	s.cache[t] = r
	return r
}

func (s *Solver) Solve(t Ticket, target int) (node Node, success bool) {
	if t.Length() <= 4 {
		n, ok := s.solutions(t)[target]
		return n, ok
	}

	if r, ok := s.solveBinary(t, target); ok {
		return r, true
	}
	for _, operation := range s.unaryOps {
		if v, ok := operation.Reverse(target); ok {
			if nn, ok := s.solveBinary(t, v); ok {
				return operation.Node(nn), true
			}
		}
	}
	return "", false
}

func (s *Solver) solveBinary(t Ticket, target int) (node Node, success bool) {
	for i := 2; i < t.Length(); i++ {
		if n, ok := s.solve(t, target, i); ok {
			return n, true
		}
	}
	if n, ok := s.solve(t, target, 1); ok {
		return n, true
	}
	return "", false
}

func (s *Solver) solve(t Ticket, target, split int) (node Node, success bool) {
	prefix, suffix := t.Split(split)
	if prefix.Length() >= suffix.Length() {
		for v, n := range s.solutions(suffix) {
			for _, operation := range s.binaryOps {
				if nv, ok := operation.ReverseA(target, v); ok {
					if nn, ok := s.Solve(prefix, nv); ok {
						return operation.Node(nn, n), true
					}
				}
			}
		}
	} else {
		for v, n := range s.solutions(prefix) {
			for _, operation := range s.binaryOps {
				if nv, ok := operation.ReverseB(target, v); ok {
					if nn, ok := s.Solve(suffix, nv); ok {
						return operation.Node(n, nn), true
					}
				}
			}
		}
	}
	return "", false
}

func main() {
	s := NewSolver()
	s.WithBinaryOps(plus, minus, times, div, power)
	s.WithUnaryOps(fact, unaryMinus, sqrt)
	if true {
		solved := 0
		for i := 0; i < 1000000; i++ {
			if i%100000 == 0 {
				println(i / 100000)
			}
			r, ok := s.Solve(Ticket{i, 6}, 100)
			if ok {
				solved++
				// println(r)
				// println(i, r)
				_ = r
			} else {
				fmt.Printf("%06d\n", i)
			}
		}
		println("Total solved:", solved)
	}
}
