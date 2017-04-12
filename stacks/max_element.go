package maxelem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myStack []int

// Push pushes a number onto the stack.
func (ms *myStack) Push(pushNum int) {
	*ms = append(*ms, pushNum)
}

// Pop removes the most recent number from the stack.
func (ms *myStack) Pop() int {
	length := len(*ms)
	poppedNum := (*ms)[length-1]
	*ms = (*ms)[:length-1]
	return poppedNum
}

// GetMax finds and returns the largest number in the stack.
func (ms *myStack) GetMax() int {
	compareNum := 0

	for _, v := range *ms {
		if v > compareNum {
			compareNum = v
		}
	}

	return compareNum
}

func main() {
	// Read the first line (aka the stack's size) in from STDIN
	// The first line contains an integer N
	bio := bufio.NewReader(os.Stdin)
	val, _, err := bio.ReadLine()
	if err != nil {
		panic(err)
	}

	size, err := strconv.Atoi(string(val))

	// Initialize a new stack
	ms := make(myStack, size)

	var input string
	var pushNum int

	// Loop through the N lines and perform the appropriate operations
	for i := 0; i < size; i++ {
		// Read in the input query
		// 1 x will push the number x on the stack
		// 2 will delete the number at the top of the stack
		// 3 will print the largest number in the stack
		rawInput, _, err := bio.ReadLine()
		if err != nil {
			panic(err)
		}

		// Get rid of extra spaces at the end of line
		input = strings.TrimRight(string(rawInput), " ")

		inputVals := strings.Split(input, " ")

		// Depending on the first value ...
		switch inputVals[0] {
		case "1":
			// Push the value onto the stack as a number
			pushNum, err = strconv.Atoi(inputVals[1])
			if err != nil {
				panic(err)
			}
			ms.Push(pushNum)
		case "2":
			// Pop the last value from the stack
			ms.Pop()
		case "3":
			// Find and print the max value of the stack
			maxNum := ms.GetMax()
			fmt.Println("the largest number in the stack is ", maxNum)
		}

	}
}
