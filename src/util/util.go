package util

import (
	"os"
	"path/filepath"
	"strings"
)

func IsOneOf[T comparable](value T, options ...T) bool {
	for _, opt := range options {
		if value == opt {
			return true
		}
	}
	return false
}

func GetExecutablePaths() []string {
	return strings.FieldsFunc(os.Getenv("PATH"), func (r rune) bool {
		return r == os.PathListSeparator
	})
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	const EXECUTE_PERMISSION = 0111

	// Check execute permission for owner, group, or others
	return info.Mode().Perm() & EXECUTE_PERMISSION != 0
}

func GetExecutableCommands() []string {
	executables := []string{}

	for _, dir := range GetExecutablePaths() {
		files, err := os.ReadDir(dir)

		if err != nil {
			continue // skip dir that is not possible to read
		}

		for _, file := range files {
			if !file.IsDir() {
				fullPath := filepath.Join(dir, file.Name())

				// Check if the file is executable
				if isExecutable(fullPath) {
					executables = append(executables, file.Name())
				}
			}
		}
	}

	return executables
}

func FindExecutablePath(command string) *string {
	for _, p := range GetExecutablePaths() {
		pathToCommand := p + "/" + command

		info, err := os.Stat(pathToCommand)

		if err == nil && !info.IsDir() {
			return &pathToCommand
		}
	}

	return nil
}
