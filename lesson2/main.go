package main

import "fmt"

func main() {
	variadicParams()
}

func variadicArgs() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := Average_args(data...) // Average_args(data...) ------> Average_args(43, 56, 87, 12, 45, 57) demektir. Listeyi açıp tek tek veriyi gönderiyorsun. Slice'ın tamamını paket halinde değil
	fmt.Println(n)
}

func variadicParams() {
	n := Average_params(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}
