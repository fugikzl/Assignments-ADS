package assignment3

type StringableHashable interface {
	Stringable
	Hashable
	Equals(StringableHashable) bool
}
