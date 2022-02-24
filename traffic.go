package main

import (
	"math/rand"
	"time"
)

// Functions to simulate the traffic of the user

// this functions smiulates the higher traffic of customers
func genLongQueue(custChan *chan Customer) {
	for true {
		var person Customer

		person = genFakeCustomer()
		*custChan <- person

		time.Sleep(time.Duration(rand.Intn(10)+5) * time.Second)
	}
}

func genLowerQueue(custChan *chan Customer) {
	for true {
		var person Customer

		person = genFakeCustomer()
		*custChan <- person

		time.Sleep(time.Duration(rand.Intn(30)+30) * time.Second)
	}
}

func genMedQueue(custChan *chan Customer) {
	for true {
		var person Customer

		person = genFakeCustomer()
		*custChan <- person

		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)

	}
}

// This function simulates a rather uniform traffic of
// Customers Taxis arriving to pick up fares
func generatePickupTraffic(order *chan OrderCoffee) {
	for true {
		var reciver OrderCoffee
		reciver = genFakeCoffee()
		// Push to channel
		*order <- reciver
		// Let us assume 1 taxi arrives evey 15 - 30s
		time.Sleep(time.Duration(rand.Intn(15)+15) * time.Second)
	}
}
