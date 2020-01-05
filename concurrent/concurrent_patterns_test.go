package test

import (
	"fmt"
	"sync"
	"testing"
)

// 1.并发模式：互斥锁 (mu.Lock(), mu.Unlock() 不在同一个goroutine中，不满足顺序一致性内存模)
func TestConcurrentPatternsTest1(t *testing.T) {
	var mu sync.Mutex
	go func() {
		fmt.Println("hello...")
		mu.Lock()
	}()
	mu.Unlock() // fatal error: sync: unlock of unlocked mutex
}

// 上例解决办法
func TestConcurrentPatternsTest2(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello...") // hello...
		mu.Unlock()
	}()
	mu.Lock()
}

// 2.并发模式：无缓冲通道 (根据Go语言内存模型规范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前)
func TestConcurrentPatternsTest3(t *testing.T) {
	done := make(chan int)
	go func() {
		fmt.Println("hello...") // hello...
		<-done
	}()
	done <- 1
}

// 上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证main退出之前后台线程能正常打印了。
// 更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步事件受管道缓存大小的影响。
func TestConcurrentPatternsTest4(t *testing.T) {
	done := make(chan int)
	go func() {
		fmt.Println("hello...") // hello...
		done <- 1
	}()
	<-done
}
