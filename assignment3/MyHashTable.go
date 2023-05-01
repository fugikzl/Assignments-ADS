package assignment3

type MyHashTable struct {
	M    int
	size int
}

func NewMyHashTable(M int) MyHashTable {
	table := MyHashTable{
		M:    M,
		size: 0,
	}

	return table
}
