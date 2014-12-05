package streamy

type filter func(int) bool

type stream struct {
	length  int
	filters []filter
}

func New(length int) *stream {
	return &stream{
		length: length,
	}
}

func (s *stream) Filter(f filter) *stream {
	s.filters = append(s.filters, f)
	return s
}

func (s *stream) GetFirst() (int, bool) {
	var (
		j int
	)
	for i := 0; i < s.length; i++ {
		for j = 0; j < len(s.filters); j++ {
			if !s.filters[j](i) {
				break
			}
		}
		if j == len(s.filters) {
			return i, true
		}
	}
	return 0, false
}

func (s *stream) GetAll() (res []int) {
	var (
		j int
	)
	for i := 0; i < s.length; i++ {
		for j = 0; j < len(s.filters); j++ {
			if !s.filters[j](i) {
				break
			}
		}
		if j == len(s.filters) {
			res = append(res, i)
		}
	}
	return res
}
