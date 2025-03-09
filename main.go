package main

import (
	"fmt"

	v1 "github.com/1ef7yy/go-uuid/v1"
	v4 "github.com/1ef7yy/go-uuid/v4"
	v5 "github.com/1ef7yy/go-uuid/v5"
)

func main() {
	uuidv1, err := v1.GenerateUUIDv1()

	if err != nil {
		fmt.Printf("error generating uuidv1: %s", err.Error())
	}

	fmt.Printf("generated uuidv1: %s\n", uuidv1.String())

	uuidv4, err := v4.GenerateUUIDv4()

	if err != nil {
		fmt.Printf("error generating uuidv4: %s\n", err.Error())
	}

	fmt.Printf("uuidv4 generated: %s\n", uuidv4.String())

	uuidv5, err := v5.GenerateUUIDv5(v5.DNSNamespace, []byte("example.com"))

	if err != nil {
		fmt.Printf("error generating uuidv5: %s\n", err.Error())
	}

	fmt.Printf("uuidv5 generated: %s\n", uuidv5.String())
}
