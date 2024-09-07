package main

import (
	"fmt"
	helper "my_module/helper"
	helper2 "my_module/helper/helper_2"
	rest2 "my_module/helper/rest"
	rest1 "my_module/rest"
)
func main() {

	fmt.Println("Hello world!")

	helper.Helper1()
	rest1.Rest1()
	helper2.Helper2()
	rest2.Rest1()

	// rest2.rest2() // is not accesible.
}