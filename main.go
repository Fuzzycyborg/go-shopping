package main

import "fmt"

//implement the car class with a Display() method that prints out the car's details

type Car struct {
	make  string
	model string
	year  int
	color string
}

func (c *Car) Display() {
	fmt.Println("Make: ", c.make)
	fmt.Println("Model: ", c.model)
	fmt.Println("Year: ", c.year)
	fmt.Println("Color: ", c.color)
}

//Implement the Product class with a Display() method that prints out the product's details

type Product struct {
	car      *Car
	quantity int
	price    int
}

func (p *Product) Display() {
	p.car.Display()
	fmt.Println("Quantity: ", p.quantity)
	fmt.Println("Price: ", p.price)
}

//IsInStock() method returns a boolean indicating whether the product is in stock or not.

func (p *Product) IsInStock() bool {
	return p.quantity > 0
}

//Implement the Store class with the following methods: AddProduct(), ListProducts(), SellProduct(), ListSoldProducts()

type Store struct {
	products        []*Product
	soldProducts    []*Product 
	totalStockValue int 
	totalSalesValue int 
}

func (s *Store) AddProduct(p *Product) {
	s.products = append(s.products, p) //add product to products
	s.totalStockValue += p.quantity * p.price //increment total stock value
}

func (s *Store) ListProducts() {
	for _, p := range s.products {
		p.Display()
	}
}

func (s *Store) SellProduct(p *Product) {
	//check if product is in stock
	
	if !p.IsInStock() {
		fmt.Println("Sorry, this product is out of stock")
		return
	}

	p.quantity-- //decrement quantity
	s.totalStockValue -= p.price //decrement total stock value
	s.totalSalesValue += p.price //increment total sales value
	s.soldProducts = append(s.soldProducts, p) //add product to sold products
}

func (s *Store) ListSoldProducts() {
	for _, p := range s.soldProducts {
		p.Display()
	}
	fmt.Println("Total Sales: ", s.totalSalesValue) //print total sales value
}
