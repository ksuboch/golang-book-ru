package math

//Average Returns the average value from the array
func Average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}

// Max Returns the maximum value from the array
func Max(xs []float64) float64 {
	max := xs[0]
	for _, val := range xs[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

// Min Returns the minimum value from the array
func Min(xs []float64) float64 {
	min := xs[0]
	for _, val := range xs[1:] {
		if val < min {
			min = val
		}
	}
	return min
}
