package vm

import (
	"fmt"
)

type Stack struct {
	data []int
}

type Instruction struct {
	Opcode OP
	Data   int
}

type Program []Instruction

type OP int

const (
	PUSH OP = iota
	POP
	DUMP
	ADD
	SUB
	MUL
	DIV
	MOD
	EQ
	GT
	LT
	AND
	OR
	NOT
)

func NewStack() *Stack {
	return &Stack{}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *Stack) Run(program Program) {
	for _, instruction := range program {
		switch instruction.Opcode {
		case PUSH:
			s.Push(instruction.Data)
		case POP:
			_, ok := s.Pop()
			must(ok)
		case DUMP:
			s.Dump()
		case ADD:
			_, ok := s.Add()
			must(ok)
		case SUB:
			_, ok := s.Sub()
			must(ok)
		case MUL:
			_, ok := s.Mul()
			must(ok)
		case DIV:
			_, ok := s.Div()
			must(ok)
		case MOD:
			_, ok := s.Mod()
			must(ok)
    case EQ:
      _, ok := s.EQ()
      must(ok)
    case GT:
      _, ok := s.GT()
      must(ok)
    case LT:
      _, ok := s.LT()
      must(ok)
    case AND:
      _, ok := s.AND()
      must(ok)
    case OR:
      _, ok := s.OR()
      must(ok)
    case NOT:
      _, ok := s.NOT()
      must(ok)
		}
	}
}

func (s *Stack) Push(data int) (int, error) {
	s.data = append(s.data, data)
	return data, nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("Stack is Empty")
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return data, nil
}

func (s *Stack) Dump() {
	for _, data := range s.data {
		fmt.Println(data)
	}
}

func (s *Stack) Add() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}
	left, _ := s.Pop()
	right, _ := s.Pop()
	s.Push(left + right)
	return left + right, nil
}

func (s *Stack) Sub() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}
	left, _ := s.Pop()
	right, _ := s.Pop()
	s.Push(left - right)
	return left - right, nil
}

func (s *Stack) Mul() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}
	left, _ := s.Pop()
	right, _ := s.Pop()
	s.Push(left * right)
	return left * right, nil
}

func (s *Stack) Div() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if right == 0 {
		return 0, fmt.Errorf("Division by zero")
	}

	s.Push(left / right)
	return left / right, nil
}

func (s *Stack) Mod() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if right == 0 {
		return 0, fmt.Errorf("Division by zero")
	}

	s.Push(left % right)

	return left % right, nil

}

func (s *Stack) EQ() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if left == right {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}

func (s *Stack) GT() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if left > right {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}

func (s *Stack) LT() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if left < right {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}

func (s *Stack) AND() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if left != 0 && right != 0 {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}

func (s *Stack) OR() (int, error) {
	if len(s.data) < 2 {
		return 0, fmt.Errorf("Stack is empty")
	}

	left, _ := s.Pop()
	right, _ := s.Pop()

	if left != 0 || right != 0 {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}

func (s *Stack) NOT() (int, error) {
	if len(s.data) < 1 {
		return 0, fmt.Errorf("Stack is empty")
	}

	data, _ := s.Pop()
	if data == 0 {
		s.Push(1)
		return 1, nil
	} else {
		s.Push(0)
		return 0, nil
	}
}
