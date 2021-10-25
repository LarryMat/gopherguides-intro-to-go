package main

import ( 
	"fmt"
	
)


type Entertain interface {
	Perform() 
	
}


type Venue struct {
	Audience int
	Entertainer string

}

func (v *Venue) Perform() {
	fmt.Printf("%s has performed for %d people.\n",  v.Entertainer, v.Audience)
	
}


func main() {
	v := Venue{Audience: 65, Entertainer: "bob"}
	entertaining(v)
	
}

func entertaining(e Entertain) {
	e.Perform()
}