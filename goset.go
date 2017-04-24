package goset

import (
	"bytes"
	"encoding/json"
)

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

func (s Set) Has(element string) bool {
	_, ok := s[element]
	return ok
}

func (s Set) Del(element string) {
	delete(s, element)
	return
}

func (s Set) ToList() []string {
	elements := make([]string, s.Len())
	i := 0
	for k := range s {
		elements[i] = k
		i++
	}
	return elements
}

func (s *Set) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	lenOfSet := s.Len()
	i := 1
	if lenOfSet > 0 {
		buf.WriteString("[\"")
	} else {
		buf.WriteString("[")
	}
	for k := range *s {
		buf.WriteString(k)
		if i < lenOfSet {
			buf.WriteString("\",\"")
		} else {
			buf.WriteString("\"")
		}
		i++
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}

func (s *Set) UnmarshalJSON(data []byte) error {
	var list []string
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}
	s.Add(list...)
	return nil
}