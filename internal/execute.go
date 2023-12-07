package internal

import (
	"bytes"
	"os/exec"
	"strings"
)

type ExecuteResult struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exitCode"`
}

// Execute runs a command and returns the result
func Execute(command string) ExecuteResult {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		panic("No command specified")
	}
	cmd := exec.Command(parts[0], parts[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_ = cmd.Run()

	return ExecuteResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}
}
