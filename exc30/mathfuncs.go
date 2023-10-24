package main

func main() {

}

func doMath(x int, y int, mathFunc func(int, int) int) int {
	return mathFunc(x, y)
}

func add(x int, y int) int {
	return x + y
}
func subtract(x int, y int) int {
	return x - y
}
