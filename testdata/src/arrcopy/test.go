package arrcopy

func iterateOverCopy() {
	type I struct {
		s string
		v int64
		a [123]int64
	}
	arr := [228228]I{}

	for idx, item := range arr { // want `for-range loop over array 'arr' found. Use for-range over '&arr' instead`
		_ = idx
		_ = item
	}

	for idx, item := range &arr {
		_ = idx
		_ = item
	}

	f := func() [228228]I {
		return arr
	}

	// Optimization is not applicable to funcion & method calls.
	for idx, item := range f() {
		_ = idx
		_ = item
	}

	// And for cast expressions.
	var anyArr any = arr
	for idx, item := range anyArr.([228228]I) {
		_ = idx
		_ = item
	}

	type S struct {
		a [228228]I
	}

	s := S{}

	for idx, item := range s.a { // want "for-range loop over array 's.a' found. Use for-range over '&s.a' instead"
		_ = idx
		_ = item
	}

	for idx, item := range [1]int{1} { // want `for-range loop over array '\[1\]int\{1\}' found\. Use for-range over '&\[1\]int\{1\}' instead`
		_ = idx
		_ = item
	}

	// Ignore cases without array's items usage.
	for idx := range [1]int{1} {
		_ = idx
	}

	for idx, _ := range [1]int{1} {
		_ = idx
	}
}
