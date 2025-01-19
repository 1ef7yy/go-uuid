package main

import (
	"fmt"

	v1 "github.com/1ef7yy/go-uuid/v1"
)

func main() {
	uuidv1, err := v1.GenerateUUIDv1()

	if err != nil {
		fmt.Printf("error generating uuidv1: %s", err.Error())
		return
	}

	fmt.Printf("generated uuidv1: %s\n", uuidv1)

}
