/*

   gofck.go -

   The language consists of eight commands, listed below. A brainfuck program is a sequence of these commands, possibly interspersed with other characters (which are ignored). The commands are executed sequentially, with some exceptions: an instruction pointer begins at the first command, and each command it points to is executed, after which it normally moves forward to the next command. The program terminates when the instruction pointer moves past the last command.

   The brainfuck language uses a simple machine model consisting of the program and instruction pointer, as well as an array of at least 30,000 byte cells initialized to zero; a movable data pointer (initialized to point to the leftmost byte of the array); and two streams of bytes for input and output (most often connected to a keyboard and a monitor respectively, and using the ASCII character encoding).

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main ... The main function.
func main() {
	parseAndExec(parseInput())
}

// parseAndExec ... Parses the string given and executes the commands sequentially.
func parseAndExec(commands []string) {
	//Pointer acceser for each cell position.
	ptr := 0
	//Typically 30,000
	cellCount := 30000
	//Creating the cells that will be used.
	cells := make([]byte, cellCount)

	execCommands := func(s string) {}
	execCommands = func(s string) {
		for i := 0; i+1 <= len(s); i++ {
			switch string(s[i : i+1]) {
			case ">":
				ptr++
			case "<":
				ptr--
			case "+":
				cells[ptr]++
			case "-":
				cells[ptr]--
			case ".":
				fmt.Printf("%v", string(cells[ptr]))
			case ",":
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Enter Char: ")
				text, _ := reader.ReadString('\n')
				cells[ptr] = []byte(text)[0]
			case "[":
				getInbetween := func(str, start, end string) (string, int) {
					s := strings.Index(str, start)
					s += len(start)
					e := strings.Index(str, end)
					return str[s:e], e
				}

				loopIndex := i
				codeToExec := ""
				codeToExec, loopIndex = getInbetween(s[loopIndex:], "[", "]")

				for cells[ptr] > 0 {
					execCommands(codeToExec)
				}

			case ";":
				fmt.Printf("%v", cells[ptr])
			}
		}
	}

	for i := 0; i+1 <= len(commands); i++ {
		command := commands[i : i+1]

		for _, str := range command {
			execCommands(str)
		}
	}
}

// parseInput ... Just parses the input and returns an array of strings.
func parseInput() []string {
	//Getings the input without the original program call from Command Line.
	args := os.Args[1:]
	return args
}
