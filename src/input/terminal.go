package input

import (
	"fmt"
	"syscall"
	"unsafe"
	"zash/src/constants"
)

var IOCTL_CONSTANTS = constants.GetIoctlConstants()

// enableRawMode sets the terminal to "raw mode" for real-time input handling.
// Raw mode disables line buffering and echoing, meaning characters are read
// immediately as the user types them instead of waiting for Enter.
func enableRawMode(fd int) (*syscall.Termios, error) {
	var oldState syscall.Termios
	

	// Get the current terminal attributes
	_, _, err := syscall.Syscall6(
		syscall.SYS_IOCTL,               // System call for terminal I/O control
		uintptr(fd),                     // File descriptor (stdin)
		uintptr(IOCTL_CONSTANTS.GetAttribute),       // Command to get terminal attributes (macOS)
		uintptr(unsafe.Pointer(&oldState)), // Pointer to store the old settings
		0, 0, 0,
	)
	if err != 0 {
		return nil, fmt.Errorf("failed to get terminal state")
	}

	newState := oldState

	// Modify the terminal attributes:
	// - ICANON: Disables canonical mode (input is processed immediately, not line by line)
	// - ECHO: Prevents characters from being echoed to the screen
	newState.Lflag &^= syscall.ICANON | syscall.ECHO

	// Apply the new settings
	_, _, err = syscall.Syscall6(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(constants.GetIoctlConstants().SetAttribute),       // Command to set terminal attributes (macOS)
		uintptr(unsafe.Pointer(&newState)),
		0, 0, 0,
	)
	if err != 0 {
		return nil, fmt.Errorf("failed to set raw mode")
	}

	return &oldState, nil // Return old settings so they can be restored later
}

// restoreTerminal restores the terminal to its original settings
// after the program finishes, ensuring normal behavior is restored.
func restoreTerminal(fd int, state *syscall.Termios) {
	if state != nil {
		syscall.Syscall6(
			syscall.SYS_IOCTL,
			uintptr(fd),
			uintptr(constants.GetIoctlConstants().SetAttribute),       // Restore original terminal settings
			uintptr(unsafe.Pointer(state)),
			0, 0, 0,
		)
	}
}
