package intcode

func Intcode(intcode []int) []int {
	length := len(intcode)
	for i := 0; i < length; i += 4 {
		if intcode[i] == 99 {
			break
		}

		opcode := intcode[i]
		operands := intcode[i+1 : i+3]
		target := intcode[i+3]

		switch opcode {
		case 1:
			intcode[target] = intcode[operands[0]] + intcode[operands[1]]
		case 2:
			intcode[target] = intcode[operands[0]] * intcode[operands[1]]
		default:
			panic("The opcode is unkown")
		}
	}
	return intcode
}
