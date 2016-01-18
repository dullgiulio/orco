package main

type rowDiff map[int]string

func makeRowDiff() rowDiff {
	return make(map[int]string)
}

func makeDiffForRow(r row) rowDiff {
	d := makeRowDiff()
	for i := 0; i < len(r); i++ {
		d[i] = r[i]
	}
	return d
}

type row []string

func makeRow(s ...string) row {
	return row(s)
}

// diff computes the difference between r and o,
// returning all elements of o that differ from r.
// When all elements match returns nil.
// Diff also returns nil if len(r) != len(o).
func (r row) diff(o row) rowDiff {
	if o == nil {
		return makeDiffForRow(r)
	}
	if r == nil {
		return makeDiffForRow(o)
	}
	if len(r) != len(o) {
		return nil
	}
	diff := makeRowDiff()
	for i := 0; i < len(r); i++ {
		if r[i] != o[i] {
			diff[i] = o[i]
		}
	}
	if len(diff) == 0 {
		return nil
	}
	return diff
}
