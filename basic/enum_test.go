package test

import (
	"fmt"
	"testing"
)

type Weapon int

const (
	Arrow Weapon = iota // 开始生成枚举值, 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)

func TestEnum1(t *testing.T) {
	// 输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)
}

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

func TestEnum2(t *testing.T) {
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}

type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func TestEnum3(t *testing.T) {
	fmt.Printf("%d:%s\n", CPU, CPU) // 1:CPU
}
