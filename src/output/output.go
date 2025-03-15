package output

import (
	"os"
	"zash/src/command"
	"zash/src/util"
)

type OutputParams struct {
	Command command.Command
}

var (
	STANDARD_OUTPUT *os.File = os.Stdout
	STANDARD_ERROR *os.File = os.Stderr
)

func withConfiguredOutput(argsPtr *[]string, action func(stdout *os.File, stderr *os.File)) {
	args := *argsPtr
	var fileOutput *os.File
	var fileError *os.File

	if (len(args) >= 2) {
		maybeOperatorIdx := len(args) - 2
		maybeOperator := args[maybeOperatorIdx]
		maybeDestination := args[len(args) - 1]

		var (
			REDIRECT_OUTPUT_OPERATOR = "1>"
			REDIRECT_OUTPUT_SHORTCUT_OPERATOR = ">"
			REDIRECT_ERROR_OPERATOR = "2>"
			APPEND_OUTPUT_OPERATOR = "1>>"
			APPEND_OUTPUT_SHORTCUT_OPERATOR = ">>"
			APPEND_ERROR_OPERATOR = "2>>"
		)

		if util.IsOneOf(maybeOperator, REDIRECT_OUTPUT_OPERATOR, REDIRECT_OUTPUT_SHORTCUT_OPERATOR) {
			fileOutput, _ = os.OpenFile(maybeDestination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			defer fileOutput.Close() // Ensure file is closed when function exits
			
			os.Stdout = fileOutput
		} else if maybeOperator == REDIRECT_ERROR_OPERATOR {
			fileError, _ = os.OpenFile(maybeDestination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			defer fileError.Close() // Ensure file is closed when function exits
			
			os.Stderr = fileError
		} else if util.IsOneOf(maybeOperator, APPEND_OUTPUT_OPERATOR, APPEND_OUTPUT_SHORTCUT_OPERATOR) {
			fileOutput, _ = os.OpenFile(maybeDestination, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			defer fileOutput.Close() // Ensure file is closed when function exits
			
			os.Stdout = fileOutput
		} else if maybeOperator == APPEND_ERROR_OPERATOR {
			fileError, _ = os.OpenFile(maybeDestination, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			defer fileError.Close() // Ensure file is closed when function exits
			
			os.Stderr = fileError
		}

		if fileOutput != nil || fileError != nil {
			*argsPtr = args[:maybeOperatorIdx]
		}
	}

	action(os.Stdout, os.Stderr)

	if fileOutput != nil {
		fileOutput.Close()
	}
	if fileError != nil {
		fileError.Close()
	}

	os.Stdout = STANDARD_OUTPUT
	os.Stderr = STANDARD_ERROR
}

func WriteOutput(params OutputParams) {
	withConfiguredOutput(&params.Command.Args, func(stdout, stderr *os.File) {
		command.HandleCommand(command.HandleCommandParams{
			Command: params.Command,
			Stdout: stdout,
			Stderr: stderr,
		})
	})
}