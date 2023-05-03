package sampletypes

import "math"

type PersonId struct {
	Value uint64
}

func (person *PersonId) String() string {
	id := person.Value
	res := ""
	for id/10 != 0 {
		res = string(uint8(id%10)+48) + res
		id = id / 10
	}
	res = string(uint8(id)+48) + res
	return res

}

//Knuth's multiplicative method:
//https://stackoverflow.com/questions/664014/what-integer-hash-function-are-good-that-accepts-an-integer-hash-key
func (p *PersonId) HashCode() int {
	var hashedValue int64 = int64(p.Value) * 2654435761
	return int(math.Abs(float64(hashedValue ^ (hashedValue >> 32))))
}
