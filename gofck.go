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
)

var brainFuck *brainFuckProgram

const (
	//INCREMENT_POINTER ... Increments the pointer.
	INCREMENT_POINTER = iota
	//DECREMENT_POINTER ... Decrements the pointer.
	DECREMENT_POINTER
	//INCREMENT_CELL ... Increments the cell value.
	INCREMENT_CELL
	//DECREMENT_CELL ... Decrements the cell value.
	DECREMENT_CELL
	//OUTPUT_VALUE ... Outputs the value of the cell at the selected pointer.
	OUTPUT_VALUE
	//INPUT_VALUE ... Takes a character input and stores that value into the cell.
	INPUT_VALUE
	//JUMP_LOOP_START ... Run code between [] if cell is non-zero.
	JUMP_LOOP_START
	//JUMP_LOOP_END ... Jump back to the [ if cell is non-zero.
	JUMP_LOOP_END
)

type brainFuckProgram struct {
	//Easier for my brain to read \-(.-.)-/
	cells [30 * 1000]byte
	ptr   int
}

func init() {
	brainFuck = new(brainFuckProgram)
	brainFuck.ptr = 0
}

// main ... The main function.
func main() {
	parseCommands(parseInput())
}

// parseAndExec ... Parses the string given and executes the commands sequentially.
func parseCommands(commands []string) {
	for _, command := range commands {
		for _, val := range []byte(command) {
			if val == '>' {
				execOpCode(INCREMENT_POINTER)
			}

			if val == '<' {
				execOpCode(DECREMENT_POINTER)

			}

			if val == '+' {
				execOpCode(INCREMENT_CELL)
			}

			if val == '-' {
				execOpCode(DECREMENT_CELL)
			}

			if val == '.' {
				execOpCode(OUTPUT_VALUE)
			}

			if val == ',' {
				execOpCode(INPUT_VALUE)
			}

			if val == '[' {
				execOpCode(JUMP_LOOP_START)
			}

			if val == ']' {
				execOpCode(JUMP_LOOP_END)
			}
		}
	}
}

func execOpCode(opCode byte) {
	cells := &brainFuck.cells
	ptr := &brainFuck.ptr
	switch opCode {
	case INCREMENT_POINTER:
		*ptr++
	case DECREMENT_POINTER:
		*ptr--
	case INCREMENT_CELL:
		cells[*ptr]++
	case DECREMENT_CELL:
		cells[*ptr]--
	case OUTPUT_VALUE:
		fmt.Printf("%v", string(cells[*ptr]))
	case INPUT_VALUE:
		reader := bufio.NewReader(os.Stdin)
		val, _ := reader.ReadString('\n')
		cells[*ptr] = ([]byte(val))[0]
	case JUMP_LOOP_START:

	case JUMP_LOOP_END:
	}
}

// parseInput ... Just parses the input and returns an array of strings.
func parseInput() []string {
	//Getings the input without the original program call from Command Line.
	args := os.Args[1:]
	return args
}
