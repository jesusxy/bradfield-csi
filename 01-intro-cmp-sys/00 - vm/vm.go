package vm

import "log"

const (
	Load  = 0x01
	Store = 0x02
	Add   = 0x03
	Sub   = 0x04
	Halt  = 0xff // 255 int
)

// Stretch goals
const (
	Addi = 0x05
	Subi = 0x06
	Jump = 0x07
	Beqz = 0x08
)

// Given a 256 byte array of "memory", run the stored program
// to completion, modifying the data in place to reflect the result
//
// The memory format is:
//
// 00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f ... ff
// __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ __ ... __
// ^==DATA===============^ ^==INSTRUCTIONS==============^
func compute(memory []byte) {
	registers := [3]byte{8, 0, 0} // PC, R1 and R2

	// Keep looping, like a physical computer's clock
	for {
		pc := registers[0]
		op := memory[pc] // fetch the opcode

		// // decode and execute
		switch op {
		case Load:
			load(&registers, memory)
		case Store:
			store(&registers, memory)
		case Add:
			add(&registers)
		case Sub:
			sub(&registers)
		case Halt:
			return
		default:
			log.Fatalf("invalid opcode %x", op)
			return
		}

		update_pc(&registers)

	}
}

func add(registers *[3]byte) {
	registers[1] = registers[1] + registers[2]
}

func sub(registers *[3]byte) {
	registers[1] = registers[1] - registers[2]
}

func load(registers *[3]byte, memory []byte) {
	reg, addr := fetch_operands(registers, memory)

	registers[reg] = memory[addr]
}

func store(registers *[3]byte, memory []byte) {
	reg, addr := fetch_operands(registers, memory)

	if addr < 8 {
		memory[addr] = registers[reg]
	}
}

func update_pc(registers *[3]byte) {
	registers[0] += 3
}

func fetch_operands(reg *[3]byte, mem []byte) (byte, byte) {
	pc := reg[0]
	x := mem[pc+1]
	y := mem[pc+2]

	return x, y
}
