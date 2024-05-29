package iterator

import (
	"math"
	"sync/atomic"
)

type Iterator interface {
	Start(x int64) Iterator  // 起始位置
	Step(x int64) Iterator   // 单次累增长度
	Offset(x int64) Iterator // 设置偏移量
	Value() int64            // 获取当前最新分配值
}

type Iter struct {
	x    int64
	step int64
}

func New() *Iter {
	return &Iter{step: 1}
}

func (it *Iter) Start(x int64) Iterator {
	atomic.AddInt64(&it.x, x)
	return it
}

func (it *Iter) Step(x int64) Iterator {
	if x == 0 {
		x = 1
	}
	it.step = x
	return it
}

func (it *Iter) Offset(x int64) Iterator {
	atomic.SwapInt64(&it.x, x)
	return it
}

func (it *Iter) Value() int64 {
	if atomic.LoadInt64(&it.x) == math.MaxInt64 {
		atomic.SwapInt64(&it.x, 0)
	}
	atomic.AddInt64(&it.x, it.step)
	return atomic.LoadInt64(&it.x)
}

type Func func() int64

// Get 获取一个迭代器
func Get() Func {
	var x int64
	return func() int64 {
		if x >= math.MaxInt64 {
			x = 0
		}
		x++
		return x
	}
}
