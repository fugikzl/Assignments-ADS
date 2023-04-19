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

	person4 := sampletypes.Person{
		Name:   "Vladimir",
		Job:    "Banker",
		Age:    23,
		Gender: true,
	}

	person5 := sampletypes.Person{
		Name:   "Daniyar",
		Job:    "manager",
		Age:    19,
		Gender: true,
	}

	var queueInStore assignment2.ListInterface = assignment2.NewMyArrayList(person1, person2, person3, person4, person5)
	fmt.Println("The name of 2 person is:", queueInStore.Get(1).(sampletypes.Person).Name)
	fmt.Println("\nPosition of Asel in queue is:", queueInStore.IndexOf(person3)+1)

}
