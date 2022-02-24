package main

import (
	"fmt"
)

type Customer struct {
	Name   string `json:name`
	Number int    `json:number`
}

type OrderCoffee struct {
	Beverage string `json:beverage`
	Quantity int    `json:qunatity`
}

func (c Customer) String() string {
	return fmt.Sprint("Name: %s, Number:%d", c.Name, c.Number)
}

func (o OrderCoffee) String() string {
	return fmt.Sprint("Beverage: %s, Quantity: %s ", o.Beverage, o.Quantity)
}
