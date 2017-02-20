package brainfuck

type InsType byte

const (
	Plus InsType = '+'
	Minus InsType = '-'
	Right InsType = '>'
	Left InsType = '<'
	PutChar InsType = '.'
	ReadChar InsType = ','
	JmpIfZero InsType = '['
	JmpIfNotZero InsType = ']'
)

type Instruction struct {
	Type InsType

	// having an argument will allow us to make the instr set
	// more dense than the brainfuck spec and enable compiler
	// optimization
	Argument int
}