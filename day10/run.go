package day10

import "fmt"

func Run() {
	fmt.Println("Running Day 10 solution...")
	// Call your solution here

	callFunc(getTime, 1000)
	fmt.Println("Solution completed.")
	callFuncAsync(getTime, 1000)
	fmt.Println("Asynchronous solution completed.")

}

func getTime() {
	fmt.Println("Time taken to run the solution:")
}
