package assignment2

type MyArrayList struct {
	objects []interface{} //field objects is slice of any type
	size    int           //field size with type of int
}

func NewMyArrayList(items ...interface{}) *MyArrayList {
	list := MyArrayList{
		size:    0,
		objects: nil,
	}

	length := len(items)

	if length > 0 {
		list.objects = make([]interface{}, length)
		for i := 0; i < length; i++ {
			list.objects[list.size] = items[i]
			list.size++
		}
	}

	return &list //we're returning pointer cause pointer weights more less when we're passing our struct as a parameter
}

func (list *MyArrayList) Size() int {
	return list.size
}

func (list *MyArrayList) Contains(item interface{}) bool {
	for i := 0; i < list.Size(); i++ {
		if list.objects[i] == item {
			return true
		}
	}
	return false
}

func (list *MyArrayList) Add(item interface{}) {
	list.objects = append(list.objects, item) // list objects equals to new slice of objects where appended
	list.size++
}

func (list *MyArrayList) AddArrayOfItems(items []interface{}) { //we're passing array of items -_-
	list.objects = append(list.objects, items...) // same like previous
	list.size = list.Size() + len(items)
}

func (list *MyArrayList) Get(index int) interface{} {
	if index < list.Size() && index >= 0 {
		return list.objects[index]
	}
	return nil
}

func (list *MyArrayList) Remove(index int) bool {
	if index < list.Size() && index >= 0 {
		list.objects = append(list.objects[:index], list.objects[index+1:]...) // берем крч элементы до индекса и обьеденяем их с элементами после: [a,b,c,e,d] -> remove(2) -> [a,b] + [e,d]
		list.size--
		return true
	}
	return false
}

func (list *MyArrayList) IndexOf(item interface{}) int {
	for i := 0; i < list.size; i++ {
		if item == list.objects[i] {
			return i
		}
	}
	return -1
}

func (list *MyArrayList) LastIndexOf(item interface{}) int {
	for i := list.size - 1; i >= 0; i-- {
		if item == list.objects[i] {
			return i
		}
	}
	return -1
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
