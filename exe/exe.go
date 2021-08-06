package exe

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type exe struct {
	cmd *exec.Cmd
}

type runResult struct {
	exe *exe

	output  []string
	success bool
	err     error
}

func New(command string) *exe {
	commandSegments := ParseCommand(command)
	cmd := exec.Command(commandSegments[0], commandSegments[1:]...)
	return &exe{cmd: cmd}
}

func (exe *exe) InDirectory(dir string) *exe {
	exe.cmd.Dir = dir
	return exe
}

func (exe *exe) Run() *runResult {
	outputBytes, err := exe.cmd.CombinedOutput()
	outputLines := strings.Split(string(outputBytes), "\n")
	if outputLines[len(outputLines)-1] == "" {
		outputLines = outputLines[0 : len(outputLines)-1]
	}
	if e, ok := err.(*exec.ExitError); ok {
		return &runResult{exe: exe, output: outputLines, success: e.Success(), err: nil}
	}
	return &runResult{exe: exe, output: outputLines, success: true, err: err}
}

func (result *runResult) NoSuccessReport() error {
	command := result.exe.cmd.Path
	for _, arg := range result.exe.cmd.Args {
		command += fmt.Sprintf(" %s", arg)
	}
	return errors.New(fmt.Sprintf("Cannot succesfully execute '%s' command:\n%s",
		command,
		strings.Join(result.output, "\n")),
	)
}

func (result *runResult) Output() []string {
	return result.output
}

func (result *runResult) Success() bool {
	return result.success
}

func (result *runResult) Err() error {
	return result.err
}

func ParseCommand(command string) []string {
	return strings.Split(command, " ")
}
