package josephus

import "github.com/youngzhu/algs4-go/fund"

// 约瑟夫环，优雅的解法
// 不过要借助数据结构——队列
// 原地址：https://algs4.cs.princeton.edu/13stacks/Josephus.java

// Solve 优雅地解决约瑟夫环
// m - 第m个退出
// n - 总人数
func Solve(m, n int) (result []int) {
	// 队列，先进先出
	queue := fund.NewQueue()
	for i := 0; i < n; i++ {
		queue.Enqueue(i)
	}

	for !queue.IsEmpty() {
		for i := 1; i < m; i++ {
			// 形成环
			queue.Enqueue(queue.Dequeue())
		}
		result = append(result, queue.Dequeue().(int))
	}

	return result
}
