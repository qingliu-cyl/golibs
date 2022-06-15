package int

// Max 返回较大的数
func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// Min 返回较小的数
func Min(i, j int) int {
	if i > j {
		return j
	}

	return i
}