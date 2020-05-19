package main

import (
	"bufio"
	"fmt"
	"go_example/codeWriter"
	"go_example/parser"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileOrPath := os.Args[1]
	filePaths := getFilePaths(fileOrPath)
	cw := codeWriter.NewCodeWriter("main.asm")

	for _, filePath := range filePaths {
		fileName := getNameFromFile(filePath)

		if !isVMFile(filePath) {
			continue
		}

		fileVMCommands := getVMCommands(filePath)
		if len(fileVMCommands) == 0 {
			continue
		}

		// 打印每个文件的
		fmt.Println("filename: " + filePath)
		fmt.Println(fileVMCommands)
		fmt.Println()

		p := parser.NewParser(filePath, fileVMCommands)
		cw.SetFileName(fileName)

		for true {
			if !p.HasMoreCommands() {
				break
			}

			p.Advance()

			if p.CommandType() == parser.C_ARITHMETIC {
				cw.WriteArithmetic(p.CurrentCommand)
			} else if p.CommandType() == parser.C_PUSH || p.CommandType() == parser.C_POP {
				index, indexError := strconv.Atoi(p.Arg2())
				if indexError != nil {
					println("error: m2")
					os.Exit(-1)
				}
				cw.WritePushPop(p.CommandType(), p.Arg1(), index)
			} else {
				println("error: m3")
				os.Exit(-1)
			}
		}
		println("")
		println("")
	}
	cw.Close()

	fmt.Println("vm translator")
}

// 从单个文件或者目录获取文件路径列表
func getFilePaths(fileOrPath string) []string {
	fop, e := os.Stat(fileOrPath)
	if e != nil {
		return []string{}
	}

	if fop.IsDir() {
		finfos, finfoError := ioutil.ReadDir(fileOrPath)
		if finfoError != nil {
			return []string{}
		}

		fileNames := []string{}
		for _, finfo := range finfos {
			fileNames = append(fileNames, finfo.Name())
		}
		return fileNames
	}

	return []string{fop.Name()}
}

// 从文件路径获取文件名
func getNameFromFile(file string) string {
	nodes := strings.Split(file, ".")
	return nodes[0]
}

// 是否是vm文件
func isVMFile(file string) bool {
	nodes := strings.Split(file, ".")
	if len(nodes) != 2 {
		return false
	}

	if nodes[1] != "vm" && nodes[1] != "VM" {
		return false
	}

	return true
}

// 从vm文件获取vm指令
func getVMCommands(file string) (commands []string) {
	f, ferr := os.Open(file)
	if ferr != nil {
		fmt.Println("error: m1")
		os.Exit(-1)
	}

	br := bufio.NewReader(f)
	for true {
		lineBytes, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lineStr := string(lineBytes)
		// 清除注释内容
		if i := strings.Index(lineStr, "//"); i >= 0 {
			lineStr = lineStr[0:i]
		}
		// 删除两侧空格
		lineStr = strings.TrimLeft(lineStr, " ")
		lineStr = strings.TrimRight(lineStr, " ")
		// 中间多空格替换为单空格
		spaceRep, _ := regexp.Compile("\\s+")
		lineStr = string(spaceRep.ReplaceAll([]byte(lineStr), []byte(" ")))
		if lineStr == "" {
			continue
		}
		commands = append(commands, lineStr)
	}
	return commands
}
