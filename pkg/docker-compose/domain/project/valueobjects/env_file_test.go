package valueobjects

import (
	"testing"
)

func TestNewEnvFile(t *testing.T) {
	fileName := ".env.test"
	fileData := []byte("DATABASE_URL=localhost\nAPI_KEY=secret123")

	envFile := NewEnvFile(fileName, fileData)

	if envFile.FileName() != fileName {
		t.Errorf("expected fileName %s, got %s", fileName, envFile.FileName())
	}

	if string(envFile.File()) != string(fileData) {
		t.Errorf("expected file data %s, got %s", string(fileData), string(envFile.File()))
	}
}

func TestEnvFile_FileName(t *testing.T) {
	expected := ".env.production"
	fileData := []byte("KEY=value")

	envFile := NewEnvFile(expected, fileData)

	if envFile.FileName() != expected {
		t.Errorf("expected %s, got %s", expected, envFile.FileName())
	}
}

func TestEnvFile_File(t *testing.T) {
	fileName := ".env"
	expected := []byte("PORT=8080\nHOST=0.0.0.0")

	envFile := NewEnvFile(fileName, expected)

	result := envFile.File()
	if string(result) != string(expected) {
		t.Errorf("expected %s, got %s", string(expected), string(result))
	}
}

func TestEnvFile_ToYaml(t *testing.T) {
	fileName := ".env"
	fileData := []byte("database:\n  host: localhost\n  port: 5432")

	envFile := NewEnvFile(fileName, fileData)

	result := envFile.ToYaml()
	expected := string(fileData)

	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestEnvFile_EmptyFile(t *testing.T) {
	fileName := ".env.empty"
	fileData := []byte("")

	envFile := NewEnvFile(fileName, fileData)

	if len(envFile.File()) != 0 {
		t.Errorf("expected empty file, got %d bytes", len(envFile.File()))
	}

	if envFile.ToYaml() != "" {
		t.Errorf("expected empty string, got %s", envFile.ToYaml())
	}
}
