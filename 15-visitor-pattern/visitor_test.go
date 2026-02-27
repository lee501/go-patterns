package visitor

import (
	"errors"
	"net/url"
	"strings"
	"testing"
)

// mockVisitor 是一个简单的 Visitor 实现，用于测试
type mockVisitor struct {
	info *Info
}

func (m *mockVisitor) Visit(fn VisitorFunc) error {
	return fn(m.info, nil)
}

func TestEagerVisitorList(t *testing.T) {
	visited := []string{}
	list := EagerVisitorList{
		&mockVisitor{info: &Info{Namespace: "ns1", Name: "obj1"}},
		&mockVisitor{info: &Info{Namespace: "ns2", Name: "obj2"}},
	}
	err := list.Visit(func(info *Info, err error) error {
		visited = append(visited, info.Name)
		return nil
	})
	if err != nil {
		t.Errorf("EagerVisitorList.Visit 不应返回错误, 实际返回: %v", err)
	}
	if len(visited) != 2 {
		t.Errorf("应访问2个对象, 实际访问了%d个", len(visited))
	}
}

func TestEagerVisitorList_WithError(t *testing.T) {
	list := EagerVisitorList{
		&mockVisitor{info: &Info{Name: "obj1"}},
		&mockVisitor{info: &Info{Name: "obj2"}},
	}
	// 第一个访问者返回错误，第二个应该仍然被访问
	count := 0
	err := list.Visit(func(info *Info, err error) error {
		count++
		if info.Name == "obj1" {
			return errors.New("访问obj1出错")
		}
		return nil
	})
	// EagerVisitorList 将错误聚合，顶层返回 nil
	if err != nil {
		t.Errorf("EagerVisitorList.Visit 应返回nil, 实际返回: %v", err)
	}
	if count != 2 {
		t.Errorf("即使有错误，也应访问所有对象, 实际访问了%d个", count)
	}
}

func TestStreamVisitor(t *testing.T) {
	r := strings.NewReader("test data")
	sv := &StreamVisitor{Reader: r, Source: "test-source"}
	var visited *Info
	err := sv.Visit(func(info *Info, err error) error {
		visited = info
		return nil
	})
	if err != nil {
		t.Errorf("StreamVisitor.Visit 不应返回错误, 实际返回: %v", err)
	}
	if visited == nil {
		t.Error("StreamVisitor.Visit 应该调用 VisitorFunc")
	}
	if visited.Name != "test-source" {
		t.Errorf("Info.Name 应为 test-source, 实际为 %s", visited.Name)
	}
}

func TestURLVisitor(t *testing.T) {
	u, _ := url.Parse("http://example.com")
	r := strings.NewReader("")
	sv := &StreamVisitor{Reader: r, Source: "http://example.com"}
	uv := &URLVisitor{URL: u, StreamVisitor: sv, HttpAttemptCount: 3}
	err := uv.Visit(func(info *Info, err error) error {
		return nil
	})
	if err != nil {
		t.Errorf("URLVisitor.Visit 不应返回错误, 实际返回: %v", err)
	}
}

