package main

import (
	"fmt"
)

type Disassembleable interface {
	disassemble()
}

type rInstruction struct {
	bits        uint32
	address     uint32
	instruction string
	opcode      uint32
	src1        uint32
	src2        uint32
	dest        uint32
	function    uint32
}

type iInstruction struct {
	bits        uint32
	address     uint32
	instruction string
	opcode      uint32
	src1        uint32
	src2        uint32
	constant    int16
}

type Mask struct {
	bits  uint32
	shift uint8
}

var opcodeMask = Mask{0xFC000000, 26}
var src1Mask = Mask{0x03D00000, 21}
var src2Mask = Mask{0x001F0000, 16}
var rDestMask = Mask{0x0000F800, 11}
var rFuncMask = Mask{0x0000003F, 0}
var iConstMask = Mask{0x0000FFFF, 0}

var pc = uint32(0x7A060)
var addressSize = uint32(4)

var input = []uint32{
	0x022DA822,
	0x8EF30018,
	0x12A70004,
	0x02689820,
	0xAD930018,
	0x02697824,
	0xAD8FFFF4,
	0x018C6020,
	0x02A4A825,
	0x158FFFF6,
	0x8E59FFF0,
}

func main() {
	// build slice of instructions from input
	instructions := buildInstructions(input)

	// decode then print each instruction
	for _, instruct := range instructions {
		instruct.disassemble()
		fmt.Println(instruct)
	}
}

func buildInstructions(input []uint32) []Disassembleable {
	instructions := make([]Disassembleable, len(input))

	// need to get opcode, if 000000, build an rInstruction, else iInstruction
	for i, inputBits := range input {
		var newInstruction Disassembleable
		opcode := maskAndShift(opcodeMask, inputBits)

		if opcode == 0 {
			// decode as R
			newInstruction = &rInstruction{bits: inputBits, opcode: opcode,
				address: pc}
		} else {
			// decode as I
			newInstruction = &iInstruction{bits: inputBits, opcode: opcode,
				address: pc}
		}
		pc += 4
		instructions[i] = newInstruction
	}

	return instructions
}

func (r *rInstruction) disassemble() {
	r.src1 = maskAndShift(src1Mask, r.bits)
	r.src2 = maskAndShift(src2Mask, r.bits)
	r.dest = maskAndShift(rDestMask, r.bits)
	r.function = maskAndShift(rFuncMask, r.bits)
	r.instruction = r.getInstruction()
}

func (i *iInstruction) disassemble() {
	i.src1 = maskAndShift(src1Mask, i.bits)
	i.src2 = maskAndShift(src2Mask, i.bits)
	i.constant = maskAndShiftShort(iConstMask, int16(i.bits))
	i.instruction = i.getInstruction()
}

func (r *rInstruction) getInstruction() string {
	switch r.function {
	case 0x20:
		return fmt.Sprintf("add")
	case 0x22:
		return fmt.Sprintf("sub")
	case 0x24:
		return fmt.Sprintf("and")
	case 0x25:
		return fmt.Sprintf("or")
	default:
		return ""
	}
}

func (i *iInstruction) getInstruction() string {

	switch i.opcode {
	case 0x4:
		return fmt.Sprintf("beq")
	case 0x5:
		return fmt.Sprintf("bne")
	case 0x23:
		return fmt.Sprintf("lw")
	case 0x2B:
		return fmt.Sprintf("sw")
	default:
		return ""
	}
}

func (r *rInstruction) String() string {
	//todo check order of source 1 & 2
	return fmt.Sprintf("%X %s $%d, $%d, $%d", r.address, r.instruction,
		r.dest, r.src1, r.src2)
}

func (i *iInstruction) String() string {
	// beq expressed as: 7A07C beq $7, $8, address 7A0F0
	// (just have to add offset to address and print)
	if i.instruction == "bne" || i.instruction == "beq" {

		// shift offset/const by 2 left to decompress, account for PREINCREMENTED pc
		branchAddress := i.address + addressSize + (uint32(i.constant) << 2)
		return fmt.Sprintf("%X %s $%d, $%d address %X", i.address, i.instruction,
			i.src1, i.src2, branchAddress)
	} else {
		return fmt.Sprintf("%X %s $%d, %d ($%d)", i.address, i.instruction,
			i.src1, i.constant, i.src2)
	}
}

func maskAndShift(mask Mask, inputBits uint32) uint32 {
	return (inputBits & mask.bits) >> mask.shift
}

func maskAndShiftShort(mask Mask, inputBits int16) int16 {
	return (inputBits & int16(mask.bits)) >> mask.shift
}
