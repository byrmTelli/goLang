package main

import "fmt"

func main() {

	var a int
	a = 10

	var p *int


	p=&a

	fmt.Printf("p nin adresindeki değer %d\na nın adresi %d",*p,a)

}