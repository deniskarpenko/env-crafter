package builders

import (
	"fmt"
	"os"
	"path/filepath"
)

type DockerfileBuilder struct {
}

const DirPerm = 0755

func (builder *DockerfileBuilder) Build(filePath string, content string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return fmt.Errorf("Can't create file: %w", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		return fmt.Errorf("Can't write to file: %w", err)
	}

	return nil
}

func GetAppRoot() (string, error) {
	if envPath := os.Getenv("APP_ROOT"); envPath != "" {
		return envPath, nil
	}

	execPath, err := os.Executable()

	if err != nil {
		return "", err
	}

	realExecPath, err := filepath.EvalSymlinks(execPath)

	if err != nil {
		return "", err
	}

	return filepath.Dir(realExecPath), nil
}
