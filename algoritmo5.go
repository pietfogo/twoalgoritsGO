	package main 

	import "fmt"

	func main()  {
		numbers := [] int {12, 7, 233, 99, 4, 44, 67, 302}
		indexNumbers := len(numbers)
		for i := 0; i < indexNumbers-1; i++ {
			for c := 0; c < indexNumbers-i-1; c++ {	
			if numbers[c] > numbers[c+1]{
				numbers[c], numbers[c+1] = numbers[c+1], numbers[c]
			}
		}
	}
		fmt.Printf("Array: %v", numbers)
	}

