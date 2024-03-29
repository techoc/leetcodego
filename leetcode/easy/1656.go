package easy

type OrderedStream struct {
	stream []string
	ptr    int
}

func OrderedStreamConstructor(n int) OrderedStream {
	//构造
	return OrderedStream{make([]string, n+1), 1}
}

func (s *OrderedStream) Insert(idKey int, value string) []string {
	s.stream[idKey] = value
	start := s.ptr
	for s.ptr < len(s.stream) && s.stream[s.ptr] != "" {
		s.ptr++
	}
	return s.stream[start:s.ptr]
}
