package v4

import (
	"testing"
)


func TestV4Generation(t *testing.T) {
	uuid, err := GenerateUUIDv4()

	if err != nil {
		t.Errorf("error generating uuidv4: %s", err.Error())
	}

	t.Logf("[success] uuidv4 generated: %s", uuid)
}
