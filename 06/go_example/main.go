package main

import (
	"bufio"
	"fmt"
	"go_example/code"
	"go_example/common"
	"go_example/parser"
	"go_example/systemTable"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	buf := bufio.NewReader(f)
	lines := []string{}
	breakFlag := false
	for {
		// 按行读取文件
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR != io.EOF {
				fmt.Println(errR.Error())
				os.Exit(-1)
			}
			breakFlag = true
		}

		// 清除不必要的东西比如注释和空格 并 格式化汇编命令
		line := clearLine(string(b))
		if line == "" {
			continue
		}
		lines = append(lines, line)
		if breakFlag {
			break
		}
	}

	for _, v := range lines {
		fmt.Printf("%s\n", v)
	}

	fp := parser.NewParser(lines) // 第一遍的解析器
	p := parser.NewParser(lines)  // 第二遍的解析器
	c := code.NewCode()           // 解码器
	st := systemTable.Construct() // 符号表

	binLines := []string{}

	// 第一遍只处理符号不处理代码,用来收集符号在，并放置到相应内存处
	for true {
		if fp.CommandType() == common.L_COMMAND {
			if st.Contains(fp.Symbol()) {
				fmt.Println("a1")
				os.Exit(-1)
			}
			st.AddEntry(fp.Symbol(), fp.CurrentCommandIndex+1)
		}

		if fp.HasMoreCommands() {
			fp.Advance()
		} else {
			break
		}
	}

	// 打印扫描的内存存储值
	for k, v := range st.Tabs {
		fmt.Println("符号 : " + k + ", 值 : " + strconv.Itoa(v))
	}

	// 第二遍处理代码并关联符号
	for true {
		tmpLine := ""

		if p.CommandType() == common.A_COMMAND {
			// A 指令
			fmt.Printf("a_commond : line / %s \n", p.Commands[p.CurrentReadLine])

			address := 0
			i, ie := strconv.Atoi(p.Symbol())
			if ie != nil {
				if !st.Contains(p.Symbol()) {
					fmt.Println("a2")
					os.Exit(-1)
				}
				address = st.GetAddress(p.Symbol())
			} else {
				address = i
			}
			tmpLine = intToBinStr(address)
		} else if p.CommandType() == common.C_COMMAND {
			// C 指令 调试输出
			fmt.Printf(
				"c_commond : line / %s, dest / %s|%s, comp / %s|%s, jump: %s|%s \n",
				p.Commands[p.CurrentReadLine],
				p.Dest(),
				c.Dest(p.Dest()),
				p.Comp(),
				c.Comp(p.Comp()),
				p.Jump(),
				c.Jump(p.Jump()),
			)
			// 生成的c指令
			tmpLine = "111" + c.Comp(p.Comp()) + c.Dest(p.Dest()) + c.Jump(p.Jump())
		} else if p.CommandType() == common.L_COMMAND {
			// L指令 伪指令
			fmt.Printf(
				"l_commond : line / %s, symbol / %s \n",
				p.Commands[p.CurrentReadLine],
				p.Symbol(),
			)
		}

		if tmpLine != "" {
			binLines = append(binLines, tmpLine)
		}

		if p.HasMoreCommands() {
			p.Advance()
		} else {
			break
		}
	}

	for _, v := range binLines {
		fmt.Printf("%s\n", v)
	}

	os.Exit(0)
}

// 格式化指令
func clearLine(line string) string {
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\n", "", -1)

	i := strings.Index(line, "//")
	if i == 0 {
		return ""
	}

	if i > 0 {
		line = line[0:i]
	}

	return line
}

// int 转二进制形式的字符串
func intToBinStr(n int) string {
	i := 0
	tmpComp := 1

	rstStr := ""

	for true {
		if tmpComp&n > 0 {
			rstStr = "1" + rstStr
		} else {
			rstStr = "0" + rstStr
		}

		i = i + 1
		tmpComp = tmpComp * 2

		if i == 15 {
			break
		}
	}
	return "0" + rstStr
}
