package test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/fatih/structs"
)

func Test01(t *testing.T) {
	type Server struct {
		Name        string `json:"name,omitempty"`
		ID          int
		Enabled     bool
		users       []string // not exported
		http.Server          // embedded
	}

	server := &Server{
		Name:    "gopher",
		ID:      123456,
		Enabled: true,
	}
	// Convert a struct to a map[string]interface{}
	// => {"Name":"gopher", "ID":123456, "Enabled":true}
	m := structs.Map(server)
	fmt.Println(m)

	// Convert the values of a struct to a []interface{}
	// => ["gopher", 123456, true]
	v := structs.Values(server)
	fmt.Println(v)

	// Convert the names of a struct to a []string
	// (see "Names methods" for more info about fields)
	n := structs.Names(server)
	fmt.Println(n)

	// Convert the values of a struct to a []*Field
	// (see "Field methods" for more info about fields)
	f := structs.Fields(server)
	fmt.Println(f)

	// Return the struct name => "Server"
	n2 := structs.Name(server)
	fmt.Println(n2)

	// Check if any field of a struct is initialized or not.
	h := structs.HasZero(server)
	fmt.Println(h)

	// Check if all fields of a struct is initialized or not.
	z := structs.IsZero(server)
	fmt.Println(z)

	// Check if server is a struct or a pointer to struct
	i := structs.IsStruct(server)
	fmt.Println(i)
}

func Test02(t *testing.T) {
	type Foo struct {
		A int `tag1:"Tag1" tag2:"Second Tag"`
		B string
	}

	// Struct
	f := Foo{A: 10, B: "Salutations"}

	// Struct类型的指针
	fPtr := &f

	fmt.Println(reflect.TypeOf(fPtr))
	fmt.Println(reflect.ValueOf(fPtr).Elem())

	// // Map
	// m := map[string]int{"A": 1, "B": 2}
	//
	// // channel
	// ch := make(chan int)
	//
	// // slice
	// sl := []int{1, 32, 34}
	//
	// //string
	// str := "string var"
	//
	// // string 指针
	// strPtr := &str
}

func Test03(t *testing.T) {
	type user struct {
		Id   int64
		Name string
		sex  bool
	}

	u := user{
		Id:   100,
		Name: "Tom",
		sex:  false,
	}
	tt := reflect.TypeOf(u)
	for i := 0; i < tt.NumField(); i++ {
		f := tt.Field(i)
		fmt.Println("field:", i, f)
	}
}
