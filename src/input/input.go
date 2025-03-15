package input

import (
	"fmt"
	"os"
	"syscall"
)

const (
	TAB byte = '\t'
	NEW_LINE byte = '\n'
	BACKSPACE_1 = 127
	BACKSPACE_2 = 8
)

type InputParams struct {
	Prompt string
	Commands []string
	input []rune
}

func readByte(fileDescriptor int) (byte, error) {
	var buf [1]byte
	n, err := syscall.Read(fileDescriptor, buf[:])

	ch := buf[0]

	if err != nil || n == 0 {
		return ch, fmt.Errorf("error reading input")
	}

	return ch, nil
}

func handleBackspace(params *InputParams) {
	if len(params.input) > 0 {
		params.input = params.input[:len(params.input) - 1]
		fmt.Print("\b \b") // Move cursor back, erase character, move back again
	}
}

func appendToInput(ch byte, params *InputParams) {
	// Append character to input and print it on the screen
	fmt.Print(string(ch))
	params.input = append(params.input, rune(ch))
}

func buildInput(params *InputParams) string {
	fmt.Println()
	return string(params.input)
}

func ReadInput(params InputParams) string {
	fmt.Print(params.Prompt)

	inputFileDescriptor := int(os.Stdin.Fd())

	// Enable raw mode for real-time key capture
	oldState, err := enableRawMode(inputFileDescriptor)
	if err != nil {
		fmt.Println("Failed to set raw mode:", err)
		return ""
	}
	defer restoreTerminal(inputFileDescriptor, oldState) // Ensure terminal is restored when function exits

	var lastChar byte

	for {
		char, err := readByte(inputFileDescriptor) // read character from stdin

		if err != nil {
			break
		}

		switch char {
		case TAB: // Handle TAB key for autocomplete
			handleAutocomplete(&params, lastChar)
		case NEW_LINE: // Handle Enter key (submit input)
			return buildInput(&params)
		case BACKSPACE_1, BACKSPACE_2: // Handle Backspace key
			handleBackspace(&params)
		default:
			appendToInput(char, &params)
		}

		lastChar = char
	}
	return ""
}