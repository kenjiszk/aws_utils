package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"syscall"
)

const (
	stdoutColor = "green"
	stderrColor = "red"
)

type SSHInfo struct {
	User    string
	Host    string
	KeyPath string
	Result  string
}

func (s *SSHInfo) execRemoteCommand(remoteCommand string) error {
	sshHostString := fmt.Sprintf("%s@%s", s.User, s.Host)
	command := exec.Command("ssh", sshHostString, "-i", s.KeyPath, remoteCommand)
	stdout, _, exitCode, err := runCommand(command)
	if err != nil {
		return err
	}
	if exitCode != 0 {
		return errors.New(fmt.Sprintf("Failed to execute %s. exitCode: %d", command, exitCode))
	}
	s.Result = stdout
	return nil
}

// Copy from https://github.com/hnakamur/execcommandexample/blob/master/main.go
func runCommand(cmd *exec.Cmd) (stdout, stderr string, exitCode int, err error) {
	outReader, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	errReader, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	var bufout, buferr bytes.Buffer
	outReader2 := io.TeeReader(outReader, &bufout)
	errReader2 := io.TeeReader(errReader, &buferr)

	if err = cmd.Start(); err != nil {
		return
	}

	go printOutputWithHeader("", stdoutColor, outReader2)
	go printOutputWithHeader("", stderrColor, errReader2)

	err = cmd.Wait()

	stdout = bufout.String()
	stderr = buferr.String()

	if err != nil {
		if err2, ok := err.(*exec.ExitError); ok {
			if s, ok := err2.Sys().(syscall.WaitStatus); ok {
				err = nil
				exitCode = s.ExitStatus()
			}
		}
	}
	return
}

func printOutputWithHeader(header string, color string, r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Printf("%s%s\n", header, scanner.Text())
	}
}
