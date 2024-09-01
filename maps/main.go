package main

import "fmt"

func main() {


	names:= map[string]int{
		"bayram":1,
		"Telli":2,
		"Özlem":3,
	}


	fmt.Println(names["bayram"])



	// var names map[string]int

	// names = make(map[string]int, 0)

	// names["bayram"] = 1
	// names["emir"] = 2
	// names["özlem"] = 3

	// fmt.Println(names["bayram"])
	// fmt.Println(names["eren"])

}