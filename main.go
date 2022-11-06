package main

import "fmt"

// product struct represents product inventory
type product struct {
	Model    string
	Quantity int
	Price    float64
}

// car struct represents the car inventory
type car struct {
	product
}

// represents product inventory list
type Store struct {
	product []product
}

func (p product) ShowStock() {
	fmt.Printf("Product: %s", p.Model)

}

//Add creates new store product and appends it to the Store

//Sold method marks a product as sold

func main() {
	//Create a store

	//Create a product

	//Add the product to the store

	//Sell the product

	//List all products in the store

	//List all sold products

}
