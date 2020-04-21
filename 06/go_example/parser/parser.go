package parser

import (
	"fmt"
	"go_example/common"
	"os"
	"strings"
)

// Parser 解析器
type Parser struct {
	CurrentCommandIndex int      // 当前指令序号
	CurrentMemoryIndex  int      // 内存当前偏移
	Commands            []string // 指令序列
	CurrentReadLine     int      // 当前读取的行数
}

// NewParser 创建新的解析器
func NewParser(commands []string) *Parser {
	return &Parser{
		CurrentMemoryIndex:  16,
		Commands:            commands,
		CurrentCommandIndex: 0,
		CurrentReadLine:     0,
	}
}

// HasMoreCommands 输入中还有更多的命令吗
func (p *Parser) HasMoreCommands() bool {
	return p.CurrentReadLine < len(p.Commands)-1
}

// Advance 从输入中读取下一条命令，将其当作当前命令
// 仅当HasMoreCommands()为真时,才能调用本程序。
// 最初始的时候没有当前命令
func (p *Parser) Advance() {
	if p.HasMoreCommands() {
		p.CurrentReadLine = p.CurrentReadLine + 1
		if p.CommandType() != common.L_COMMAND {
			p.CurrentCommandIndex = p.CurrentCommandIndex + 1
		}
	}
	return
}

// CommandType 返回当前命令的类型:
// A_COMMAND 当@xxx中的xxx是符号或者是十进制数字时
// C_COMMAND 用于dest=comp;jmp
// L_COMMAND 实际上是伪指令 当(xxx)中的xxx是符号时
func (p *Parser) CommandType() int {
	if isACommand(p.Commands[p.CurrentReadLine]) {
		return common.A_COMMAND
	}

	if isCCommand(p.Commands[p.CurrentReadLine]) {
		return common.C_COMMAND
	}

	if isLCommand(p.Commands[p.CurrentReadLine]) {
		return common.L_COMMAND
	}

	fmt.Println("e1")
	os.Exit(-1)
	return 0
}

// Symbol 返回形如(xxx) 或 @xxx 的当前命令的符号或者十进制
// 仅当 CommandType() 是 A_COMMAND 或者 L_COMMAND 才能调用
func (p *Parser) Symbol() string {
	if p.CommandType() == common.A_COMMAND {
		return p.Commands[p.CurrentReadLine][1:]
	}

	if p.CommandType() == common.L_COMMAND {
		l := len(p.Commands[p.CurrentReadLine])
		return p.Commands[p.CurrentReadLine][1 : l-1]
	}

	fmt.Println("e2")
	os.Exit(-1)
	return ""
}

// Dest 返回当前C指令的 dest 助记符(具有8种可能的形式)
// 当 CommandType() 是 C_COMMAND 时才能使用
func (p *Parser) Dest() string {

	if p.CommandType() == common.C_COMMAND {
		items := strings.Split(p.Commands[p.CurrentReadLine], "=")
		if len(items) == 1 {
			return ""
		}
		return items[0]
	}

	fmt.Println("e3")
	os.Exit(-1)
	return ""
}

// Comp 返回当前 C_COMMAND 的 comp助记符(具有28种可能的形式)
// 当 CommandType() 是 C_COMMAND 时才能使用
func (p *Parser) Comp() string {
	if p.CommandType() == common.C_COMMAND {
		items := strings.Split(p.Commands[p.CurrentReadLine], "=")
		subCommand := ""
		if len(items) == 1 {
			subCommand = items[0]
		} else {
			subCommand = items[1]
		}
		subItems := strings.Split(subCommand, ";")
		return subItems[0]
	}

	fmt.Println("e4")
	os.Exit(-1)
	return ""
}

// Jump 返回当前C指令的 jump 助记符(具有8种可能的形式)
// 当 CommandType() 是 C_COMMAND 时才能使用
func (p *Parser) Jump() string {
	if p.CommandType() == common.C_COMMAND {
		items := strings.Split(p.Commands[p.CurrentReadLine], "=")
		subCommand := ""
		if len(items) == 1 {
			subCommand = items[0]
		} else {
			subCommand = items[1]
		}
		subItems := strings.Split(subCommand, ";")

		if len(subItems) == 1 {
			return ""
		}
		return subItems[1]
	}

	fmt.Println("e5")
	os.Exit(-1)
	return ""
}

// 是否是a指令
func isACommand(command string) bool {
	if strings.HasPrefix(command, "@") {
		return true
	}
	return false
}

// 是否是c指令
func isCCommand(command string) bool {
	i := strings.Index(command, "=")
	j := strings.Index(command, ";")
	return i > 0 || j > 0
}

// 是否是l指令
func isLCommand(command string) bool {
	if strings.HasPrefix(command, "(") && strings.HasSuffix(command, ")") {
		return true
	}
	return false
}
