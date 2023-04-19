package assignment2

type ListInterface interface {
	Size() int
	Contains(interface{}) bool
	Add(interface{})
	// AddWithIndex(interface{}, int)
	Remove(interface{}) bool
	Clear()
	Get(int) interface{}
	IndexOf(interface{}) int
	LastIndexOf(interface{}) int
}

// public interface MyList<T> {
//     int size();
//     boolean contains(Object o);
//     void add(T item);
//     void add(T item, int index);
//     boolean remove(T item);
//     T remove(int index);
//     void clear();
//     T get(int index);
//     int indexOf(Object o);
//     int lastIndexOf(Object o);

// }
