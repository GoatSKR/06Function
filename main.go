package main

import "fmt"

// Function that adds two integers
// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(3, 4)
// 	fmt.Println("Sum:", result) // Output: Sum: 7
// }

// Function that divides two integers and returns the result and an error
// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	// result, err := divide(10, 2)
// 	result, err := divide(10, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Result:", result) // Output: Result: 5
// 	}
// }

// // Function with named return values
// func swap(a int, b int) (x int, y int) {
// 	x = b
// 	y = a
// 	return // Named return values are returned automatically
// }

// func main() {
// 	a, b := 1, 2
// 	x, y := swap(a, b)
// 	fmt.Println("Swapped:", x, y) // Output: Swapped: 2 1
// }

// // Function that takes a variable number of integers and returns their sum
// func sum(numbers ...int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// func main() {
// 	result := sum(1, 2, 3, 4, 5)
// 	fmt.Println("Sum:", result) // Output: Sum: 15
// }

// func main() {
// 	// Anonymous function assigned to a variable
// 	greet := func(name string) {
// 		fmt.Println("Hello,", name)
// 	}
// 	greet("Alice") // Output: Hello, Alice

// 	// Anonymous function used directly
// 	result := func(a, b int) int {
// 		return a * b
// 	}(3, 4) // Immediate invocation
// 	fmt.Println("Product:", result) // Output: Product: 12
// }

// Function that takes another function as an argument
// func operate(a int, b int, op func(int, int) int) int {
// 	return op(a, b)
// }

// func add(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	result := operate(5, 3, add)
// 	fmt.Println("Result of addition:", result) // Output: Result of addition: 8
// }

// // Function that returns a closure
// func makeMultiplier(factor int) func(int) int {
// 	return func(x int) int {
// 		return x * factor
// 	}
// }

// func main() {
// 	multiplier := makeMultiplier(3)
// 	fmt.Println("Multiplier result:", multiplier(5)) // Output: Multiplier result: 15
// }

// Define a struct
type Rectangle struct {
	Width  int
	Height int
}

// Method on Rectangle to calculate area
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of Rectangle:", rect.Area()) // Output: Area of Rectangle: 50
}
