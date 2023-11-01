package model

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type EmployeeAddress struct {
	EmployeeId int
	AddresId   int
}

type Address struct {
	Id       int
	Address1 string
}
