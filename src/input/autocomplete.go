package input

import (
	"fmt"
	"sort"
	"strings"
)

func getLongestCommonPrefix(strings []string) string {
	prefix := strings[0]

	for _, str := range strings[1:] {
		// Compare the prefix with each match
		i := 0
		for i < len(prefix) && i < len(str) && prefix[i] == str[i] {
			i++
		}
		prefix = prefix[:i]
	}

	return prefix
}

func getMatchingCommands(input string, commands []string) []string {
	var matches []string
	
	for _, cmd := range commands {
		if strings.HasPrefix(cmd, input) {
			matches = append(matches, cmd)
		}
	}

	if len(matches) == 0 {
		return matches
	}

	if len(matches) == 1 {
		return []string{matches[0] + " "}
	}

	sort.Strings(matches)

	prefix := getLongestCommonPrefix(matches)

	// If the prefix length is less than the original input, just return the matches
	if len(prefix) > len(input) {
		return []string{prefix}
	}
	
	return matches
}

func GetMatchingCommands(input string, commands []string) []string {
	return getMatchingCommands(input, commands)
}

func ringBell() {
	fmt.Print("\a") // Rings a bell sound
}

func handleAutocomplete(params *InputParams, lastChar byte) {
	matches := getMatchingCommands(string(params.input), params.Commands)

	// single match found
	if len(matches) == 1 {
		// Auto-complete with the single match
		fmt.Print(matches[0][len(params.input):])
		params.input = []rune(matches[0])
		return
	}

	// multiple matches found
	if len(matches) > 1 && lastChar == TAB {
		// Show possible completions
		fmt.Println("\n" + strings.Join(matches, "  "))
		fmt.Print(params.Prompt + string(params.input))
		return
	}

	// No match found, ring the terminal bell
	ringBell()
}