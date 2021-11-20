package algorithms

type StringStack struct {
	list []string
}

func NewStringStack() *StringStack {
	return &StringStack{
		list: make([]string, 0),
	}
}

func (s *StringStack) Pop() string {
	val := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return val
}

func (s *StringStack) Push(val string) {
	s.list = append(s.list, val)
}

func (s *StringStack) Peek() string {
	return s.list[len(s.list)-1]
}

func (s *StringStack) Length() int {
	return len(s.list)
}

func (s *StringStack) IsEmpty() bool {
	return len(s.list) == 0
}

//  =========================
//	INT STACK
//  =========================

type IntStack struct {
	list []int
}

func NewIntStack() *IntStack {
	return &IntStack{
		list: make([]int, 0),
	}
}

func (s *IntStack) Pop() int {
	val := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	return val
}

func (s *IntStack) Push(val int) {
	s.list = append(s.list, val)
}

func (s *IntStack) Peek() int {
	return s.list[len(s.list)-1]
}

func (s *IntStack) Length() int {
	return len(s.list)
}

func (s *IntStack) IsEmpty() bool {
	return len(s.list) == 0
}
