package main

import "fmt"

func Average_params(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}
