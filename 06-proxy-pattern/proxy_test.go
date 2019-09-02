package proxy

import "testing"

func TestProxyObject(t *testing.T) {
	object := &Object{action: "run"}

	proxyObject :=  new(ProxyObject)
	proxyObject.object = object
	proxyObject.ObjDo("run")
}
