package JackTokenizer

import "os"

type TknType uint8

const (
	KEYWORD      TknType = 1
	SYMBOL       TknType = 2
	IDENTIFIER   TknType = 3
	INT_CONST    TknType = 4
	STRING_CONST TknType = 5
)

type KeywordType string

const (
	CLASS       KeywordType = "CLASS"
	METHOD      KeywordType = "METHOD"
	INT         KeywordType = "INT"
	FUNCTION    KeywordType = "FUNCTION"
	BOOLEAN     KeywordType = "BOOLEAN"
	CONSTRUCTOR KeywordType = "CONSTRUCTOR"
	CHAR        KeywordType = "CHAR"
	VOID        KeywordType = "VOID"
	VAR         KeywordType = "VAR"
	STATIC      KeywordType = "STATIC"
	FIELD       KeywordType = "FIELD"
	LET         KeywordType = "LET"
	DO          KeywordType = "DO"
	IF          KeywordType = "IF"
	ELSE        KeywordType = "ELSE"
	WHILE       KeywordType = "WHILE"
	RETURN      KeywordType = "RETURN"
	TRUE        KeywordType = "TRUE"
	FALSE       KeywordType = "FALSE"
	NULL        KeywordType = "NULL"
	THIS        KeywordType = "THIS"
)

// JackTokenizer ...
type JackTokenizer struct {
	I *os.File
}

// NewJackTokenizer 打开输入输入文件准备进行字元转换操作
func NewJackTokenizer(fi *os.File) *JackTokenizer {
	return &JackTokenizer{fi}
}

// HasMoreTokens 输入中是否还有字元
func (p *JackTokenizer) HasMoreTokens() bool {
	return true
}

// Advance 从输入中读取下一个字元，使其成为当前字元
// 该函数仅当hasMoreTokens()返回i真时才能调用
// 最初始状态没有当前字元
func (p *JackTokenizer) Advance() {
	return
}

// TokenType 返回当前字元的类型
func (p *JackTokenizer) TokenType() TknType {
	return 0
}

// Keyword 返回当前资源的关键字，仅当tokenType()返回KEYWORD时才被调用
func (p *JackTokenizer) Keyword() KeywordType {
	return ""
}

// Symbol 返回当前字元的字符，仅当tokenType()的返回值为SYMBOL时才被调用
func (p *JackTokenizer) Symbol() string {
	return ""
}

// Identifierr 返回当前字元的字符，仅当tokenType()的返回值为IDENTIFIER时才被调用
func (p *JackTokenizer) Identifierr() string {
	return ""
}

// IntVal 返回当前字元的整数值，仅当tokenType()的返回值为INT_CONST时才被调用
func (p *JackTokenizer) IntVal() int {
	return 0
}

// StringVal 返回当前字元的整数值，仅当tokenType()的返回值为STRING_CONST时才被调用
func (p *JackTokenizer) StringVal() int {
	return 0
}
