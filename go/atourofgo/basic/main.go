package main

import "fmt"

func main() {
	stringBased()
	iotaBased()
	fmt.Println(split(17))
}

// named return values
func split(sum int) (z int) {
	a := 0
	z = sum/2 + a
	return z
}

// enum
func stringBased() {
	type Color string
	const (
		Red    Color = "Red"
		Yellow Color = "Yellow"
		Green  Color = "Green"
	)

	colors := []Color{Red, Yellow, Green}

	for _, c := range colors {
		if c == Red {
			fmt.Println(c)
		} else {
			fmt.Println("NO")
		}
	}
}

func iotaBased() {
	type Color int
	const (
		Red Color = iota
		Yellow
		Green
	)

	colors := []Color{Red, Yellow, Green}

	for _, c := range colors {
		if c == 0 {
			fmt.Println(c)
		} else {
			fmt.Println("NO")
		}
	}
}
