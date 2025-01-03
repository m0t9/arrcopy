package testdata

func iterateOverCopy() {
	arr := [228228]struct {
		s string
		v int64
		a [123]int64
	}{}

	for idx, item := range arr { // want "for-range loop over array 'arr' found. Use for-range over '&arr' instead"
		_ = idx
		_ = item
	}
}

func iterateOverPtr() {
	arr := [228228]struct {
		s string
		v int64
		a [123]int64
	}{}

	for idx, item := range &arr {
		_ = idx
		_ = item
	}
}
