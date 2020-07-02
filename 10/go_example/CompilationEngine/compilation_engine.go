package CompilationEngine

import (
	"fmt"
	"go_example/JackTokenizer"
	jt "go_example/JackTokenizer"
	"os"
)

// CompilationEngine ...
type CompilationEngine struct {
	jt *jt.JackTokenizer
	o  *os.File
}

// NewCompilationEngine ...
func NewCompilationEngine(jt *jt.JackTokenizer, fo *os.File) *CompilationEngine {
	return &CompilationEngine{jt, fo}
}

// CompileClass 编译整个类
func (p *CompilationEngine) CompileClass() {
	outputstring := ""
	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_001")
	}
	tknType := p.jt.TokenType()
	if tknType != JackTokenizer.TKN_KEYWORD {
		panic("c_002")
	}
	smb := p.jt.Keyword()
	if smb != JackTokenizer.KEY_CLASS {
		panic("c_003")
	}
	outputstring += "<class>"
	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_004")
	}
	tknType2 := p.jt.TokenType()
	if tknType2 != JackTokenizer.TKN_IDENTIFIER {
		panic("c_005")
	}
	outputstring += fmt.Sprintf("<identifier>%s</identifier>", p.jt.Identifierr())
	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_006")
	}
	tknType3 := p.jt.TokenType()
	if tknType3 != JackTokenizer.TKN_SYMBOL {
		panic("c_007")
	}
	if p.jt.Symbol() != JackTokenizer.SYM_LEFT_PARENTHESIS {
		panic("c_008")
	}
	outputstring += fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol())

	// TODO

	outputstring += "</class>"
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
