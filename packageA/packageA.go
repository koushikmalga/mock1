package packageA

import (
	"mock/packageB"
)

type backendA struct {
	b packageB.BackendB
}

func NewA(b packageB.BackendB) backendA {
	return backendA{b: b}
}

func (a backendA) sum(ele1 int, ele2 int) int {
	result := a.b.Sum1(ele1, ele2) * 2
	return result
}
