package test

import (
	"fmt"
	"sort"
	"testing"
)

type intSlice []int

func (i intSlice) Len() int {
	return len(i)
}

func (i intSlice) Less(x, y int) bool {
	return i[x] <= i[y]
}

func (i intSlice) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

// 1.自己实现sort包的Interface接口
func TestSort1(t *testing.T) {
	is := intSlice{3, 1, 2, 5, 4}
	sort.Sort(is)
	fmt.Println("is:", is) // is: [1 2 3 4 5]
}

// 2.通过sort包的内置类型
func TestSort2(t *testing.T) {
	is := sort.IntSlice{3, 1, 2, 5, 4}
	sort.Sort(is)
	fmt.Println("is:", is) // is: [1 2 3 4 5]

	ss := sort.StringSlice{"Tom", "Jerry", "Bob"}
	sort.Sort(ss)
	fmt.Println("ss:", ss) // ss: [Bob Jerry Tom]

	fs := sort.Float64Slice{1.0, 100.23, 0.13}
	sort.Sort(fs)
	fmt.Println("fs:", fs) // fs: [0.13 1 100.23]

	s := []int{3, 1, 2, 5, 4}
	is2 := sort.IntSlice(s)
	sort.Sort(is2)
	fmt.Println("is2:", is2) // is2: [1 2 3 4 5]
}

type user struct {
	id   int
	name string
}

type users []user

func (u users) Len() int {
	return len(u)
}

func (u users) Less(x, y int) bool {
	if u[x].id != u[y].id {
		return u[x].id <= u[y].id
	}
	return u[x].name <= u[y].name
}

func (u users) Swap(x, y int) {
	u[x], u[y] = u[y], u[x]
}

// 3.对结构体数据进行排序
func TestSort3(t *testing.T) {
	us := users{
		{1, "Tom"},
		{3, "Jerry"},
		{2, "Tom"},
	}
	sort.Sort(us)
	fmt.Println("us:", us) // us: [{1 Tom} {2 Tom} {3 Jerry}]
}

// 4.通过sort.Slice()方法对切片进行排序
func TestSort4(t *testing.T) {
	u := []user{
		{1, "Tom"},
		{3, "Jerry"},
		{2, "Tom"},
	}
	sort.Slice(u, func(x, y int) bool {
		if u[x].id != u[y].id {
			return u[x].id <= u[y].id
		}
		return u[x].name <= u[y].name
	})
	fmt.Println("u:", u) // u: [{1 Tom} {2 Tom} {3 Jerry}]
}
