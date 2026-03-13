package visitor

import (
	"io"
	"net/url"
)

type Info struct {
	// Namespace will be set if the object is namespaced and has a specified value.
	Namespace string
	Name      string
	/*
		...
	*/
}

// visitor 接口
type Visitor interface {
	Visit(VisitorFunc) error
}

// VisitorFunc对应这个对象的方法，也就是定义中的“操作”
type VisitorFunc func(*Info, error) error

// 将多个[]Visitor封装为一个Visitor
type EagerVisitorList []Visitor

// 返回的错误暂存到[]error中，统一聚合
func (l EagerVisitorList) Visit(fn VisitorFunc) error {
	errs := []error(nil)
	for i := range l {
		if err := l[i].Visit(func(info *Info, err error) error {
			if err != nil {
				errs = append(errs, err)
				return nil
			}
			if err := fn(info, nil); err != nil {
				errs = append(errs, err)
			}
			return nil
		}); err != nil {
			errs = append(errs, err)
		}
	}
	return nil
}

type StreamVisitor struct {
	// 读取信息的来源，实现了Read这个接口，这个"流式"的概念，包括了常见的HTTP、文件、标准输入等各类输入
	io.Reader
	//*mapper

	Source string
	//Schema ContentValidator
}

func (s *StreamVisitor) Visit(fn VisitorFunc) error {
	info := &Info{
		Namespace: s.Source,
		Name:      s.Source,
	}
	return fn(info, nil)
}

// url visit
type URLVisitor struct {
	URL *url.URL
	*StreamVisitor
	// 提供错误重试次数
	HttpAttemptCount int
}

func (u *URLVisitor) Visit(fn VisitorFunc) error { return nil }
