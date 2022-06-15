package slice

// IArrayCopy copy一个int型二维数组
func IArrayCopy(board [][]int) [][]int {
	if len(board) == 0 {
		return [][]int{}
	}
	cow := len(board)
	res := make([][]int, cow)

	for i := 0; i < cow; i++ {
		dst := make([]int, len(board[i]))
		copy(dst, board[i])
		res[i] = dst
	}

	return res
}
