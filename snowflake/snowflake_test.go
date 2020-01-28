package snowflake_test

import (
	"fmt"
	"hello-go/snowflake"
	"log"
	"testing"
)

func TestSnowflake(t *testing.T) {
	g, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal("create new node err:", err)
	}

	for i := 0; i < 100; i++ {
		id := g.GenerateId()
		fmt.Println(id)
		// fmt.Printf("%2d: %064b\n", i, id)
	}
}

func TestNewNode(t *testing.T) {
	_, err := snowflake.NewNode(0)
	if err != nil {
		t.Fatalf("error creating NewNode, %s", err)
	}

	_, err = snowflake.NewNode(5000)
	if err == nil {
		t.Fatalf("no error creating NewNode, %s", err)
	}
}

func TestGenerateDuplicateID(t *testing.T) {
	node, _ := snowflake.NewNode(1)
	var x, y int64
	for i := 0; i < 1000000; i++ {
		y = node.GenerateId()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

func TestRace(t *testing.T) {

	node, _ := snowflake.NewNode(1)

	go func() {
		for i := 0; i < 1000000000; i++ {

			snowflake.NewNode(1)
		}
	}()

	for i := 0; i < 4000; i++ {

		node.GenerateId()
	}

}
