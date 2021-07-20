package stack

func main() {
	// stack top is len(s)-1
	// stack bottom is 0
	var s = []int{}
	// push
	s = append(s, 0)
	// pop
	s = s[:len(s)-1]
}
