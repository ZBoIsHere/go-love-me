package queue

func main() {
	// queue top is len(queue)-1
	// queue bottom is 0
	q := []int{1, 2, 3}
	// enqueue
	q = append(q, 4)
	// dequeue
	q = q[1:]
}