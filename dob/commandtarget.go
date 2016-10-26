package dob

import (
	"os/exec"
	"strings"
)

type CommandTarget struct {
	deps []Target

	stdin     string
	command   string
	arguments []string
}

func (t *CommandTarget) GetDeps() []Target {
	return t.deps
}
func (t *CommandTarget) AddDependencies(toAdd []Target) {
	t.deps = append(t.deps, toAdd...)
}

func (t *CommandTarget) Build() bool {
	cmd := exec.Command(t.command, t.arguments...)
	cmd.Stdin = strings.NewReader(t.stdin)
	// this blocks
	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

// If you want this to rebuild, add dependencies so it works
func (t *CommandTarget) NeedsRebuild() bool {
	return false
}

func NewCommandTarget(command string, arguments []string, stdin string) *CommandTarget {
	return &CommandTarget{
		stdin:     stdin,
		command:   command,
		arguments: arguments,
	}
}
