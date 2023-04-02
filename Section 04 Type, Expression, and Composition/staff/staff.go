package staff

var OverPaidLimit = 30000
var underPaidLimit = 20000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (e *Office) All() []Employee {
	return e.AllStaff
}

func (e *Office) Overpaid() []Employee {
	var overpaid []Employee
	for _, v := range e.AllStaff {
		if v.Salary > OverPaidLimit {
			overpaid = append(overpaid, v)
		}
	}
	return overpaid
}

func (e *Office) Underpaid() []Employee {
	var underpaid []Employee
	for _, v := range e.AllStaff {
		if v.Salary < underPaidLimit {
			underpaid = append(underpaid, v)
		}
	}
	return underpaid
}
