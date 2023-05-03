package sampletypes

type Person struct {
	Name   string
	Job    string
	Age    uint8
	Gender bool // true - for male and false for female
}

func (person *Person) String() string {
	return person.Name + " " + person.Job + " " + person.ageToString() + " " + person.genderToString()
}

func (person *Person) genderToString() string {
	if person.Gender {
		return "M"
	}
	return "F"
}

func (person *Person) ageToString() string {
	age := person.Age
	res := ""
	for age/10 != 0 {
		res = string(age%10+48) + res
		age = age / 10
	}
	res = string(age+48) + res
	return res
}
