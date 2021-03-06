package main

import "fmt"

/*
func main() {
    var x [5]int
    x[4] = 100
    fmt.Println(x)
}

func main() {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83

    var total float64 = 0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    fmt.Println(total / float64(len(x)))
}

func main() {
    var x [5]float64
    x[0] = 98
    x[1] = 93
    x[2] = 77
    x[3] = 82
    x[4] = 83

    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))
}

func main() {
    x := [5]float64{ 98, 93, 77, 82, 83 }
    var total float64 = 0
    for _, value := range x {
        total += value
    }
    fmt.Println(total / float64(len(x)))
}

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2)
}

func main() {
    slice1 := []int{1, 2, 3}
    slice2 := make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)
}

func main() {
    x := make(map[string]int)
    x["key"] = 10
    fmt.Println(x["key"])
}

func main() {
    x := make(map[int]int)
    x[1] = 10
    fmt.Println(x[1])
}

func main() {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"

    if name, ok := elements["Un"]; ok {
        fmt.Println(name, ok)
    }
}

func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}

func main() {
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

	if el, ok := elements["F"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}*/

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	minimal := x[0]
	for _, value := range x {
		if value < minimal {
			minimal = value
		}
	}
	fmt.Println(minimal)
}
