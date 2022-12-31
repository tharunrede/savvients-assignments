package defs

import "fmt"

func Prime(n int) bool {

	count := 0

	for i := 2; i < n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 0 {
		return true
	}
	return false
}

func Fib(n int) {

	f := 0

	s := 1

	fmt.Println(f)
	fmt.Println(s)
	next := f + s
	//fmt.Println("nxt,n", next, n)
	for i := 1; i <= n-2; i++ {
		fmt.Println(next)
		//fmt.Println("control")
		f = s
		s = next
		next = f + s

	}

}

func Pali(n int) int {

	quo := n / 10
	rem := n % 10
	finl := 0
	//fmt.Println("quo  rem finl out", quo, rem, finl)
	for {

		if quo == 0 {
			break
		} else {
			finl = (finl * 10) + rem
			//fmt.Println("quo  rem finl insde before", quo, rem, finl)
			rem = quo % 10
			quo = quo / 10

			//fmt.Println("quo  rem finl insde after", quo, rem, finl)

		}

	}
	finl = (finl * 10) + rem
	return finl

}
