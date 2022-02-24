package main

import (
	"log"
	"time"

	"syreclabs.com/go/faker"
	"syreclabs.com/go/faker/locales"
)

var totalCount int
var orderWaitStart float64
var customerWaitStart float64

func assignOrderPicker(order *chan OrderCoffee, custChan *chan Customer, coffeeConcurrency int, customerBufferSize int) {

	var customers chan Customer
	var orders chan OrderCoffee

	customers = make(chan Customer, customerBufferSize)
	orders = make(chan OrderCoffee, coffeeConcurrency)

	go stats(*custChan, *order, customers, orders)

	for true {
		var ord OrderCoffee
		var custom Customer
		var orderStart time.Time
		var customerStart time.Time

		go func() {
			var custom Customer
			custom = <-*custChan
			log.Printf("Customer %s arrived ", custom)
			customerStart = time.Now()
			customers <- custom
		}()

		// Order Served Buffer
		go func() {
			var ord OrderCoffee
			ord = <-*order
			log.Printf("Order %s Serverd", ord)
			orderStart = time.Now()
			orders <- ord
		}()

		ord = <-orders
		custom = <-customers

		orderWaitStart += time.Now().Sub(orderStart).Seconds()
		customerWaitStart += time.Now().Sub(customerStart).Seconds()

		totalCount += 1

		log.Printf("Assigned Customer \"%s\" Order Sereved at \"%s\"", ord, custom)
	}

}

func stats(customers chan Customer, orders chan OrderCoffee, customerBuffer chan Customer, orderBuffer chan OrderCoffee) {
	var template = "--[stats: waiting Customers: %d, waiting orders: %d, Customer wait time (avg): %.2f, Order wait time (avg): %.2f ]--\n"
	for true {
		var waitPassTime float64
		var waitTaxiTime float64
		var waitPassCount int
		var waitTaxiCount int

		waitTaxiCount = len(orders) + len(orderBuffer)
		waitPassCount = len(customers) + len(customerBuffer)

		// Averages
		if totalCount > 0 {
			waitPassTime = orderWaitStart / float64(totalCount)
			waitTaxiTime = customerWaitStart / float64(totalCount)
		}

		log.Printf(template, waitPassCount, waitTaxiCount, waitPassTime, waitTaxiTime)
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func main() {

	var custChan chan Customer
	var order chan OrderCoffee

	faker.Locale = locales.En_IND

	order = make(chan OrderCoffee, 1000)
	custChan = make(chan Customer, 500)

	go genLongQueue(&custChan)
	go generatePickupTraffic(&order)

	assignOrderPicker(&order, &custChan, 3, 10)
}
