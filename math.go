package main

func getBasicArithmetic(a int, b int) (int, int, int, float64) {
	add := a + b
	subtract := a - b
	multiply := a * b
	divide := float64(a) / float64(b)
	return add, subtract, multiply, divide
}
