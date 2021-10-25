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
		fmt.Println("Program larger than RAM! Exiting.")
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
		fmt.Printf("ip: %04X // running instr: %02X %s\n", cpu.ip, op.Val(), op.String())

		switch int(op.Val()) {
		case opcode.HALT:
			cpu.state = 0
		case opcode.ADD:
			reg, a, b := cpu.get3Val()
	
			av := cpu.registers[a]
			bv := cpu.registers[b]
			cpu.registers[reg] = av + bv
			cpu.debugPrintReg(reg)
		case opcode.SUB:
			reg, a, b := cpu.get3Val()
			av, bv := cpu.registers[a], cpu.registers[b]

			cpu.registers[reg] = av + bv
			cpu.debugPrintReg(reg)
		case opcode.MUL:
			reg, a, b := cpu.get3Val()
			av, bv := cpu.registers[a], cpu.registers[b]

			cpu.registers[reg] = av * bv
			cpu.debugPrintReg(reg)
		case opcode.DIV:
			reg, a, b := cpu.get3Val()
			av, bv := cpu.registers[a], cpu.registers[b]

			cpu.registers[reg] = av / bv
			cpu.debugPrintReg(reg)
		case opcode.MV:
			reg := cpu.atIP()
			val := cpu.atIP()

			cpu.registers[reg] = int(val)
			cpu.debugPrintReg(reg)
		default:
			fmt.Printf("Encountered unknown opcode: %02X\n", op.Val())
			os.Exit(1)
		}
		cpu.ip++
	}
}

func (cpu *CPU) get3Val() (byte, byte, byte) {
	a := cpu.atIP()
	b := cpu.atIP()
	c := cpu.atIP()

	return a, b, c
}

func (cpu *CPU) atIP() byte {
	cpu.ip++
	return cpu.ram[cpu.ip]
}

func (cpu *CPU) debugPrintReg(register byte) {
	fmt.Printf("@%d register -> %d\n", register, cpu.registers[register])
}

