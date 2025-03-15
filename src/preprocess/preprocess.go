package preprocess

import "zash/src/command"

func preprocessCommand(commandStr string) (string, []string) {
	if len(commandStr) == 0 {
		return "", []string{}
	}

	args := parseArgs(commandStr)
	
	if len(args) == 0 {
		return "", []string{}
	}

	return args[0], args[1:]
}

func PreprocessCommand(commandStr string) command.Command {
	commandName, args := preprocessCommand(commandStr)

	return command.Command{
		Name: commandName,
		Args: args,
	}
}