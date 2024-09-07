package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type Flower struct {
	desi int
}

func (book *Book) ShippingCost() int {
	return 5 + book.desi*2
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + electronic.desi*3
}

func (flower *Flower) ShippingCost() int {
	return 12 + flower.desi*3
}

func main() {

	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 30},
		&Flower{desi: 15},
	}

	fmt.Println(products[3].ShippingCost())

	// var product IShippable

	// fmt.Printf("Book shipping cost: %d\n",product.ShippingCost())

	// var electronic IShippable
	// electronic= &Electronic{desi : 10}

	// fmt.Printf("Electronic shipping cost: %d\n",electronic.ShippingCost())

	// book1 := &Book{desi: 10}
	// book2 := &Book{desi: 20}
	// fmt.Println(book1.ShippingCost(),book2.ShippingCost())

	// electronic1 := &Electronic{desi:20}
	// fmt.Println("Electronic shipping cost: ",electronic1.ShippingCost())

	// books := []Book{Book{desi:20},Book{desi:30}}
	// fmt.Println("Books array first item shipping cost: ",books[0].ShippingCost())
	// fmt.Println(calcualteTotalShippingCostOfBooks(books))

}

func calculateTotalShippingCost(products []IShippable) int {
	total := 0

	for _, product := range products {
		total += product.ShippingCost()
	}

	return total
}

// func calcualteTotalShippingCostOfBooks(books []Book) int {

// 	total := 0

// 	for _, book := range books {
// 		total += book.ShippingCost()
// 	}

// 	return total
// }

// func calcualteTotalShippingCostOfElectronics(electronics []Electronic) int {

// 	total := 0

// 	for _, electronic := range electronics {
// 		total += electronic.ShippingCost()
// 	}

// 	return total
// }