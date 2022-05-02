package unsettled

// DistributeEqually 将一个数组切片（original）分割成n等份
// （尽量平均分配）
func DistributeEqually(original []int, n int) (parts [][]int) {
	parts = make([][]int, n)

	total := len(original)
	for i := 0; i < n; i++ {
		for j := i * total / n; j < (i+1)*total/n; j++ {
			parts[i] = append(parts[i], original[j])
		}
	}

	return
}
