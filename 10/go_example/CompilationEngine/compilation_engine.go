package CompilationEngine

import (
	jt "go_example/JackTokenizer"
	"os"
)

// CompilationEngine ...
type CompilationEngine struct {
	jt jt.JackTokenizer
	o  *os.File
}

// NewCompilationEngine ...
func NewCompilationEngine(jt jt.JackTokenizer, fo *os.File) *CompilationEngine {
	return &CompilationEngine{jt, fo}
}

// CompileClass 编译整个类
func (p *CompilationEngine) CompileClass() {

}

// CompileClassVarDec 编译类静态声明或字段声明
func (p *CompilationEngine) CompileClassVarDec() {

}

// CompileSubroutine 编译方法，函数或构造函数
func (p *CompilationEngine) CompileSubroutine() {

}

// CompileParameterList 编译参数列表(可能为空)不包含"(",")"
func (p *CompilationEngine) CompileParameterList() {

}

// CompileVarDec 编译var声明
func (p *CompilationEngine) CompileVarDec() {

}

// CompileStatements 编译一系列语句，不包含大括号 "{", "}"
func (p *CompilationEngine) CompileStatements() {

}

// CompileDo 编译do语句
func (p *CompilationEngine) CompileDo() {

}

// CompileLet 编译let语句
func (p *CompilationEngine) CompileLet() {

}

// CompileWhile 编译while语句
func (p *CompilationEngine) CompileWhile() {

}

// CompileReturn 编译return语句
func (p *CompilationEngine) CompileReturn() {

}

// CompileIf 编译if语句,包含可选的else从句
func (p *CompilationEngine) CompileIf() {

}

// CompileExpression 编译一个表达式
func (p *CompilationEngine) CompileExpression() {

}

// CompileTerm 编译一个term
func (p *CompilationEngine) CompileTerm() {

}

// CompileExpressionList 编译逗号分割的表达式列表(可能为空)
func (p *CompilationEngine) CompileExpressionList() {

}
