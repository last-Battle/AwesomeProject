package abc

type Queue []interface{}

func (q *Queue) Shift() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

type Stack []interface{}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	tail := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tail.(int)
}
