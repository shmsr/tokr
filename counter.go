package tokr

func (doc document) Count(find document) int {
	mm := make(map[string]struct{})
	for _, w := range find {
		mm[w] = struct{}{}
	}

	n, ok := 0, false
	for _, w := range doc {
		if _, ok = mm[w]; ok {
			n++
		}
	}

	return n
}
