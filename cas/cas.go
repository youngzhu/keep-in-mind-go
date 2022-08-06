package cas

import "sync"

// CAS: Compare and Swap
// CAS的含义是：我认为V的值应该是A。如果是，那么将V的值更新为B，
// 否则不更新并告诉V的当前值是什么。
//
// CAS是一项乐观技术，它希望能成功地执行更新操作，并且如果有另一个线程在最近一次检查后更新了该变量，
// 那么CAS能检测到这个错误。
//
// CAS典型的使用模式是：首先从V中读取值A，并根据A计算新值B，然后再通过CAS以原子方式将V的值由A更新为B。
//
// 来自《Java并发编程实践》


type Simulator struct {
	value int
	mu    sync.Mutex
}

func (s *Simulator) Get() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.value
}

func (s *Simulator) CompareAndSwap(expectedVal, newVal int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	oldVal := s.value
	if oldVal == expectedVal {
		s.value = newVal
	}
	return oldVal
}

func (s *Simulator) CompareAndSet(expectedVal, newVal int) bool {
	return expectedVal == s.CompareAndSwap(expectedVal, newVal)
}

// Counter 基于CAS实现的非阻塞的计数器
type Counter struct {
	value *Simulator
}

func NewCounter() *Counter {
	return &Counter{&Simulator{}}
}

func (c *Counter) Get() int {
	return c.value.Get()
}

func (c *Counter) Increment() int {
	for {
		current := c.Get()
		next := current + 1
		if c.value.CompareAndSet(current, next) {
			return next
		}
	}
}
