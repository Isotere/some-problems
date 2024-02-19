package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem

const (
	NumberOfPizzas = 10
)

var (
	pizzasMade   int64
	pizzasFailed int64
	total        int64
)

// makePizza attempts to make a pizza. We generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order: #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza: #%d. It will take %d seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		switch {
		case rnd <= 2:
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d ***", pizzaNumber)
		case rnd <= 4:
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d ***", pizzaNumber)
		default:
			{
				success = true
				msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
			}
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		currentPizza := makePizza(i)
		// try to make a pizza
		// decision

	}
}

func main() {
	// print out message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message
}
