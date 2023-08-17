package packageC

type BackendC interface {
	Sum2(int, int) int
}

type C struct{}

func (c C) Sum2(ele1 int, ele2 int) int {
	return ele1 + ele2
}
