package test

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 没有使用锁，会导致最后输出结果错误
func TestLock1(t *testing.T) {
	fmt.Println("use go lock==>>")
	runtime.GOMAXPROCS(4)

	num := 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i:=0; i<500; i++ {
			num++
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i:=0; i<500; i++ {
			num--
			time.Sleep(time.Microsecond)
		}
	}()

	wg.Wait()
	fmt.Println("num:", num)
}

// 使用原子函数
func TestLock2(t *testing.T) {
	fmt.Println("use go lock==>>")
	runtime.GOMAXPROCS(4)

	var num int64 = 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i:=0; i<500; i++ {
			atomic.AddInt64(&num, 1)
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i:=0; i<500; i++ {
			atomic.AddInt64(&num, -1)
			time.Sleep(time.Microsecond)
		}
	}()

	wg.Wait()
	fmt.Println("num:", num)
}

// atomic其他方法：atomic.StoreInt64(), atomic.LoadInt64() TODO
