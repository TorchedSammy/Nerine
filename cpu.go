package main

import (
	"fmt"
	"os"

	"github.com/TorchedSammy/Nerine/opcode"
)

type CPU struct {
	ram [0xffff]byte // 65535 bytes of ram
	registers [15]int
	ip int // Instruction pointer
	state int
}

func NewCPU() *CPU {
	c := &CPU{}
	return c
}

func (cpu *CPU) Load(data []byte) {
	if len(data) >= 0xffff {
		fmt.Println("Program larger than RAM")
		os.Exit(1)
	}

	for i := 0; i < len(data); i++ {
		cpu.ram[i] = data[i]
	}
}
func (cpu *CPU) Run() {
	cpu.state = 1 // Running
	for cpu.state == 1 {
		op := opcode.New(cpu.ram[cpu.ip])
		fmt.Printf("ip: %04X // %02X %s\n", cpu.ip, op.Val(), op.String())

		switch int(op.Val()) {
		case opcode.HALT:
			cpu.state = 0
		case opcode.ADD:
			cpu.ip++
			reg := cpu.ram[cpu.ip]
			cpu.ip++
			a := cpu.ram[cpu.ip]
			cpu.ip++
			b := cpu.ram[cpu.ip]

			av := cpu.registers[a]
			bv := cpu.registers[b]
			cpu.registers[reg] = av + bv
			fmt.Printf("%d reg -> %d\n", reg, av + bv)
		case opcode.MV:
			cpu.ip++
			reg := cpu.ram[cpu.ip]
			cpu.ip++
			val := cpu.ram[cpu.ip]

			cpu.registers[reg] = int(val)
			fmt.Printf("%d reg -> %d\n", reg, val)
		}
		cpu.ip++
	}
}
