package main

import "fmt"

type Register struct {
	A byte
	B byte
	C byte
	D byte
}

func NewRegister() *Register {
	return &Register{
		A: 0x00,
		B: 0x00,
		C: 0x00,
		D: 0x00,
	}
}

func main() {
	// Register型のデータが格納されたメモリ番地
	rgst := NewRegister()
	// 変数rgstのメモリ番地
	addr_rgst := &rgst
	fmt.Println(&rgst)
	fmt.Println(*addr_rgst)

	a := 1
	fmt.Println(&a) // (0xc000016110)
	fmt.Println(a)  // 1
	// fmt.Println(*a) // エラー
}
