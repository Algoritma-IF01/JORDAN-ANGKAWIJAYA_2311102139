package main

import (
	"bufio"   // this package is used to read input from the user
	"fmt"     // this package is used to print to the console
	"os"      // this package is used to read input from the user
	"strings" // this package is used to manipulate strings
)

func main() {
	reader := bufio.NewReader(os.Stdin) // create a reader to read input from the user
	fmt.Println("Choose a function to run:")
	fmt.Println("1: soalA1()")
	fmt.Println("2: soalA2()")
	fmt.Println("3: soalA3()")
	fmt.Print("Enter choice (1/2/3): ")

	choice, err := reader.ReadString('\n') // read input from the user
	if err != nil {                        // if there is an error reading the input
		fmt.Println("Error reading input:", err) // print the error
		return
	}

	choice = strings.TrimSpace(choice) // remove leading/trailing whitespace

	switch choice {
	case "1":
		soalA1()
	case "2":
		soalA2()
	case "3":
		soalA3()

	default:
		fmt.Println("Invalid choice!")
	}
}
