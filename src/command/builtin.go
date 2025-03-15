package command

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"zash/src/util"
)

var (
	EXIT string = "exit"
	ECHO string = "echo"
	TYPE string = "type"
	PWD string = "pwd"
	CD string = "cd"
)

var BUILT_IN_COMMANDS = []string{EXIT, ECHO, TYPE, PWD, CD}

func isBuiltinCommand(command string) bool {
	return slices.Contains(BUILT_IN_COMMANDS, command)
}

func handleExit(args []string) {
	if len(args) == 0 || args[0] != "0" {
		fmt.Println("Invalid args for command: ", args)
	} else {
		os.Exit(0)
	}
}

func handleEcho(args []string) {
	const SPACE = " "

	for i := 0; i < len(args); i++ {
		cur := args[i]
		
		fmt.Print(cur)
		
		if i + 1 < len(args) && args[i + 1] != SPACE && cur != " " {
			fmt.Print(" ")
		}
	}
	
	fmt.Print("\n")
}

func handleType(args []string) {
	for _, command := range args {
		if isBuiltinCommand(command) {
			fmt.Println(command + " is a shell builtin")
		} else if executablePath := util.FindExecutablePath(command); executablePath != nil {
			fmt.Println(command + " is " + *executablePath)
		} else {
			fmt.Println(command + ": not found")
		}
	}
}

func handlePwd(args []string) {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error handling command: ", err)
		return
	}

	fmt.Println(dir)
}

func handleCd(args []string) {
	homeDir, _ := os.UserHomeDir()

	var dir string

	if len(args) > 0 {
		dir = args[0]
	}

	dir = strings.ReplaceAll(dir, "~", homeDir)
	dir = strings.ReplaceAll(dir, "//", "/")

	err := os.Chdir(dir)

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dir)
		return
	}
}

var BUILTIN_COMMAND_HANDLERS = map[string]func([]string) {
	EXIT: handleExit,
	ECHO: handleEcho,
	TYPE: handleType,
	PWD: handlePwd,
	CD: handleCd,
}

type BuiltinCommandHandler struct {
}

func (handler *BuiltinCommandHandler) shouldHandle(params HandleCommandParams) bool {
	return isBuiltinCommand(params.Command.Name)
}

func (handler *BuiltinCommandHandler) handle(params HandleCommandParams) {
	commandHandler, commandHandlerExists := BUILTIN_COMMAND_HANDLERS[params.Command.Name]
			
	if (commandHandlerExists) {
		commandHandler(params.Command.Args)
	}
}