package goset

type Set map[string]struct{}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Add(element ...string) {
	for _, e := range element {
		s[e] = struct{}{}
	}
	return
}

func (s Set) Has(element string) {
	_, ok := s[element]
	return ok
}

func (s Set) Del(element string) {
	delete(s, element)
	return
}

func (s Set) ToList() {
	elements := make([]string, s.Len())
	i := 0
	for k := range s {
		elements[i] = k
		i++
	}
	return elements
}