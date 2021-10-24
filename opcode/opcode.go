package opcode

// Opcodes
var (
	HALT = 0x00
	MV = 0x01
	ADD = 0x02
	SUB = 0x03
	MUL = 0x04
	DIV = 0x05
)

type Op struct {
	instr byte
}

func New(instr byte) *Op {
	opcode := &Op{instr}
	return opcode
}

func (o *Op) Val() byte {
	return o.instr
}

func (o *Op) String() string {
	switch int(o.instr) {
		case HALT: return "HALT"
		case MV: return "MV"
		case ADD: return "ADD"
		case SUB: return "SUB"
		case MUL: return "MUL"
		case DIV: return "DIV"
	}

	return "UNK"
}
