package main

import "fmt"

func main() {

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	defer fmt.Println("Bue-bue")

	xs := []int{98, 93, 77, 82, 83}

	avg := average(xs)
	fmt.Println(avg)
	_, dup := duplicate(int(avg))
	fmt.Println(dup)

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println(factorx(5))

	panic("PANIC")
}

func average(xs []int) float64 {
	return float64(add(xs...) / len(xs))
}

func duplicate(x int) (int, int) {
	return x, 2 * x
}

func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorx(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorx(x-1)
}
