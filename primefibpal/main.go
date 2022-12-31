package main

import (
	"fmt"
	defs "primefibpal/functions"
)

func main() {

chkflag:
	for {
		var num int

		fmt.Println("1. Check if Number is prime.\n2. Print number of required numbers from fibonaci series\n3. Create a Palindrome.\n Any other Exit.\n Enter your choice")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			//var num int
			fmt.Print("Enter a number to check if it is prime: ")
			fmt.Println("here")
			fmt.Scanln(&num)
			fmt.Println(defs.Prime(num))

		case 2:
			//var numseq int
			fmt.Print("How many numbers you want : ")
			fmt.Scanln(&num)
			defs.Fib(num)

		case 3:
			//var numrev int
			fmt.Print("Enter the number to reverse: ")
			fmt.Scanln(&num)
			fmt.Println(defs.Pali(num))

		default:
			break chkflag

		}

	}
}
