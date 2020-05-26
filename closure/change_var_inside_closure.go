package main

func main() {
	i := 0
	func() {
		i = 17
	}()
	print(i)
}
