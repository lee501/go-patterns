package iterator

import "testing"

func TestIterator_Next(t *testing.T) {
	teacher := new(Teacher)
	analysis := new(Analysis)
	//迭代器
	iterator := NewIterator()
	iterator.Add(teacher)
	iterator.Add(analysis)
	if len(iterator.list) != 2 {
		t.Error("期望的count is 2")
	}
	for ; iterator.HasNext(); {
		iterator.Next().Visit()
	}
}
