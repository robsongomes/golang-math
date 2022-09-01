//math contains the math basic operations
package math

//Mul calculate the sum between a and b
func Add(a, b int) int {
	return a + b
}

//Mul calculate the subtraction between a and b
func Sub(a, b int) int {
	return a - b
}

//Mul calculate the multiplication between a and b
func Mul(a, b int) int {
	return a * b
}

//Div calculate the division between a and b. If b == 0 then the result will be 0
func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
