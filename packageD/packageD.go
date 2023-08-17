package packageD

type BackendD interface {
	Prod(int, int) int
}

type D struct{}

func (d *D) Prod(ele1 int, ele2 int) int {

	return ele1 * ele2
}
