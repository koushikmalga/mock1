package packageA

import (
	"fmt"
	"mock/packageB"
	"mock/packageC"
	"mock/packageD"
	"testing"
)

func TestA(t *testing.T) {
	c := packageC.C{}
	d := packageD.D{}
	b := packageB.NewB(&c, &d)
	a := NewA(b)

	var ele1, ele2 int
	ele1 = 2
	ele2 = 3
	result := a.sum(ele1, ele2)

	fmt.Println(result)

	exp_result := 22

	if result != exp_result {
		t.Errorf("Expected result = %d, but got =%d", exp_result, result)
	}

}
