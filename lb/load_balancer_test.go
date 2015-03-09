package lb

import (
	"testing"
)

func TestCalcWeightPercentages(t *testing.T) {
	weights := []int{1, 2, 3}
	expected := []int{16, 33, 50}
	actual := calcWeightPercentages(weights)

	for i,v := range(actual) {
		if v != expected[i] {
			t.Error(v, "not equal", expected[i])
		}
	}
}
