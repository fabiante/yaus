package drivers

import (
	"bytes"
	"fmt"
	"os/exec"
)

type CLIDriver struct {
	sourceFile string
}

func NewCLIDriver(sourceFile string) *CLIDriver {
	return &CLIDriver{sourceFile: sourceFile}
}

func (d *CLIDriver) ShortenURL(input string) (string, error) {
	cmd := exec.Command("go", "run", d.sourceFile, input)

	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf("invoking CLI failed: %w\nstdout: %s", err, stdout.String())
	}

	return stdout.String(), nil
}
