package tokr

func (doc document) Count(find document) int {
	mm := make(map[string]struct{})
	for _, w := range find {
		mm[w] = struct{}{}
	}

	var ok bool
	n := 0
	for _, w := range doc {
		_, ok = mm[w]
		if ok {
			n++
		}
	}

	return n
}
