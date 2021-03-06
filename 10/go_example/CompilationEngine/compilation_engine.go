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
	p.o.WriteString("<class>")

	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_004")
	}
	tknType2 := p.jt.TokenType()
	if tknType2 != JackTokenizer.TKN_IDENTIFIER {
		panic("c_005")
	}
	p.o.WriteString(fmt.Sprintf("<identifier>%s</identifier>", p.jt.Identifierr()))

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
	p.o.WriteString(fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol()))

	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_009")
	}
	tknType4 := p.jt.TokenType()

	// 类下面一层似乎之需要定义类静态声明/字段声明 或 定义方法
	// 其他都是嵌套在其他结构吧，后面碰到再说吧
	// 直到遇到他命中注定的右括号才算完美结束, 其他的和咱都不搭....
	for {
		if tknType4 == JackTokenizer.TKN_KEYWORD {
			switch p.jt.Keyword() {
			case JackTokenizer.KEY_FIELD, JackTokenizer.KEY_STATIC:
				p.CompileClassVarDec()
			case JackTokenizer.KEY_CONSTRUCTOR, JackTokenizer.KEY_FUNCTION, JackTokenizer.KEY_METHOD:
				p.CompileSubroutine()
			default:
				panic("c_011")
			}

			// 还要继续向下pick一个
			p.jt.Advance()
			tknType4 = p.jt.TokenType()
		} else if tknType4 == JackTokenizer.TKN_SYMBOL && p.jt.Symbol() == JackTokenizer.SYM_RIGHT_PARENTHESIS {
			p.o.WriteString(fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol()))
			break
		} else {
			panic("c_012")
		}
	}

	// 类终结符
	p.o.WriteString("</class>")
}

// CompileClassVarDec 编译类静态声明或字段声明
// 编译VarDec时，肯定是别人取过后判断出是这个方法出来才转过来的
// 所以一开始,咱就不p.jt.Advance()了，直接拿就行了
func (p *CompilationEngine) CompileClassVarDec() {
	p.o.WriteString(fmt.Sprintf("<keyword>%s</keyword>", p.jt.Keyword()))

	p.jt.Advance()
	if !p.jt.HasMoreTokens() {
		panic("c_013")
	}
	if p.jt.TokenType() != JackTokenizer.TKN_KEYWORD {
		panic("c_014")
	}
	if p.jt.Keyword() != JackTokenizer.KEY_INT && p.jt.Keyword() != JackTokenizer.KEY_CHAR {
		panic("c_015")
	}
	p.o.WriteString(fmt.Sprintf("<keyword>%s</keyword>", p.jt.Keyword()))

	lastFlag := 1 // 1 代表 ',' 2 代表 varName
	// 这段代码用来解析 声明 多个变量的情况
	for {
		p.jt.Advance()
		if !p.jt.HasMoreTokens() {
			panic("c_016")
		}
		switch p.jt.TokenType() {
		case JackTokenizer.TKN_IDENTIFIER:
			if lastFlag != 1 {
				panic("c_017")
			}
			p.o.WriteString(fmt.Sprintf("<identifier>%s</identifier>", p.jt.Identifierr()))
			lastFlag = 2
		case JackTokenizer.TKN_SYMBOL:
			p.o.WriteString(fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol()))
			switch p.jt.Symbol() {
			case JackTokenizer.SYM_COMMA:
				lastFlag = 1
			case JackTokenizer.SYM_SEMICOLON:
				return
			}
		default:
			panic("c_018")
		}
	}
}

// CompileSubroutine 编译方法，函数或构造函数
func (p *CompilationEngine) CompileSubroutine() {
	p.o.WriteString(fmt.Sprintf("<subroutineDec>"))
	p.o.WriteString(fmt.Sprintf("<keyword>%s</keyword>", p.jt.Keyword()))

	p.jt.Advance()
	if p.jt.HasMoreTokens() {
		panic("c_019")
	}
	if p.jt.TokenType() != JackTokenizer.TKN_IDENTIFIER {
		panic("c_020")
	}
	p.o.WriteString(fmt.Sprintf("<identifier>%s</identifier>", p.jt.Identifierr()))

	p.jt.Advance()
	if p.jt.HasMoreTokens() {
		panic("c_021")
	}
	if p.jt.TokenType() != JackTokenizer.TKN_SYMBOL {
		panic("c_022")
	}
	if p.jt.Symbol() != JackTokenizer.SYM_LEFT_BRACES {
		panic("c_023")
	}

	p.CompileParameterList() // (int a, string b)

	p.jt.Advance()
	if p.jt.HasMoreTokens() {
		panic("c_024")
	}
	if p.jt.TokenType() != JackTokenizer.TKN_SYMBOL {
		panic("c_025")
	}
	if p.jt.Symbol() != JackTokenizer.SYM_LEFT_PARENTHESIS {
		panic("c_026")
	}
	p.o.WriteString(fmt.Sprintf("<subroutineBody>"))
	p.o.WriteString(fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol())) // {
	p.CompileStatements()
	p.o.WriteString(fmt.Sprintf("<symbol>%s</symbol>", p.jt.Symbol())) // }
	p.o.WriteString(fmt.Sprintf("</subroutineBody>"))

	p.o.WriteString(fmt.Sprintf("</subroutineDec>"))
}

// CompileParameterList 编译参数列表(可能为空)不包含"(",")"
func (p *CompilationEngine) CompileParameterList() {

	lastFlag := 0 // 0 代表 "(", 1 代表 keyword( int / char ), 2 代表 identifier 3 代表 ","
	for {
		p.jt.Advance()
		if p.jt.HasMoreTokens() {
			panic("c_027")
		}

		switch p.jt.TokenType() {
		case JackTokenizer.TKN_SYMBOL:
			if lastFlag == 0 || lastFlag == 2 {
				break
			} else {
				panic("c_031")
			}
		case JackTokenizer.TKN_KEYWORD:
			if lastFlag != 0 && lastFlag != 3 {
				panic("c_028")
			}
			p.o.WriteString(fmt.Sprintf("<keyword>%s</keyword>", p.jt.Keyword()))
			continue
		case JackTokenizer.TKN_IDENTIFIER:
			if lastFlag != 1 {
				panic("c_029")
			}
			if p.jt.Keyword() != JackTokenizer.KEY_INT && p.jt.Keyword() != JackTokenizer.KEY_CHAR {
				panic("c_030")
			}
			p.o.WriteString(fmt.Sprintf("<identifier>%s</identifier>", p.jt.Identifierr()))
			continue
		}
	}
}

// CompileVarDec 编译var声明
func (p *CompilationEngine) CompileVarDec() {
	// TODO
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
