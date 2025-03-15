package main

import (
	"zash/src/command"
	"zash/src/input"
	"zash/src/output"
	"zash/src/preprocess"
	"zash/src/util"
)

func startREPL() {
	allCommands := append(command.BUILT_IN_COMMANDS, util.GetExecutableCommands()...)
	uniqueCommands := util.NewSet(allCommands...).ToSlice()

	for {
		input := input.ReadInput(input.InputParams{
			Prompt: "$ ",
			Commands: uniqueCommands,
		})

		command := preprocess.PreprocessCommand(input)

		output.WriteOutput(output.OutputParams{
			Command: command,
		})
	}
}

func main() {
	startREPL()
}
