package interpreter

// AbstractExpression 抽象表达式
type AbstractExpression interface {
	// 解释的操作
	Interpret(ctx Context)
}

// Context 上下文，包含解释器之外的一些全局信息
type Context struct {}

// TerminalExpression 终结符表达式
type TerminalExpression struct {}

func (t *TerminalExpression) Interpret(ctx Context) {
	// 实现与语法规则中的终结符相关联的解释操作
}

// NonTerminalExpression 非终结符表达式
type NonTerminalExpression struct {}

func (t *NonTerminalExpression) Interpret(ctx Context) {
	// 实现与语法规则中的终结符相关联的解释操作
}