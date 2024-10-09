package lib

import (
	"fmt"
	"os/exec"
)

func FindExecutablePath(executable string) (string, error) {
	path, err := exec.LookPath(executable)
	if err != nil {
		return "", fmt.Errorf("executable %s not found", executable)
	}
	return path, nil
}
