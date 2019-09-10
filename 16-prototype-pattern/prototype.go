package prototype

/*
	通过将已经存在的实例赋值给新的变量来完成clone, 可定制clone对象
*/
type Example struct {
	Description string
}

//实现Clone
func (e *Example) Clone() *Example {
	res := *e
	return &res
}

func New(des string) *Example{
	return &Example{
		Description: des,
	}
}
