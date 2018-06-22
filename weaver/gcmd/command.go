package gcmd

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

var (
	ErrCmdTerminated = errors.New("command terminated")
)

// Execute is a concurrent wrapper around Go's os/exec Output() method.
// It runs a command, and returns its standard output as a byte slice.
// If a long-running command is being executed, it can easily be killed at
// any time using the terminate channel as the process is spawned in a
// Goroutine.
func Execute(c []string, terminate <-chan struct{}) ([]byte, error) {
	cmd := exec.Command(c[0], c[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cout := make(chan []byte, 1)
	cerr := make(chan error, 1)

	go func(cmd *exec.Cmd, cout chan<- []byte, cerr chan<- error) {
		var stderr bytes.Buffer
		cmd.Stderr = &stderr
		out, err := cmd.Output()
		if err != nil {
			cerr <- fmt.Errorf("%+v : %+v", err, stderr.String())
			return
		}
		cout <- out
	}(cmd, cout, cerr)

	select {
	case out := <-cout:
		close(cerr)
		return out, nil
	case err := <-cerr:
		close(cout)
		return nil, err
	case <-terminate:
		log.Println("exiting")
		// if (cmd.ProcessState == nil || cmd.ProcessState.Exited() == false) && cmd.Process != nil {
		if cmd.Process != nil {
			if err := cmd.Process.Kill(); err != nil {
				return nil, err
			}
		}
		return nil, ErrCmdTerminated
	}
}
