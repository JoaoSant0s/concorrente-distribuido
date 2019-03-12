package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type Philosopher struct {
	name string
	availableFork chan bool
	neighbor *Philosopher
}

func main() {	
	philosophers := createPhilosophers()
	finishedEating := make(chan *Philosopher)	

	startDinner(philosophers, finishedEating)
}

func startDinner(philosophers []*Philosopher, finishedEating chan *Philosopher){
	fmt.Println("----  Start Dinner ----")

	for _, philo := range philosophers {
		go philo.dine(finishedEating)
	}
	philosLen := len(philosophers)
	for i := 0; i < philosLen; i++ {
		philo := <- finishedEating
		fmt.Println(philo.name, "- ate.")
	}
	
	fmt.Println("----  the " + strconv.Itoa(philosLen) + " philosophers ate ----\n")
	startDinner(philosophers, finishedEating)
}

func createPhilosophers()[]*Philosopher{
	names := []string{"philosopher 1", "philosopher 2", "philosopher 3", "philosopher 4", "philosopher 5"}

	philosophers := make([]*Philosopher, len(names))
	var philo *Philosopher
	for i, name := range names {
		philo = createPhilosopher(name, philo)
		philosophers[i] = philo
	}
	philosophers[0].neighbor = philo
	return philosophers
}

func createPhilosopher(name string, neighbor *Philosopher) *Philosopher {
	philo := &Philosopher{name, make(chan bool, 1), neighbor}
	philo.availableFork <- true
	return philo
}

func (philo *Philosopher) dine(finishedEating chan *Philosopher) {
	philo.think()
	philo.getForks()
	philo.eat()
	philo.returnforks()	
	finishedEating <- philo
}

func (philo *Philosopher) think() {
	fmt.Println(philo.name, "- thinking.")
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (philo *Philosopher) getForks() {
	timeout := make(chan bool, 1)
	go func() { time.Sleep(1e9); timeout <- true }()
	<-philo.availableFork
	fmt.Println(philo.name, "- fork 1.")
	select {
		case <-philo.neighbor.availableFork:		
			fmt.Println(philo.name, "- fork 2.")
			return
		case <-timeout:
			fmt.Println(philo.name, "- release fork.")
			philo.availableFork <- true
			philo.think()
			philo.getForks()
	}
}

func (philo *Philosopher) eat() {
	fmt.Println(philo.name, "- eating.")
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (philo *Philosopher) returnforks() {
	philo.availableFork <- true
	philo.neighbor.availableFork <- true
}
