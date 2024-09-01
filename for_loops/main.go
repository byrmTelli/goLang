package main

import "fmt"

func main() {
	// index := 1

	// for index <= 10 {
	// 	fmt.Println(index)
	// 	index ++
	// }

	var names = [3]string{"bayram","emir","özlem"}

	for index:= 0 ; index <= len(names);index ++ {
		fmt.Println(names[index])
		}
}