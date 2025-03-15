package command

import "fmt"

type InvalidCommandHandler struct{}

func (handler *InvalidCommandHandler) shouldHandle(params HandleCommandParams) bool {
	return true
}

func (handler *InvalidCommandHandler) handle(params HandleCommandParams) {
	fmt.Println(params.Command.Name + ": command not found")
}