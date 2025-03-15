package command

import "os"

type Command struct {
	Name string
	Args []string
}

type HandleCommandParams struct {
	Command Command
	Stdout *os.File
	Stderr *os.File
}

type CommandHandler interface {
	handle(params HandleCommandParams)
	shouldHandle(params HandleCommandParams) bool
}

var COMMAND_HANDLERS = []CommandHandler {
	&BuiltinCommandHandler{},
	&ExecutableHandler{},
	&InvalidCommandHandler{},
}

func HandleCommand(params HandleCommandParams) {
	for _, handler := range COMMAND_HANDLERS {
		if handler.shouldHandle(params) {
			handler.handle(params)
			break
		}
	}
}