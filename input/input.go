package main

// library to print out msg
import "fmt"

// declare constant setting my favorite color
const favColor string = "blue"

// main function
func main() {

	// init guess var
	var guess string

	// loop for wait the input
	for {

		// get the input
		fmt.Println("Guess my favorite color:")

		// validate if input is null
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// validate if you guess my favorite color
		if favColor == guess {
			fmt.Printf("%q is my favorite color!", favColor)
			return
		}

		fmt.Printf("Sorry, %q is not my favorite color!Guess again.\n", guess)
	}
}
