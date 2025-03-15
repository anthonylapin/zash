package preprocess

import "strings"

const (
	BACKSLASH = '\\'
	DOUBLE_QUOTE = '"'
	SINGLE_QUOTE = '\''
	SPACE = ' '
)

func parseArgs(input string) []string {
	var (
		args []string
		current strings.Builder
		inQuotes bool
		quoteChar byte
		i int
	)

	for i < len(input) {
		char := input[i]

		if char == BACKSLASH && i + 1 < len(input) {
			// Handle escaped characters
			i++
			next := input[i]

			if inQuotes && next != quoteChar && next != BACKSLASH {
				current.WriteByte(char)
			}

			current.WriteByte(next)
		} else if char == DOUBLE_QUOTE || char == SINGLE_QUOTE {
			// Handle opening/closing quotes
			if inQuotes && quoteChar == char {
				inQuotes = false
			} else if !inQuotes {
				inQuotes = true
				quoteChar = char
			} else {
				current.WriteByte(char)
			}
		} else if char == SPACE && !inQuotes {
			// Handle spaces outside quotes
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		} else {
			current.WriteByte(char)
		}
		i++
	}
	
	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}