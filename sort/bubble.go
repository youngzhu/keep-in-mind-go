package sort

/*
冒泡排序，又叫选择排序，首先选出最小的元素，再选出剩下的元素中最小的，依此类推

将数组顺时针旋转90度，想象成一个水桶，越小的气泡（数字）越快往上浮
         ****   /\
    a[0] *  *   ||
         ****   ||
         *  *   ||
         ****   ||
         *  *   ||
         ****   ||
    a[n] *  *   ||
         ****
*/

func BubbleSort(x []int) {
	n := len(x)
	for i := 0; i < n; i++ { // 为x[i]找到正确的元素
		idx := i
		minFromI := x[idx] // x[i...n-1]最小的元素
		for j := i + 1; j < n; j++ {
			if x[j] < minFromI {
				idx = j
				minFromI = x[j]
			}
		}
		if idx > i {
			x[i], x[idx] = x[idx], x[i]
		}
	}
}
