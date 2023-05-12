package interfaces

type StringableHashable interface {
	Stringable
	Hashable
	Equals(StringableHashable) bool
}
