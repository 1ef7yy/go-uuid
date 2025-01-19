package v1_test

import (
	"testing"

	v1 "github.com/1ef7yy/go-uuid/v1"
)

func TestV1Generation(t *testing.T) {
	uuid, err := v1.GenerateUUIDv1()

	if err != nil {
		t.Errorf("error generating uuidv1: %s", err.Error())
	}

	t.Logf("[success] uuidv1 generated: %s", uuid)

}
