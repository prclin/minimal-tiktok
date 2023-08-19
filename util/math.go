package util

func Intersection(ar, br []uint64) []uint64 {
	is := make([]uint64, 0, len(ar))
	hm := make(map[uint64]*struct{}, len(ar))
	for _, item := range ar {
		hm[item] = &struct{}{}
	}
	for _, item := range br {
		if hm[item] != nil {
			is = append(is, item)
		}
	}
	return is
}
