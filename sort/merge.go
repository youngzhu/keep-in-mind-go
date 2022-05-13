package sort

// 归并排序（分治法）
// 最常见的是自顶向下法（top-down），即向数组分成2部分分别排序，之后再合并。递归调用
// 还有自底向上法（bottom-up），即，1-1，2-2，4-4 ... 依次合并

func MergeSort(x []int) {
	// 这里不考虑效率，使用常见的自顶向下归并法
	topDown(x)
}

// 使用一个辅组切片能是算法更简单
func topDown(x []int) {
	n := len(x)
	// 看似多余，只有最后合并的时候才用到
	// 但是如果不在这里创建，后面会有很多次make
	aux := make([]int, n)
	mergeSort(x, aux, 0, n-1)
}

func mergeSort(x []int, aux []int, lo int, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(x, aux, lo, mid)
	mergeSort(x, aux, mid+1, hi)
	merge(x, aux, lo, mid, hi)
}

// 辅助切片的作用体现在这里
// 通过aux将x[lo...mid]和x[mid+1...hi]合并为有序数组
// x[lo...mid]和x[mid+1...hi]各自有序
func merge(x []int, aux []int, lo int, mid int, hi int) {
	// copy to aux
	copy(aux, x)

	// merge back to x[]
	i, j := lo, mid+1 // 双指针
	k := lo
	// 交叉
	for i <= mid && j <= hi {
		if aux[i] > aux[j] {
			x[k] = aux[j]
			j++
		} else {
			x[k] = aux[i]
			i++
		}
		k++
	}

	// 只剩左半边，依次复制
	for i <= mid {
		x[k] = aux[i]
		i++
		k++
	}
	// 只剩右半边，依次复制
	for j <= hi {
		x[k] = aux[j]
		j++
		k++
	}
}
