//go:build !windows

package tmux

import "syscall"

// killProcessGroup sends a signal to a process group.
// pgid should be positive; it will be negated internally.
func killProcessGroup(pgid int, sig syscall.Signal) error {
	return syscall.Kill(-pgid, sig)
}

// Signals for process termination
const (
	sigTERM = syscall.SIGTERM
	sigKILL = syscall.SIGKILL
)
