//go:build windows

package tmux

// Signal type for Windows compatibility
type windowsSignal int

// killProcessGroup is a no-op on Windows since tmux is not supported.
func killProcessGroup(pgid int, sig windowsSignal) error {
	return nil
}

// Signals for process termination (no-op on Windows)
const (
	sigTERM windowsSignal = 15
	sigKILL windowsSignal = 9
)
