package HelperFunctions

import (
	"bufio"
	"io"
)

// RunBrainfuck runs a Brainfuck program on 30000 8-bit cells. , reads a byte from Input (0 at EOF) and . writes one to Output.
func RunBrainfuck(Program string, Input io.Reader, Output io.Writer) {
	Code := []byte{}
	for _, Char := range []byte(Program) {
		switch Char {
		case '+', '-', '<', '>', '[', ']', '.', ',':
			Code = append(Code, Char)
		}
	}
	Jumps := make([]int, len(Code))
	Open := []int{}
	for Idx, Char := range Code {
		if Char == '[' {
			Open = append(Open, Idx)
		}
		if Char == ']' {
			Start := Open[len(Open)-1]
			Open = Open[:len(Open)-1]
			Jumps[Idx], Jumps[Start] = Start, Idx
		}
	}

	Reader := bufio.NewReader(Input)
	Tape := make([]byte, 30000)
	Ptr := 0
	for Pc := 0; Pc < len(Code); Pc++ {
		switch Code[Pc] {
		case '+':
			Tape[Ptr]++
		case '-':
			Tape[Ptr]--
		case '>':
			Ptr++
		case '<':
			Ptr--
		case '.':
			Output.Write(Tape[Ptr : Ptr+1])
		case ',':
			Tape[Ptr], _ = Reader.ReadByte()
		case '[':
			if Tape[Ptr] == 0 {
				Pc = Jumps[Pc]
			}
		case ']':
			if Tape[Ptr] != 0 {
				Pc = Jumps[Pc]
			}
		}
	}
}
