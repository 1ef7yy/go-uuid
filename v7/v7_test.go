package v7

import (
	"testing"
)


func TestV7Generation(t *testing.T) {
	uuid, err := GenerateUUIDv7()

	if err != nil {
		t.Errorf("error generating uuidv7: %s", err.Error())
	}

	t.Logf("[success] uuidv7 generated: %s", uuid)
}
