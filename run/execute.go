package run

import (
	"os/exec"

	"github.com/haquenafeem/executor/internal/consts"
)

func Execute(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	bytes, err := cmd.Output()
	if err != nil {
		return consts.EmptyString, err
	}

	return string(bytes), nil
}
