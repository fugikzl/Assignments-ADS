package main

import (
	"Assignments-ADS/assignment2"
	"Assignments-ADS/sampletypes"
	"fmt"
)

func main() {

	person1 := sampletypes.Person{
		Name:   "Nariman",
		Job:    "casier",
		Age:    35,
		Gender: true,
	}

	person2 := sampletypes.Person{
		Name:   "Madiyar",
		Job:    "journalist",
		Age:    21,
		Gender: true,
	}

	person3 := sampletypes.Person{
		Name:   "Asel",
		Job:    "python developer",
		Age:    11,
		Gender: false,
	}

	arrayList := assignment2.NewMyArrayList(person1, person2, person3)
	fmt.Println("The name of 2 person is:", arrayList.Get(1).(sampletypes.Person).Name)
	fmt.Println("The name of 2 person is:", arrayList.Get(1).(sampletypes.Person).Name)

}
