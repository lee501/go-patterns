package interpreter

import "testing"

func TestParser_Result(t *testing.T) {
	p := &Parser{}
	p.Parse("1 + 3 + 3 + 3 - 3")
	e := p.Result()
	res := e.Interpret()
	t.Run("test result", func(t *testing.T) {
		if res != 7 {
			t.Errorf("test failed, return %v, want 7", res)
		}
	})
}
