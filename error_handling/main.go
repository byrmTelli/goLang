package main

import (
	"errors"
	"fmt"
)

func divide(x int, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("divider value may not be 0")
	}

	return x / y, nil
}

type Product struct {
	id int
	name string
	price int
}

type ProductService struct {

}
func (productService *ProductService) Add(product Product) error {

	if len(product.name) == 0 {
		return ValidationError{text:"Product name may not be empty.",code:"400"}
	}

	if product.price < 0 {
		return ValidationError{text:"Product price value may not be lover than 0",code:"400"}
	}

	fmt.Println("Product added successfully.")

	return nil
}


// Creating Custom Error

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata: %s,\nKod: %s",validationError.text,validationError.code)
}

func main() {


	productService := ProductService{}
	
	err:= productService.Add(Product{id:1,name:"asd",price:-1})

	if err != nil {
		fmt.Println(err)
	}

	//Function Error Handling
	// result,err := divide(15,3)
	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(result)
	// }



	//File Read
	// fileData,err := os.ReadFile("../note.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(fileData)
	// }

	// //for instance you create a variable that type is intger but did not assign any value in it
	// var x int
	// fmt.Println("integer: ",x)

	// //or float
	// var y float32
	// fmt.Println("float: ",y)

	// // or pointer
	// var pointer1 *int
	// fmt.Println("pointer: ",pointer1)

	// if pointer1 == nil {
	// 	// fmt.Sprintf kullanarak C#'taki interpolasyon gibi yazdırma
	// 	fmt.Println(fmt.Sprintf("pointer1 (%T) herhangi bir yeri göstermiyor!", pointer1))
	// }
}