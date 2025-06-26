package builders

import (
	"os"
	"testing"
)

func TestEnvBuilder_Build(t *testing.T) {
	eb := EnvBuilder{}

	fileName := "mysql.env"

	content := []string{
		"MYSQL_ROOT_PASSWORD: secret",
		"MYSQL_DATABASE: attendance",
		"MYSQL_USER: attendance_admin",
		"MYSQL_PASSWORD: secret",
	}

	fName, err := eb.Build(content, fileName)

	if err != nil {
		t.Error("Error:", err)
	}

	if _, err = os.Stat(fName); os.IsNotExist(err) {
		t.Error("Error:", err)
	}
}
