package tokr

func (doc document) Find(find document) bool {
	mm := make(map[string]struct{})
	for _, w := range find {
		mm[w] = struct{}{}
	}

	var ok bool
	for _, w := range doc {
		_, ok = mm[w]
		if ok {
			return true
		}
	}

	return false
}
