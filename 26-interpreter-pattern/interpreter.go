package interpreter

import (
	"strconv"
	"strings"
)

/*
	使用解释器模式实现加减法
*/
// 定义接口表达式
type Expression interface {
	Interpret() int
}

// 实现具体的类

// 整数表达式
type NumExpression struct {
	val int
}

// 解析整数值
func (n *NumExpression) Interpret() int {
	return n.val
}

// 加法表达式
type AddExpression struct {
	left, right Expression
}

// 解释-加法操作
func (n *AddExpression) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

// 减法表达式
type SubExpression struct {
	left, right Expression
}

// 解释-减法操作
func (n *SubExpression) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

// 定义解析器
type Parser struct {
	exp   []string
	index int //exp游标
	prev  Expression
}

func (p *Parser) Parse(str string) {
	p.exp = strings.Split(str, " ")
	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddExpression()
		case "-":
			p.prev = p.newSubExpression()
		default:
			p.prev = p.newNumExpression()
		}
	}
}

func (p *Parser) newAddExpression() Expression {
	p.index++
	return &AddExpression{
		left:  p.prev,
		right: p.newNumExpression(),
	}
}

func (p *Parser) newSubExpression() Expression {
	p.index++
	return &SubExpression{
		left:  p.prev,
		right: p.newNumExpression(),
	}
}

func (p *Parser) newNumExpression() Expression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &NumExpression{
		val: v,
	}
}

func (p *Parser) Result() Expression {
	return p.prev
}
