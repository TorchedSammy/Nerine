package main

import (
	"fmt"
)

func main() {
	cpu := NewCPU()
	data := []byte{0x01, 0x01, 0x02, 0x01, 0x02, 0x02, 0x02, 0x03, 0x01, 0x02}
	cpu.Load(data)
	cpu.Run()
	fmt.Println("exited!")
}
