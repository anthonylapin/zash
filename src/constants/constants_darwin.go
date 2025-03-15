//go:build darwin

package constants

import "syscall"

// macOS-specific IOCTL constants
func GetIoctlConstants() IoctlConstants {
	return IoctlConstants{
		GetAttribute: int(syscall.TIOCGETA),
		SetAttribute: int(syscall.TIOCSETA),
	}
}