package assignment2

type MyArrayList struct {
	Objects  []interface{} //field Objects is slice of any type
	ListSize int           //field Size with type of int
}

func NewMyArrayList(items ...interface{}) *MyArrayList {
	list := MyArrayList{
		ListSize: 0,
		Objects:  nil,
	}

	length := len(items)

	if length > 0 {
		list.Objects = make([]interface{}, length)
		for i := 0; i < length; i++ {
			list.Objects[list.ListSize] = items[i]
			list.ListSize++
		}
	}

	return &list //we're returning pointer cause pointer weights more less when we're passing our struct as a parameter
}

func (list *MyArrayList) Size() int {
	return list.ListSize
}

func (list *MyArrayList) Contains(item interface{}) bool {
	for i := 0; i < list.Size(); i++ {
		if list.Objects[i] == item {
			return true
		}
	}
	return false
}

func (list *MyArrayList) Add(item interface{}) {
	list.Objects = append(list.Objects, item) // list Objects equals to new slice of Objects where appended
	list.ListSize++
}

func (list *MyArrayList) AddArrayOfItems(items []interface{}) { //we're passing array of items -_-
	list.Objects = append(list.Objects, items...) // same like previous
	list.ListSize = list.Size() + len(items)
}

func (list *MyArrayList) Get(index int) interface{} {
	if index < list.Size() && index >= 0 {
		return list.Objects[index]
	}
	return nil
}

func (list *MyArrayList) Remove(item interface{}) bool {
	for i := 0; i < list.ListSize; i++ {
		if item == list.Objects[i] {
			list.Objects = append(list.Objects[:i], list.Objects[i+1:]...)
			list.ListSize--
			return true
		}
	}
	return false
}

func (list *MyArrayList) RemoveFomIndex(index int) bool {
	if index < list.ListSize && index >= 0 {
		list.Objects = append(list.Objects[:index], list.Objects[index+1:]...)
		// берем крч элементы до индекса и обьеденяем их с элементами после: [a,b,c,e,d] -> remove(2) -> [a,b] + [e,d]
		list.ListSize--
		return true
	}
	return false
}

func (list *MyArrayList) IndexOf(item interface{}) int {
	for i := 0; i < list.ListSize; i++ {
		if item == list.Objects[i] {
			return i
		}
	}
	return -1
}

func (list *MyArrayList) LastIndexOf(item interface{}) int {
	for i := list.ListSize - 1; i >= 0; i-- {
		if item == list.Objects[i] {
			return i
		}
	}
	return -1
}

func (list *MyArrayList) AddWithIndex(item interface{}, index int) {
	if index <= list.ListSize && index >= 0 {
		list.Objects = append(list.Objects[:index], item, list.Objects[index:])
	}

	if index > list.ListSize && index >= 0 {
		list.Objects = append(list.Objects, item)
	}
}

func (list *MyArrayList) Clear() {
	list.Objects = nil
	list.ListSize = 0
}

// type ListInterface interface {
// 	Size() int
// 	Contains(interface{}) bool
// 	Add(interface{})
// 	AddWithIndex(interface{}, int)
// 	Remove(interface{}) bool
// 	Clear()
// 	Get(int) interface{}
// 	IndexOf(interface{}) int
// 	LastIndexOf(interface{}) int
// }
