package algorithm

type SingleQueue struct {
	queue []int
	size  int
}

func (s *SingleQueue) Push(index, value int) {
	if len(s.queue) > 0 && index-s.queue[0] >= s.size {
		s.queue = s.queue[1:]
	}
	for len(s.queue) > 0 && s.queue[len(s.queue)-1] > value {
		s.queue = s.queue[:len(s.queue)-1]
	}
	s.queue = append(s.queue, index)
}

func (s *SingleQueue) Pop() int {
	if len(s.queue) <= 0 {
		return -1
	}

	return s.queue[len(s.queue)-1]
}
