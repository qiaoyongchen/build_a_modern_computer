package parser

import (
	"fmt"
	"os"
	"strings"
)

// Parser 解析器
type Parser struct {
	file                *os.File
	CurrentCommand      string
	CurrentCommandIndex int
	Commands            []string
}

// NewParser 构造函数
// inputFile 输入文件流
func NewParser(inputFile string, commands []string) *Parser {
	p := &Parser{Commands: commands}

	file, fileErr := os.Open(inputFile)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		os.Exit(-1)
	}
	p.file = file

	return p
}

// HasMoreCommands 输入中还有更多的命令吗
func (p *Parser) HasMoreCommands() bool {
	return p.CurrentCommandIndex < (len(p.Commands) - 1)
}

// Advance 从输入读取下一条指令,并将其指定为当前指令
// 仅当 HasMoreCommands 为真时才能调用该命令
func (p *Parser) Advance() {
	p.CurrentCommandIndex++
	p.CurrentCommand = p.Commands[p.CurrentCommandIndex]
}

const (
	C_NONE       = 0
	C_ARITHMETIC = 1
	C_PUSH       = 2
	C_POP        = 3
	C_LABEL      = 4
	C_GOTO       = 5
	C_IF         = 6
	C_FUNCTION   = 7
	C_RETURN     = 8
	C_CALL       = 9
)

// CommandType 返回VM命令的类型, 对于所有算术命令总是返回 C_ARITHMETIC
func (p *Parser) CommandType() int {
	commandParts := strings.Split(p.CurrentCommand, " ")

	switch strings.ToLower(commandParts[0]) {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return C_ARITHMETIC
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	}

	return C_NONE
}

// Arg1 返回当前类型的地一个参数
// 如果当前类型为 C_ARITHMETIC 则返回命令本身(如 add sub)
// 如果当前类型为 C_RETURN 不应该调用本程序
func (p *Parser) Arg1() string {
	ct := p.CommandType()
	commandParts := strings.Split(p.CurrentCommand, " ")

	if ct == C_NONE {
		return ""
	}

	switch ct {
	case C_ARITHMETIC:
		return commandParts[0]
	case C_RETURN:
		return ""
	default:
		return commandParts[1]
	}
}

// Arg2 返回当前命令的第二个参数
// 当前命令为 C_PUSH C_POP C_FUNCTION C_CALL 时才可调用
func (p *Parser) Arg2() string {
	ct := p.CommandType()
	commandParts := strings.Split(p.CurrentCommand, " ")
	if ct == C_PUSH || ct == C_POP || ct == C_FUNCTION || ct == C_CALL {
		return commandParts[2]
	}
	return ""
}
