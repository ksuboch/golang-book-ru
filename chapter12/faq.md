# Задачи
### Написать хороший набор тестов не всегда легко, но даже сам процесс их написания, зачастую, может выявить много проблем для первой реализации функции. Например, что произойдет с нашей функцией Average, если ей передать пустой список ([]float64{})? Как нужно изменить функцию, чтобы она возвращала 0 в таких случаях?
```go
//Average Returns the average value from the array
func Average(xs []float64) float64 {
	total := 0.0
	if len(xs) == 0 {
		return 0
	}

	for _, val := range xs {
		total += val
	}

	return total / float64(len(xs))
}
```
***
### Напишите серию тестов для функций Min и Max из предыдущей главы.
```go
package math

import "testing"

type testpair struct {
	values []float64
	res    float64
}

var minTests = []testpair{
	{[]float64{1, 2, 3}, 1},
	{[]float64{2, 1, 0, -1, -2}, -2},
	{[]float64{-0.1, 0.1}, -0.1},
}

var maxTests = []testpair{
	{[]float64{1, 2, 3}, 3},
	{[]float64{-2, -1, 0, 1, 2}, 2},
	{[]float64{-0.1, 0.1}, 0.1},
}
func TestMin(t *testing.T) {
	for _, pair := range minTests {
		v := Min(pair.values)
		if v != pair.res {
			t.Error(
				"For", pair.values,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxTests {
		v := Max(pair.values)
		if v != pair.res {
			t.Error(
				"For", pair.values,
				"expected", pair.res,
				"got", v,
			)
		}
	}
}
```