package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	type user struct{}

	u1 := user{}
	u2 := &user{}
	var i int = 10
	var f float32 = 1.0
	var b bool = true

	slice := []interface{}{
		u1, u2, i, f, b,
	}
	for _, val := range slice {
		printTypeAndKind(val)
	}
}

func printTypeAndKind(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Printf("name: %v, kind: %v\n", t.Name(), t.Kind())
}
