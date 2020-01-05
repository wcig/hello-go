package test

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 1.没有使用锁，会导致最后输出结果错误
func TestLock1(t *testing.T) {
	fmt.Println("no use lock==>>")
	runtime.GOMAXPROCS(4)

	num := 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 500; i++ {
			num++
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 500; i++ {
			num--
			time.Sleep(time.Microsecond)
		}
	}()

	wg.Wait()
	fmt.Println("num:", num)
}

// 2.使用原子函数 (atomic其他方法：atomic.StoreInt64(), atomic.LoadInt64())
func TestLock2(t *testing.T) {
	fmt.Println("use sync/atomic==>>")
	runtime.GOMAXPROCS(4)

	var num int64 = 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 500; i++ {
			atomic.AddInt64(&num, 1)
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 500; i++ {
			atomic.AddInt64(&num, -1)
			time.Sleep(time.Microsecond)
		}
	}()

	wg.Wait()
	fmt.Println("num:", num)
}

// 3.互斥锁 (开销比原子操作大)
var total struct {
	sync.Mutex
	val int64
}

func TestLock3(t *testing.T) {
	fmt.Println("use sync.Mutex==>>")
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		total.Lock()
		for i := 0; i < 500; i++ {
			total.val++
			time.Sleep(time.Microsecond)
		}
		total.Unlock()
	}()

	go func() {
		defer wg.Done()

		total.Lock()
		for i := 0; i < 500; i++ {
			total.val--
			time.Sleep(time.Microsecond)
		}
		total.Unlock()
	}()

	wg.Wait()
	fmt.Println("val:", total.val)
}

// 4.通过互斥锁+原子操作实现线程安全的单例模式
type singleton struct {
	val int
}

var (
	instance    *singleton
	initialized uint32
	mutext      sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mutext.Lock()
	defer mutext.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

func TestSingleton(t *testing.T) {
	s1 := Instance()
	fmt.Printf("s1 val: %d, addr:%p\n", s1.val, s1)
	s1.val++

	go func() {
		s2 := Instance()
		fmt.Printf("s2 val: %d, addr:%p\n", s2.val, s2)
		s2.val++
	}()

	time.Sleep(time.Millisecond)
	s3 := Instance()
	fmt.Printf("s3 val: %d, addr:%p\n", s3.val, s3)
}

// output:
// s1 val: 0, addr:0xc00008a1c8
// s2 val: 1, addr:0xc00008a1c8
// s3 val: 2, addr:0xc00008a1c8

// 5.sync.Once的实现
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if atomic.LoadUint32(&o.done) == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

// 6.通过sync.Once实现单例模式
var (
	instance2 *singleton
	once2     sync.Once
)

func Instance2() *singleton {
	once2.Do(func() {
		instance2 = &singleton{}
	})
	return instance2
}

func TestSingleton2(t *testing.T) {
	s1 := Instance2()
	fmt.Printf("s1 val: %d, addr:%p\n", s1.val, s1)
	s1.val++

	go func() {
		s2 := Instance2()
		fmt.Printf("s2 val: %d, addr:%p\n", s2.val, s2)
		s2.val++
	}()

	time.Sleep(time.Millisecond)
	s3 := Instance2()
	fmt.Printf("s3 val: %d, addr:%p\n", s3.val, s3)
}

// 7.sync.Once作用：传入的函数只执行一次，与init()函数区别：init()在go文件被加载时默认执行一次，sync.Once.Don()在需要时候执行一次
func TestSyncOnce(t *testing.T) {
	onceFunc := func() {
		fmt.Println("hello...")
	}

	var once sync.Once
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(onceFunc)
		}()
	}
	wg.Wait()
}
