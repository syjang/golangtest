package math

func AbsInt(v int) int {
	if v < 0 {
		v = -v
	}

	return v
}

func AbsInt64(v int64) int64 {
	if v < 0 {
		v = -v
	}
	return v
}

func AbsInt32(v int32) int32 {
	if v < 0 {
		v = -v
	}
	return v
}
