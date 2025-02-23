package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Diner struct {
	number 	int
	left	*chan bool
	right	*chan bool
}

func (self Diner) diner_loop() {
	time.After(time.Duration(rand.IntN(5)) * time.Second)

	for {
		select {
		case <-*self.left:
			fmt.Printf("Diner %d: acquired left chopstick\n", self.number)

			select	{
			case <-*self.right:
				fmt.Printf("Diner %d: Acquired both chopsticks!\n", self.number)
				time.After(time.Duration(rand.IntN(10)) * time.Second)

				fmt.Printf("Diner %d: releasing....\n", self.number)
				*self.right <- true
				*self.left <- true
			case <-time.After(3 * time.Second):
				fmt.Printf("Diner %d, failed to acquire right chopstick, releasing left\n", self.number)
				*self.left <- true
			}
		case <-*self.right:
			fmt.Printf("Diner %d: acquired right chopstick\n", self.number)

			select {
			case <-*self.left:
				fmt.Printf("Diner %d: Acquired both chopsticks!\n", self.number)
				time.After(time.Duration(rand.IntN(10)) * time.Second)

				fmt.Printf("Diner %d: releasing....\n", self.number)
				*self.left <- true
				*self.right <- true
			case <-time.After(3 * time.Second):
				fmt.Printf("Diner %d: failed to acquire left chopstick, releasing right\n", self.number)
				*self.right <- true
			}
		}
	}
}

func main() {
	c1 := make(chan bool, 1)
	c1 <- true	// unlocked
	c2 := make(chan bool, 1)
	c2 <- true
	c3 := make(chan bool, 1)
	c3 <- true
	c4 := make(chan bool, 1)
	c4 <- true
	c5 := make(chan bool, 1)
	c5 <- true

	d1 := Diner{ number: 1, left: &c5, right: &c1 }
	d2 := Diner{ number: 2, left: &c1, right: &c2 }
	d3 := Diner{ number: 3, left: &c2, right: &c3 }
	d4 := Diner{ number: 4, left: &c3, right: &c4 }
	d5 := Diner{ number: 5, left: &c4, right: &c5 }
	
	go d1.diner_loop()
	go d2.diner_loop()
	go d3.diner_loop()
	go d4.diner_loop()
	go d5.diner_loop()

	for{}
}
