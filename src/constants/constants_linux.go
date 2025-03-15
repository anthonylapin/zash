//go:build linux

package constants

import "syscall"

// Linux-specific IOCTL constants
func GetIoctlConstants() IoctlConstants {
	return IoctlConstants{
		GetAttribute: int(syscall.TCGETS),
		SetAttribute: int(syscall.TCSETS),
	}
}