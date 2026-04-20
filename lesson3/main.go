package main

import (
	"lesson3/crm"
	"fmt"
)

func main(){
	a := crm.Address{City: "Istanbul", Street: "Erzincan"}
	u := crm.User{Name: "Cenk", Address: a}

	fmt.Println(u.Name, u.Address)
}