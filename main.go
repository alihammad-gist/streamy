package streamy

type Filter func(int) bool

var (
	// the order in which list of values will be evaluted
	OrderAsc  = true
	OrderDesc = false
)

type stream struct {
	length  int
	filters []Filter
}

func New(length int) *stream {
	return &stream{
		length: length,
	}
}

func (s *stream) Filter(f Filter) *stream {
	s.filters = append(s.filters, f)
	return s
}

func (s *stream) GetN(n int, order bool) (res []int) {
	if n < 0 {
		// what u want!
		return
	}
	fChain := func(i int) bool {
		var j int
		for j = 0; j < len(s.filters); j++ {
			if !s.filters[j](i) {
				break
			}
		}
		if j == len(s.filters) {
			return true
		}
		return false
	}

	if order == OrderAsc {
		for i := 0; i < s.length; i++ {
			if fChain(i) {
				res = append(res, i)
				if len(res) == n {
					return
				}
			}
		}
	} else {
		for i := s.length - 1; i >= 0; i-- {
			if fChain(i) {
				res = append(res, i)
				if len(res) == n {
					return
				}
			}
		}
	}
	return
}

func (s *stream) GetAll() (res []int) {
	return s.GetN(0, OrderAsc)
}

func (s *stream) GetFirst() (int, bool) {
	res := s.GetN(1, OrderAsc)
	if len(res) == 0 {
		return -1, false
	} else {
		return res[0], true
	}
}

func (s *stream) GetLast() (int, bool) {
	res := s.GetN(1, OrderDesc)
	if len(res) == 0 {
		return -1, false
	} else {
		return res[0], true
	}
}
