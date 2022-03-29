package main

import "fmt"

func main() {
	// number1 := 17
	// number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string
	if number1, number2 := 17, 24; number1 < number2 {
		resultMessage = "24 is bigger than 17 by"
		difference := number2 - number1
		fmt.Println(resultMessage, difference)
	}
}
