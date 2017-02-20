package brainfuck

// the compiler will turn brainfuck source code into a slice
// of brainfuck machine instructions
type Compiler struct {
	code string
	codeLength int
	position int

	instructions []*Instruction
}

func NewCompiler(code string) *Compiler {
	return &Compiler {
		code: code,
		codeLength: len(code),
		instructions: []*Instruction{},
	}
}

func (c *Compiler) Compile() []*Instruction {
	// for tracking loop instruction stack
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.EmitWithArg(JmpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// get the last jmpifzero off of stack
			openInstr := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]

			// add the jmpifnotzero instruction with its correct position
			closeInstructionPos := c.EmitWithArg(JmpIfNotZero, openInstr)

			// update the matching jmpifzero instruction with the new pos
			c.instructions[openInstr].Argument = closeInstructionPos
		case '+':
			c.CompileFoldableInstruction('+', Plus)
		case '-':
			c.CompileFoldableInstruction('-', Minus)
		case '<':
			c.CompileFoldableInstruction('<', Left)
		case '>':
			c.CompileFoldableInstruction('>', Right)
		case '.':
			c.CompileFoldableInstruction('.', PutChar)
		case ',':
			c.CompileFoldableInstruction(',', ReadChar)
		}

		c.position++
	}

	return c.instructions
}

// perform some simple lookahead logic to group like instructions into one
func (c *Compiler) CompileFoldableInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.EmitWithArg(insType, count)
}

func (c *Compiler) EmitWithArg(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}