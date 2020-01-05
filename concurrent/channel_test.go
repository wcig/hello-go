package test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 1.通道: 无缓冲通道 (创建方式: c := make(chan type))
func TestChannel1(t *testing.T) {
	fmt.Println("start work...")

	c := make(chan bool)
	// <- c // 错误: fatal error: all goroutines are asleep - deadlock!
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
		c <- true
	}()
	<-c

	fmt.Println("end work...")
}

// 2.通道: 有缓冲通道 (创建方式: c := make(chan type, num))
func TestChannel2(t *testing.T) {
	fmt.Println("start work...")

	workerNum := 4
	jobNum := 10
	tasks := make(chan string, workerNum)
	var wg sync.WaitGroup
	wg.Add(jobNum)

	for i := 0; i < workerNum; i++ {
		go work(tasks, i, &wg)
	}

	for i := 0; i < jobNum; i++ {
		tasks <- fmt.Sprintf("Task:%d ", i)
	}

	close(tasks)
	wg.Wait()

	fmt.Println("end work...")
}

func work(tasks chan string, i int, wg *sync.WaitGroup) {
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("worker %d shutdown\n", i)
			return
		}

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("job %s over with worker %d\n", task, i)
		wg.Done()
	}
}

// 3.无缓冲通道：channel的接收通常应该在发送之前
var (
	over = make(chan bool)
	msg  string
)

func aGoroutine1() {
	msg = "hello..."
	over <- true
}

func TestChannel3(t *testing.T) {
	go aGoroutine1()
	<-over
	fmt.Println(msg) // hello...

	// <-over // fatal error: all goroutines are asleep - deadlock!
	// go aGoroutine1()
	// fmt.Println(msg)
}

// true
// hello...

// 4.无缓冲通道
func aGoroutine2() {
	msg = "hello..."
	close(over) // 关闭通道也可以实现同样功能，因为关闭通道后依然可以从通道接受值，只是此时接收到的是零值
}

func TestChannel4(t *testing.T) {
	go aGoroutine2()
	val, ok := <-over
	fmt.Println("val:", val, ", ok:", ok)
	fmt.Println(msg)
}

// val: false , ok: false
// hello...

// 5.无缓冲通道：交换channel发送和接收
func aGoroutine3() {
	msg = "hello..."
	<-over
}

func TestChannel5(t *testing.T) {
	go aGoroutine3()
	over <- true
	fmt.Println(msg) // hello...

	// over <- true // fatal error: all goroutines are asleep - deadlock!
	// go aGoroutine3()
	// fmt.Println(msg)
}

// 6.有缓冲通道
func TestChannel6(t *testing.T) {
	var limit = make(chan int, 3)

	workFunc := func() {
		var sum int
		for i := 0; i < 10000; i++ {
			sum += i
		}
		fmt.Println("sum:", sum)
	}

	var works []func()
	for i := 0; i < 5; i++ {
		works = append(works, workFunc)
	}

	for _, w := range works {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
	// select {} // fatal error: all goroutines are asleep - deadlock!
}
