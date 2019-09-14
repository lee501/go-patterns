package template

import (
	"testing"
)

func TestBoy_BeforeAction(t *testing.T) {
	boy := &Boy{}
	person := new(Person)
	person.SetName("boy")
	person.Concrete = boy
	//赋值boy的内容, 注意要在设定了person具体值之后赋值，否则为空
	boy.Person = *person
	person.Exit()
}

func TestGirl_BeforeAction(t *testing.T) {
	girl := &Girl{}
	person := new(Person)
	person.Concrete = girl
	person.SetName("girl")
	//赋值girl的内容
	girl.Person = *person
	person.Exit()
}
