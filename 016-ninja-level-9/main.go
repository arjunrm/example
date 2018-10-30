package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Hello, I am ", p.name)
}

func (p *person) renamePtrReceiver() {
	p.name += "_renamed"
	fmt.Println("I am renamed: ", p.name)
}

func (p person) rename() {
	p.name += "_renamed"
	fmt.Println("I am renamed: ", p.name)
}

type human interface {
	speak()
	renamePtrReceiver()
	rename()
}

func saySomething(h human) {
	fmt.Printf("Hello. Who are you?\n")
	h.speak()
}

func renameHumanPtrReceiver(h human) {
	fmt.Println("\nBefore renameHumanPtrReceiver: ")
	h.speak()
	h.renamePtrReceiver()
	fmt.Println("After renameHumanPtrReceiver: ")
	h.speak()
}

func renameHuman(h human) {
	fmt.Println("\nBefore renameHuman: ")
	h.speak()
	h.rename()
	fmt.Println("After renameHuman: ")
	h.speak()
}

func main() {
	mulGoRoutinesPrintingStr()
	methodSetsExample()
}

func mulGoRoutinesPrintingStr() {
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Println("Hello James: ", x)
			wg.Done()
		}(i)
	}

	for j := 0; j < 10; j++ {
		go func(y int) {
			fmt.Println("Hello Bond: ", y)
			wg.Done()
		}(j)
	}

	fmt.Println("Waiting for go routines to finish")
	wg.Wait()
	fmt.Println("Exit")
}

func methodSetsExample() {
	p1 := person{"James", 23}
	p2 := person{"Bond", 43}
	p3 := &p2

	saySomething(&p1)
	saySomething(&p2)
	saySomething(p3)

	p1.speak()
	p2.speak()
	p3.speak()

	renameHuman(&p1)
	renameHumanPtrReceiver(&p1)
}
