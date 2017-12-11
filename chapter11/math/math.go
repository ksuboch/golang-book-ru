package math

//Average Returns the average value from the array
func Average(xs []float64) float64 {
	total := 0.0
	for _, val := range xs {
		total += val
	}
	return total / float64(len(xs))
}
