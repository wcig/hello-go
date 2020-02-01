package test

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"testing"
)

// 1.write gob file
func TestGob01(t *testing.T) {
	maps := map[string]interface{}{
		"id":   100,
		"name": "Tom",
	}

	gobFile, err := os.Create("user.gob")
	if err != nil {
		log.Fatal("create gob file err:", err)
	}
	defer gobFile.Close()

	encoder := gob.NewEncoder(gobFile)
	if err := encoder.Encode(maps); err != nil {
		log.Fatal("gob encode err:", err)
	}
	fmt.Println("create gob file success!!!")
}

// 2.read gob file
func TestGob02(t *testing.T) {
	gobFile, err := os.Open("user.gob")
	if err != nil {
		log.Fatal("open gob file err:", err)
	}
	defer gobFile.Close()

	maps := make(map[string]interface{})
	decoder := gob.NewDecoder(gobFile)
	if err := decoder.Decode(&maps); err != nil {
		log.Fatal("gob decode err:", err)
	}
	fmt.Println("maps:", maps)
}

// maps: map[id:100 name:Tom]
