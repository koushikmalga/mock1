package packageB

import (
	"mock/packageC"
	"mock/packageD"
)

type BackendB interface {
	Sum1(int, int) int
}

type backendBmockB interface {
	sum2(int, int) int
}

type mockC struct{}

func (m mockC) sum2(ele1 int, ele2 int) int {
	return 5
}

type mockD struct{}

func (m mockD) prod(ele1 int, ele2 int) int {
	return 6
}

type mockB struct {
	c mockC
	d mockD
}
type B struct {
	c packageC.BackendC
	d packageD.BackendD
}

func NewB(c packageC.BackendC, d packageD.BackendD) BackendB {
	return &B{c: c, d: d}
}

func mockNewB(c mockC, d mockD) backendBmockB {
	return &mockB{c: c, d: d}
}

func (b *mockB) sum2(ele1 int, ele2 int) int {
	result := b.c.sum2(ele1, ele2) + b.d.prod(ele1, ele2)
	return result
}
func (b *B) Sum1(ele1 int, ele2 int) int {

	result := b.d.Prod(ele1, ele2) + b.c.Sum2(ele1, ele2)
	return result
}
