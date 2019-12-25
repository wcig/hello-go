package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// go协程使用：主线程执行完成，但此时创建的goroutine没有执行完
func TestGoroutine1(t *testing.T) {
	fmt.Println("go routine test==>>")

	// 创建一个匿名函数的goroutine
	go func() {
		for i:=0; i<100; i++ {
			fmt.Println(i)
		}
	}()
}

// 通过睡眠3秒达到执行完goroutine任务再退出程序
func TestGoroutine2(t *testing.T) {
	fmt.Println("go routine test==>>")

	// 创建一个匿名函数的goroutine
	go func() {
		for i:=0; i<100; i++ {
			fmt.Println(i)
		}
	}()
	time.Sleep(time.Second*3)
}

// 通过sync.WaitGroup等待所有goroutine执行完成再退出程序
func TestWaitGroup1(t *testing.T) {
	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个匿名函数的goroutine
	go func() {
		defer wg.Done()

		for i:=0; i<100; i++ {
			fmt.Println("goroutine1:", i)
		}
	}()

	// 创建另一个匿名函数的goroutine
	go func() {
		defer wg.Done()

		for i:=0; i>-100; i-- {
			fmt.Println("goroutine2:", i)
		}
	}()

	fmt.Println("wait to finish==>")
	wg.Wait()

	fmt.Println("main execute over==>>")
}

// 分配多个逻辑处理器给调度器使用
func TestWaitGroup2(t *testing.T) {
	// runtime.GOMAXPROCS(1) // 分配一个
	// runtime.GOMAXPROCS(2) // 分配两个
	runtime.GOMAXPROCS(runtime.NumCPU()) // 分配当前所有可用cpu个数个

	// wg 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	// 创建一个匿名函数的goroutine
	go func() {
		defer wg.Done()

		for i:=0; i<100; i++ {
			fmt.Println("goroutine1:", i)
		}
	}()

	// 创建另一个匿名函数的goroutine
	go func() {
		defer wg.Done()

		for i:=0; i>-100; i-- {
			fmt.Println("goroutine2:", i)
		}
	}()

	fmt.Println("wait to finish==>")
	wg.Wait()

	fmt.Println("main execute over==>>")
}