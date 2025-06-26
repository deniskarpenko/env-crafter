package docker

import (
	"testing"
)

func TestWebBuilder_Build(t *testing.T) {
	_, err := NewWebBuilder()

	if err != nil {
		t.Fatal(err)
	}
}
