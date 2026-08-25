package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial

	for _, n := range s {
		result = fn(result, n)
	}

	return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial

	for i := s.Length() - 1; i >= 0; i-- {
		result = fn(s[i], result)
	}

	return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}

	for _, n := range s {
		if fn(n) {
			result = result.Append(IntList{n})
		}
	}

	return result
}

func (s IntList) Length() int {
	length := 0

	for range s {
		length += 1
	}

	return length
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, n := range s {
		s[i] = fn(n)
	}

	return s
}

func (s IntList) Reverse() IntList {
	result := IntList{}

	for i := s.Length() - 1; i >= 0; i-- {
		result = result.Append(IntList{s[i]})
	}

	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, s.Length()+lst.Length())

	for i, n := range s {
		result[i] = n
	}

	for i, n := range lst {
		result[s.Length()+i] = n
	}

	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}

	return s
}
