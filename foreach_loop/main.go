package main

import "fmt"

func main() {

	var names = map[string]int{
		"bayram": 10,
		"telli":  20,
	}

	for key, value := range names {
		fmt.Printf("key: %s , value: %d \n",key,value)
	}

	// var numbers = []int{1, 2, 3, 4, 5}

	// for _,value :=range numbers{
	// 	fmt.Println(value)
	// }

	// for index := 0; index < len(numbers); index++ {
	// 	fmt.Println(numbers[index])
	// }

	// var language = "go langauge"

	// for _,value := range language{
	// 	fmt.Printf("%c ",value)
	// }

}