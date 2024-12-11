package main

import (
	"fmt"
)

func isPerfectNumber(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if i*i != n {
				sum += i + n/i
			} else {
				sum += i
			}
		}
	}
	return sum == n && n != 1
}

func printPerfectNumbersInRange(a, b int) {
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}
}

func soalA2() {
	var a, b int
	fmt.Print("")
	fmt.Print("-=-=- Jordan's Perfect Number Program (A2-2) -=-=-\n")
	fmt.Print("First, tolong inputkan dua bilangan positif (a & b) untuk dijadikan range!\n")
	fmt.Print("Second, perlu diingat bahwa a harus lebih kecil atau sama dengan b.\n")
	fmt.Print("\nOke, inputlah range a hingga b: ")
	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Loh, a yang anda input lebih besar dari b. Input ulang!")
		fmt.Print("\n")
		return
	}

	fmt.Printf("Nice, perfect numbers di dalam range [%d, %d] adalah ", a, b)
	printPerfectNumbersInRange(a, b)
	fmt.Print("\n")
}