package pool

import (
	"fmt"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestNewPool(t *testing.T) {
	p := NewPool(5)
	if len(*p) != 5 {
		t.Error("线程池构造错误")
	}
	// 从池中获取对象并使用
	obj := p.Acquire()
	fmt.Println(len(*p))
	if len(*p) != 4 {
		t.Errorf("Acquire后池中应有4个对象，实际有%d个", len(*p))
	}
	obj.Do()
	// 归还对象到池中
	p.Release(obj)
	if len(*p) != 5 {
		t.Errorf("Release后池中应有5个对象，实际有%d个", len(*p))
	}
}

func TestPoolConcurrent(t *testing.T) {
	p := NewPool(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			obj := p.Acquire()
			obj.Do()
			p.Release(obj)
		}()
	}
	wg.Wait()
	if len(*p) != 3 {
		t.Errorf("并发使用后池中应有3个对象，实际有%d个", len(*p))
	}
}
