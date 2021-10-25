package main

import ( 
	"fmt"
	"log"
)


type Entertainer interface {
	Entertain() string
	Perform() string
	
}

type Setuper interface {
	Setup() string
}

type Teardowner interface{
	Teardown() string
}

type AllThree interface {
	Entertainer
	Setuper
	Teardowner 
}

type venue struct {
	Log
}

func (v venue) Entertain() string {
	return "%s has performed for %d people.\n"
}

func (v venue) Setup() string {
	return "%s has completed setup.\n"
}

func (v venue) Teardown() string {
	return "%s has completed teardown.\n"
} 

func process(venues  AllThree) {
	fmt.Println("Test....", venues.Entertain)
	fmt.Println("Test....", venues.Setup)
	fmt.Println("Test....", venues.Teardown)
}

func main() {
	perform := venue{}
	process(perform)
	
}

