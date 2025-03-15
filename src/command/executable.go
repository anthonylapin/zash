package command

import (
	"io"
	"os/exec"

	"zash/src/util"
)

type ExecutableHandler struct{}

func runCommand(executablePath string, args []string, stdout io.Writer, stderr io.Writer) {
	cmd := exec.Command(executablePath, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	cmd.Run()
}

func (handler *ExecutableHandler) shouldHandle(params HandleCommandParams) bool {
	return util.FindExecutablePath(params.Command.Name) != nil
}

func (handler *ExecutableHandler) handle(params HandleCommandParams) {
	runCommand(params.Command.Name, params.Command.Args, params.Stdout, params.Stderr)
}