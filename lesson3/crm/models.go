package crm

type Address struct {
	City string
	Street string
}

type User struct{
	Name string
	Address Address
}
