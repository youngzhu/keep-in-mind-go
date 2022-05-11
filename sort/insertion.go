package sort

// 插入排序算法
// 将x[i]插入到x[0...i]的合适位置

// InsertionSort
// 跟之前版本比，少了赋值操作（出处：《编程珠玑》）
func InsertionSort(x []int) {
	for i := 1; i < len(x); i++ {
		t := x[i]
		j := i
		for ; j > 0 && x[j-1] > t; j-- {
			x[j] = x[j-1]
		}
		if j != i {
			x[j] = t
		}
	}
}

// 最直观的实现
func _InsertionSort(x []int) {
	for i := 1; i < len(x); i++ {
		for j := i; j > 0; j-- {
			if x[j] < x[j-1] {
				x[j], x[j-1] = x[j-1], x[j]
			} else {
				break
			}
		}
	}
}
