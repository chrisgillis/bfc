package brainfuck

import "io"

// Machine represents a machine capable of executing brainfuck
// source code
type Machine struct {
	code []*Instruction // the parsed source code the machine will execute
	memory [30000]int // 30k byte sized zeroed memory cells
	readBuf []byte // for ',' instructions

	ip int // instruction pointer (points to source instruction)
	dp int // data pointer (points to memory cell)

	input io.Reader
	output io.Writer
}

func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code: instructions,
		input: in,
		output: out,
		readBuf: make([]byte, 1),
	}
}

func (m *Machine) Execute() {
	// iterate through every character in the source code
	// no brainfuck instruction is greater than 1 char
	for m.ip < len(m.code) {
		instruction := m.code[m.ip]

		switch instruction.Type {
		// increment the value of the cell pointed to by dp
		case Plus:
			m.memory[m.dp] += instruction.Argument
		// decrement the value of the cell pointed to by dp
		case Minus:
			m.memory[m.dp] -= instruction.Argument
		// move the data pointer right one cell
		case Right:
			m.dp += instruction.Argument
		// move the data pointer left one cell
		case Left:
			m.dp -= instruction.Argument
		// store chars from input
		case ReadChar:
			for i := 0; i < instruction.Argument; i++ {
				m.readChar()
			}
		// write stored chars to output
		case PutChar:
			for i := 0; i < instruction.Argument; i++ {
				m.putChar()
			}
		// handle the loop instructions
		case JmpIfZero:
			if m.memory[m.dp] == 0 {
				m.ip = instruction.Argument
				continue // avoid incr instr pointer
			}
		case JmpIfNotZero:
			if m.memory[m.dp] != 0 {
				m.ip = instruction.Argument
				continue // avoid incr instr pointer
			}
		}

		// after handling the instruction, incr the instr pointer
		m.ip++
	}
}

// read a char from input and store it at the memory cell pointed to
// by dp
func (m *Machine) readChar() {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("read more than 1 byte")
	}
	m.memory[m.dp] = int(m.readBuf[0])
}

// write the value at the memory cell pointed to by dp to output
func (m *Machine) putChar() {
	m.readBuf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrote more than 1 byte")
	}
}