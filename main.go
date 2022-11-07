package main

import "fmt"

// product struct represents product inventory
type Product struct {
	Model    string
	Quantity int
	Price    float64
}

// car struct represents the car inventory
type car struct {
	Product
}

func (p Product) ShowStock() {
	fmt.Printf("Product: %s", p.Model)
}

func (p Product) itemStatus() {
	if p.Quantity > 0 {
		fmt.Println(" In Stock")
	} else {
		fmt.Println(" Not In Stock")
	}
}

// represents product inventory list
type Store struct {
	Inventory []Product
	inventoryList int
	sold int
	gainTotal float64

}

//add method adds a product to the store
func (s *Store) add(p Product) {
	s.Inventory = append(s.Inventory, p)
	s.inventoryList++
}

// InventoryList list all products
func (s *Store) InventoryList() {
	for _, p := range s.Inventory {
		p.ShowStock()
	}
}

//Sell method marks a product as sold
func (s *Store) sell(Model string, quantity int) {
	for i, p := range s.Inventory {
		if p.Model == Model {
			if p.Quantity >= quantity {
				s.Inventory[i].Quantity -= quantity
				s.sold += quantity
				s.gainTotal += p.Price * float64(quantity)
			} else {
				fmt.Println("Not enough quantity")
			}
		}
	}
}
func main() {

	auto1 := car{"Honda", "50", 50000}
	auto2 := car{"Honda", "50", 50000}

	//Create a store
	Store := Store{}

	//Create new product
	pro1 := Product{[]car{auto1}, "Model" 50, 50000}

	//Add the product to the store
	Store.add(auto1)
	//Sell the product
	Store.Sell()

	//List all products in the store
	Store.inventoryList
	//List all sold products
	Store.sold

}
