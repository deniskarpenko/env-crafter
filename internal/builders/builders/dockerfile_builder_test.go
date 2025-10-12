package builders

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDockerfileBuilder_Build(t *testing.T) {
	db := DockerfileBuilder{}

	filePath := "php.Dockerfile"

	content := "!"

	err := db.Build(filePath, content)

	if err != nil {
		t.Error("Error:", err)
	}
}

func TestGetAppRoot_FromEnv(t *testing.T) {
	expectedPath := "tes/var"

	err := os.Setenv("APP_ROOT", expectedPath)

	if err != nil {
		t.Fatalf("Error  setting environment variable:%v", err)
	}

	defer os.Unsetenv("APP_ROOT")

	path, err := GetAppRoot()

	if err != nil {
		t.Fatalf("Error  getting app root:%v", err)
	}

	if path != expectedPath {
		t.Fatalf("Wrong app root:expected %v, got %v", expectedPath, path)
	}
}

func TestGetAppRoot(t *testing.T) {
	os.Unsetenv("APP_ROOT")

	path, err := GetAppRoot()

	if err != nil {
		t.Fatalf("Error  getting app root:%v", err)
	}

	realPath, err := os.Executable()

	if err != nil {
		t.Fatalf("Error  getting app root:%v", err)
	}

	realExecPath, err := filepath.EvalSymlinks(realPath)

	if err != nil {
		t.Fatalf("Error getting real app root :%v", err)
	}

	expected := filepath.Dir(realExecPath)

	if path != expected {
		t.Fatalf("Wrong app root:expected %v, got %v", expected, path)
	}
}
