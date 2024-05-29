package iterator

import "testing"

func TestIter_Get(t *testing.T) {
	iter := New().Start(0).Step(1)
	for i := 0; i < 1000; i++ {
		println(iter.Value())
	}
}

func TestIter_Offset(t *testing.T) {
	iter := New().Start(0).Offset(1e10)
	println(iter.Value())
}

func TestIterFunc(t *testing.T) {
	iter := Get()
	for i := 0; i < 10; i++ {
		println(iter())
	}
}
