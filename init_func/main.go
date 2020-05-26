package main

var (
	x int = 5
)

func main() {
	y := func(x int) int {
		z := x &^ 02
		return z
	}

	println(x, y(x))
}

func init() {
	// you can't save function to the variable here
	x &= 6
}
