package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Greet the user
	fmt.Println("Hi Bharathi Bhumireddy, I am a calculator app ....")

	for {
		// Create a reader to read user input from stdin
		reader := bufio.NewReader(os.Stdin)

		// Prompt the user for input
		fmt.Print("Enter any calculation (Example: 1 + 2 or 2 * 5 -> Please maintain spaces as shown in example): ")
		text, _ := reader.ReadString('\n')

		// Remove the newline character from input
		text = strings.TrimSpace(text)

		// Exit condition
		if text == "exit" {
			break
		}

		// Split the input into operands and operator
		parts := strings.Split(text, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input format. Please enter like: 1 + 2")
			continue
		}

		// Convert the first operand to integer
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid first number. Try again.")
			continue
		}

		// Convert the second operand to integer
		right, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Invalid second number. Try again.")
			continue
		}

		// Perform calculation based on operator
		var result int
		switch parts[1] {
		case "+":
			result = left + right
		case "-":
			result = left - right
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				fmt.Println("Cannot divide by zero.")
				continue
			}
			result = left / right
		default:
			fmt.Println("Invalid operator. Please use +, -, *, or /.")
			continue
		}

		// Display the result
		fmt.Printf("Result: %d\n", result)
	}
}
