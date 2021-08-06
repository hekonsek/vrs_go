package exe

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type exe struct {
	cmd    *exec.Cmd
	output []string
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

func (exe *exe) Run() (output []string, success bool, err error) {
	outputBytes, err := exe.cmd.CombinedOutput()
	outputLines := strings.Split(string(outputBytes), "\n")
	exe.output = outputLines
	if outputLines[len(outputLines)-1] == "" {
		outputLines = outputLines[0 : len(outputLines)-1]
	}
	if e, ok := err.(*exec.ExitError); ok {
		return outputLines, e.Success(), nil
	}
	return outputLines, true, err
}

func (exe *exe) NoSuccessReport() error {
	command := exe.cmd.Path
	for _, arg := range exe.cmd.Args {
		command += fmt.Sprintf(" %s", arg)
	}
	return errors.New(fmt.Sprintf("Cannot succesfully execute '%s' command:\n%s",
		command,
		strings.Join(exe.output, "\n")),
	)
}

func ParseCommand(command string) []string {
	return strings.Split(command, " ")
}
