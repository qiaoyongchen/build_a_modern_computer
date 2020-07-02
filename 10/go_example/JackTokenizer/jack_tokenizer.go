package JackTokenizer

import "os"

type TknType uint8

const (
	TKN_KEYWORD      TknType = 1
	TKN_SYMBOL       TknType = 2
	TKN_IDENTIFIER   TknType = 3
	TKN_INT_CONST    TknType = 4
	TKN_STRING_CONST TknType = 5
)

type KeywordType string

const (
	KEY_CLASS       KeywordType = "CLASS"
	KEY_METHOD      KeywordType = "METHOD"
	KEY_INT         KeywordType = "INT"
	KEY_FUNCTION    KeywordType = "FUNCTION"
	KEY_BOOLEAN     KeywordType = "BOOLEAN"
	KEY_CONSTRUCTOR KeywordType = "CONSTRUCTOR"
	KEY_CHAR        KeywordType = "CHAR"
	KEY_VOID        KeywordType = "VOID"
	KEY_VAR         KeywordType = "VAR"
	KEY_STATIC      KeywordType = "STATIC"
	KEY_FIELD       KeywordType = "FIELD"
	KEY_LET         KeywordType = "LET"
	KEY_DO          KeywordType = "DO"
	KEY_IF          KeywordType = "IF"
	KEY_ELSE        KeywordType = "ELSE"
	KEY_WHILE       KeywordType = "WHILE"
	KEY_RETURN      KeywordType = "RETURN"
	KEY_TRUE        KeywordType = "TRUE"
	KEY_FALSE       KeywordType = "FALSE"
	KEY_NULL        KeywordType = "NULL"
	KEY_THIS        KeywordType = "THIS"
)

type SymbolType string

const (
	SYM_LEFT_PARENTHESIS  = "{"
	SYM_RIGHT_PARENTHESIS = "}"
	SYM_LEFT_BRACES       = "("
	SYM_RIGHT_BRACES      = ")"
	SYM_LEFT_BRACKET      = "["
	SYM_RIGHT_BRACKET     = "]"
	SYM_DOT               = "."
	SYM_COMMA             = ","
	SYM_SEMICOLON         = ";"
	SYM_PLUS              = "+"
	SYM_MINUS             = "-"
	SYM_MUL               = "*"
	SYM_DIV               = "-"
	SYM_COMBINE           = "&"
	SYM_VERTICAL          = "|"
	SYM_GT                = ">"
	SYM_LT                = "<"
	SYM_EQ                = "="
	SYM_WAVE              = "~"
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
