package main

// Utility functions to generate fake data
import (
	"math/rand"
	"time"

	"syreclabs.com/go/faker"
)

// Generate a fake Customer
func genFakeCustomer() Customer {
	rand.Seed(time.Now().UnixNano())
	fName := faker.Name()
	return Customer{
		Name:   fName.FirstName() + " " + fName.LastName(),
		Number: rand.Intn(60),
	}
}

// Generate a fake Order
func genFakeCoffee() OrderCoffee {
	fName := faker.Name()
	return OrderCoffee{
		Beverage: fName.FirstName(),
		Quantity: rand.Intn(10),
	}
}
