package main

import (
	"Assignments-ADS/sampletypes"
	"fmt"

	// "fmt"
	"Assignments-ADS/assignment3"
)

func main() {

	// person1 := sampletypes.Person{
	// 	Name:   "Nariman",
	// 	Job:    "casier",
	// 	Age:    35,
	// 	Gender: true,
	// }

	t := assignment3.NewMyHashTable(11)
	for i := 0; i < 10000; i++ {
		var p assignment3.Stringable = &sampletypes.Person{
			Name:   "John",
			Job:    "sample",
			Age:    48,
			Gender: true,
		}

		var pid assignment3.StringableHashable = &sampletypes.PersonId{
			Value: uint64(i),
		}

		t.Put(pid, p)
	}

	fmt.Println(t.Get(&sampletypes.PersonId{Value: 1}))

}
