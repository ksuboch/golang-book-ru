package main

import (
	"fmt"
)

func main() {
	var a [5]float64
	a[0] = 97
	a[1] = 98
	a[2] = 77
	a[3] = 82
	a[4] = 65

	b := [5]float64{
		97,
		98,
		77,
		82,
		65,
	}

	avg := 0.0
	for i := 0; i < len(a); i++ {
		avg += a[i] / float64(len(a))
	}
	fmt.Println(avg)

	total := 0.0
	for _, val := range b {
		total += val
	}
	fmt.Println(total / float64(len(a)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	slice1[0] = 3
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
	slice3[0] = 3
	fmt.Println(slice3, slice4)

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
