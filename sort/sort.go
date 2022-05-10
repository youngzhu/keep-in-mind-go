package sort

// 各类排序算法最直接的实现
// **不包含**实际应用中可能需要的优化，如快速排序时，如果切片长度小于一定数值，使用插入排序会更高效

// InsertionSort 插入排序
func InsertionSort(x []int) {
	n := len(x)
	for i := 1; i < n; i++ {
		j := i
		for ; j > 0; j-- {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			}
		}
	}
}
