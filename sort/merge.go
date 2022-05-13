package sort

// 归并排序（分治法）
// 最常见的是自顶向下法（top-down），即向数组分成2部分分别排序，之后再合并。递归调用
// 还有自底向上法（bottom-up），即，1-1，2-2，4-4 ... 依次合并

func MergeSort(x []int) {
	// 这里不考虑效率，使用常见的自顶向下归并法
	topDown(x)

	//bottomUp(x) // 测试
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
	merge(x, aux, lo, mid+1, hi)
}

// 辅助切片的作用体现在这里
// 将两个有序的数组合并
// 通过aux将x[lo...mid]和x[mid+1...hi]合并为有序数组
// x[lo...mid]和x[mid+1...hi]各自有序
func merge(x []int, aux []int, lo int, mid int, hi int) {
	// copy to aux
	copy(aux[lo:hi+1], x[lo:hi+1])

	// merge back to x[]
	i, j := lo, mid // 双指针
	k := lo
	// 交叉
	for i < mid && j <= hi {
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
	for i < mid {
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

/******************************************************
自底向上合并法
跟自顶向下比，少了个排序的过程
自顶向下：首先要对x[lo...mid]和x[mid+1...hi]排序
自底向上：将每一个元素看作一个子数组，一步步向上合并，1-1，2-2 ...
*******************************************************/
func bottomUp(x []int) {
	n := len(x)
	aux := make([]int, n)
	var subLength = 1 // 子数组的长度，从1开始
	for ; subLength < n; subLength *= 2 {
		idx1, idx2 := 0, subLength // 待合并的2部分的起始索引
		var maxIdx int             // 最大索引（包含）
		for idx2 < n {
			maxIdx = min(idx2+subLength-1, n-1)
			merge(x, aux, idx1, idx2, maxIdx)
			idx1 = maxIdx + 1
			idx2 = idx1 + subLength
		}
	}
}

func min(i int, j int) int {
	if i > j {
		return j
	}
	return i
}
