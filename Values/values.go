package main 

import "fmt"

func main() {
	
	//Strings, which can be added together with +.
	fmt.Println("go" + "lang")

	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("10-4 =", 10-4)
	fmt.Println("2*2 =", 2*2)
	fmt.Println("4.0/2.0 =", 4.0/2.0)

	//Booleans, with boolean operators as you’d expect.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
