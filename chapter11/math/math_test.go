package math
import "math"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1,2})
	if v!= 1.5 {
		t.Error("Expexted 1.5, got ", v)
	}
}
