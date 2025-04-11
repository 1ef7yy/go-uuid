package v1

import (
	"testing"

)

func TestV1Generation(t *testing.T) {
	uuid, err := GenerateUUIDv1()

	if err != nil {
		t.Errorf("error generating uuidv1: %s", err.Error())
	}

	t.Logf("[success] uuidv1 generated: %s", uuid)

}
