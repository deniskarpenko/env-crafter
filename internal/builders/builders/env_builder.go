package builders

import (
	"fmt"
	"os"
)

type EnvBuilder struct{}

func (e *EnvBuilder) Build(envData []string, filename string) (string, error) {

	file, err := os.Create(filename)

	if err != nil {
		return "", fmt.Errorf("Can't create file: %w", err)
	}

	defer file.Close()

	for _, line := range envData {
		_, err = file.WriteString(line + "\n")

		if err != nil {
			return "", fmt.Errorf("Can't write to file: %w", err)
		}
	}

	return file.Name(), nil
}
