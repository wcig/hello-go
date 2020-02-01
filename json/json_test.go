package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name, omitempty"`
}

// 1.struct -> json str
func TestJson01(t *testing.T) {
	u1 := &user{
		Id:   100,
		Name: "Tom",
	}
	u2 := &user{
		Id:   101,
		Name: "",
	}

	bytes1, err := json.Marshal(u1)
	if err == nil {
		jsonStr1 := string(bytes1)
		fmt.Println("json str1:", jsonStr1)
	}

	bytes2, err := json.Marshal(u2)
	if err == nil {
		jsonStr2 := string(bytes2)
		fmt.Println("json str2:", jsonStr2)
	}
}

// json str1: {"id":100,"name":"Tom"}
// json str2: {"id":101,"name":""}

// 2.json str -> struct
func TestJson02(t *testing.T) {
	jsonStr1 := `{"id":100,"name":"Tom"}`
	jsonStr2 := `{"id":101,"name":""}`

	var (
		u1 user
		u2 user
	)

	if err := json.Unmarshal([]byte(jsonStr1), &u1); err == nil {
		fmt.Println("u1:", u1)
	}
	if err := json.Unmarshal([]byte(jsonStr2), &u2); err == nil {
		fmt.Println("u2:", u2)
	}
}

// u1: {100 Tom}
// u2: {101 }

// 3.map -> json str
func TestJson03(t *testing.T) {
	maps := map[string]interface{}{
		"id":   100,
		"name": "Tom",
	}

	if bytes, err := json.Marshal(maps); err == nil {
		fmt.Println("json str:", string(bytes))
	}
}

// json str: {"id":100,"name":"Tom"}

// 4.json str -> map
func TestJson04(t *testing.T) {
	jsonStr := `{"id":100,"name":"Tom"}`
	maps := make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsonStr), &maps); err == nil {
		fmt.Println("maps:", maps)
	}
}

// maps: map[id:100 name:Tom]

// 5.write json file
func TestJson05(t *testing.T) {
	u := &user{
		Id:   100,
		Name: "Tom",
	}

	jsonFile, err := os.Create("user.json")
	if err != nil {
		log.Fatal("create file err:", err)
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	if err := encoder.Encode(u); err != nil {
		log.Fatal("json encoder err:", err)
	}

	fmt.Println("create file success!!!")
}

// 6.read json file
func TestJson6(t *testing.T) {
	jsonFile, err := os.Open("user.json")
	if err != nil {
		log.Fatal("open file err:", err)
	}
	defer jsonFile.Close()

	var u user
	decoder := json.NewDecoder(jsonFile)
	if err := decoder.Decode(&u); err != nil {
		log.Fatal("json decoder err:", err)
	}

	fmt.Println("read json file success!!!")
	fmt.Println("user:", u)
}

// 7.map -> format json str
func TestJson07(t *testing.T) {
	maps := map[string]interface{}{
		"id":   100,
		"name": "Tom",
	}

	if bytes, err := json.MarshalIndent(maps, "", "    "); err == nil {
		fmt.Println(string(bytes))
	}
}

// {
//    "id": 100,
//    "name": "Tom"
// }

// 8.json str format
func TestJson08(t *testing.T) {
	inputJsonStr := `{"id":100,"name":"Tom"}`

	var jsonBuffer bytes.Buffer
	err := json.Indent(&jsonBuffer, []byte(inputJsonStr), "", "    ")
	if err != nil {
		log.Fatal("json str format err:", err)
	}
	resultJsonStr := jsonBuffer.String()
	fmt.Println(resultJsonStr)
}

// {
//    "id": 100,
//    "name": "Tom"
// }
