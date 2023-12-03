package sum

import "testing"

func TestSum(t *testing.T) {
	if sum := Sum(1, 2); sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}
