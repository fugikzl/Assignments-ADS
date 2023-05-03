package assignment3

import (
	"Assignments-ADS/sampletypes"
	"fmt"
)

func TestHashTable() {
	t := NewMyHashTable(11)
	for i := 0; i < 10000; i++ {
		var p Stringable = &sampletypes.Person{
			Name:   "John",
			Job:    "sample",
			Age:    48,
			Gender: true,
		}

		var pid StringableHashable = &PersonId{
			Value: uint64(i),
		}

		t.Put(pid, p)
	}

	fmt.Println(t.Get(&PersonId{Value: 5}))
	fmt.Println(t.Get(&PersonId{Value: 9565}))
	fmt.Println(t.Get(&PersonId{Value: 9}))
	fmt.Println(t.Get(&PersonId{Value: 845}))

	fmt.Print("\n------------------------------------------------------\n")

	t.PrintChainArrayElement(0)
	t.PrintChainArrayElement(1)
	t.PrintChainArrayElement(2)
}
