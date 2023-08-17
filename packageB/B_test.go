package packageB

import (
	"testing"
)

/*type mockC struct{}

func (m mockC) sum2(ele1 int, ele2 int) int {
	return ele1 + ele2
}

type mockD struct{}

func (m mockD) prod(ele1 int, ele2 int) int {
	return ele1 * ele2
}*/

func TestB(t *testing.T) {
	b := mockNewB(mockC{}, mockD{})
	result := b.sum2(2, 3)

	exp_result := 11
	if result != exp_result {
		t.Errorf("expected result = %d , but result = %d", exp_result, result)
	}
}
