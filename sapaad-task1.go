//Print first name if the number is even else print second name
package main

import "fmt"

func main() {
	var first_name = "Akhil"
	var last_name = "G Krishnan"
	var odd_count = 0
	var even_count = 0
	for a := 1; a <= 100; a++ {
		if a%2 == 0 {
			fmt.Printf("Even :%s", first_name)
			fmt.Printf("\n")
			odd_count++
		} else {
			fmt.Printf("Odd :%s", last_name)
			fmt.Printf("\n")
			even_count++
		}

	}
	fmt.Printf("%d :%d", odd_count, even_count)

}
