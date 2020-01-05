package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 1.顺序一致性只在一个goroutine有效
var (
	a    string
	done bool
)

func setup() {
	time.Sleep(time.Millisecond)
	a = "hello..."
	done = true
}

func TestGoroutineSeq1(t *testing.T) {
	go setup()
	for !done {
	}
	fmt.Println(a)
}

// 2.通过同步语给两个goroutine明确排序
func TestGoroutineSeq2(t *testing.T) {
	done := make(chan bool)
	go func() {
		fmt.Println("hello...")
		done <- true
	}()
	<-done
}

// 3.通过互斥锁给两个goroutine明确排序
func TestGoroutineSeq3(t *testing.T) {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		fmt.Println("hello...")
		mu.Unlock()
	}()
	mu.Lock()
}
