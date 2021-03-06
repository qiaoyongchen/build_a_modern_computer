package codeWriter

import (
	"fmt"
	"go_example/parser"
	"os"
	"strconv"
)

const (
	SEG_ARGUMENT = "argument"
	SEG_LOCAL    = "local"
	SEG_STATIC   = "static"
	SEG_CONSTANT = "constant"
	SEG_THIS     = "this"
	SEG_THAT     = "that"
	SEG_POINTER  = "pointer"
	SEG_TEMP     = "temp"
)

var segmentAllow = []string{
	SEG_ARGUMENT,
	SEG_LOCAL,
	SEG_STATIC,
	SEG_CONSTANT,
	SEG_THIS,
	SEG_THAT,
	SEG_POINTER,
	SEG_TEMP,
}

func isAllowSegment(segment string) bool {
	ok := false
	for _, v := range segmentAllow {
		if v == segment {
			ok = true
			break
		}
	}
	return ok
}

// CodeWriter 汇编源码输出器
type CodeWriter struct {
	outputFileNmae string
	outputFile     *os.File
	fileName       string
}

// NewCodeWriter 构造器
func NewCodeWriter(outputFileNmae string) *CodeWriter {
	f, ferror := os.OpenFile(outputFileNmae, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if ferror != nil {
		println("error: cw2")
		os.Exit(-1)
	}

	cw := &CodeWriter{
		outputFileNmae: outputFileNmae,
		outputFile:     f,
	}
	return cw
}

// SetFileName 通知代码写入程序新的VM文件翻译已开始
func (cw *CodeWriter) SetFileName(fileName string) {
	cw.fileName = fileName
	return
}

// eq/gt/lt 需要在代码中插入符号，判断后进行跳转
var eqORgtORltIndex = 0

// WriteArithmetic 将给定的算数操作所对应的汇编代码写至输出
// command 算数命令
func (cw *CodeWriter) WriteArithmetic(command string) {
	commandStrPopTwice := ""
	commandStrPopTwice += "@SP\r\n"    //A寄存器加载statck地址
	commandStrPopTwice += "AM=M-1\r\n" //sp减一
	commandStrPopTwice += "D=M\r\n"    //D寄存器指向减一后的寄存器
	commandStrPopTwice += "A=A-1\r\n"  //A寄存器再减一，指向原始sp减二处,此时M指向sp-2的值,D位于sp-1处

	commandStrPopOnce := ""
	commandStrPopOnce += "@SP\r\n"    //A寄存器加载statck地址
	commandStrPopOnce += "AM=M-1\r\n" //sp减一

	CommandEqORgtORltStr := ""
	CommandEqORgtORltStr += "@SP\r\n"          //A寄存器加载statck地址
	CommandEqORgtORltStr += "AM=M-1\r\n"       //sp减一
	CommandEqORgtORltStr += "D=M\r\n"          //D寄存器指向减一后的寄存器
	CommandEqORgtORltStr += "A=A-1\r\n"        //A寄存器再减一，指向原始sp减二处,此时M指向sp-2的值,D位于sp-1处
	CommandEqORgtORltStr += "D=M-D\r\n"        //计算相减,存放在D寄存器
	CommandEqORgtORltStr += "@TRUE%d\r\n"      //A指令指向TRUEx代码处准备跳转（如果符合条件的话）
	CommandEqORgtORltStr += "D;%s\r\n"         //判断d，符合条件跳转
	CommandEqORgtORltStr += "@SP\r\n"          //
	CommandEqORgtORltStr += "AM=M-1\r\n"       //sp - 2
	CommandEqORgtORltStr += "M=0\r\n"          //sp - 2 设置为false
	CommandEqORgtORltStr += "@SP\r\n"          //
	CommandEqORgtORltStr += "M=M+1\r\n"        //设置回 sp - 1 (sp - 1 - 1 + 1)
	CommandEqORgtORltStr += "@CONTINUE%d\r\n"  //
	CommandEqORgtORltStr += "0;JMP\r\n"        //
	CommandEqORgtORltStr += "(TRUE%d)\r\n"     //
	CommandEqORgtORltStr += "@SP\r\n"          //
	CommandEqORgtORltStr += "AM=M-1\r\n"       //
	CommandEqORgtORltStr += "M=-1\r\n"         //sp - 2 设置为true
	CommandEqORgtORltStr += "@SP\r\n"          //
	CommandEqORgtORltStr += "M=M+1\r\n"        //
	CommandEqORgtORltStr += "(CONTINUE%d)\r\n" //

	commandStr := ""

	switch command {
	case "add":
		commandStr = commandStrPopTwice + "M=M+D\r\n" //计算结果并把结果放在sp-2处，类似于弹出两个栈，计算结果后再压入栈
	case "sub":
		commandStr = commandStrPopTwice + "M=M-D\r\n" //计算结果并把结果放在sp-2处，类似于弹出两个栈，计算结果后再压入栈
	case "neg":
		commandStr = commandStrPopOnce + "M=-M\r\n" //sp-1处原地计算，类似于先弹出，再计算，再压入栈
		commandStr += "@SP\r\n"                     //sp 指针再次指向未分配的栈顶
		commandStr += "M=M+1\r\n"
	case "eq":
		eqORgtORltIndex++
		commandStr = fmt.Sprintf(CommandEqORgtORltStr, eqORgtORltIndex, "JEQ", eqORgtORltIndex, eqORgtORltIndex, eqORgtORltIndex)
	case "gt":
		eqORgtORltIndex++
		commandStr = fmt.Sprintf(CommandEqORgtORltStr, eqORgtORltIndex, "JGT", eqORgtORltIndex, eqORgtORltIndex, eqORgtORltIndex)
	case "lt":
		eqORgtORltIndex++
		commandStr = fmt.Sprintf(CommandEqORgtORltStr, eqORgtORltIndex, "JLT", eqORgtORltIndex, eqORgtORltIndex, eqORgtORltIndex)
	case "and":
		commandStr = commandStrPopTwice + "M=M&D\r\n"
	case "or":
		commandStr = commandStrPopTwice + "M=M|D\r\n"
	case "not":
		commandStr = commandStrPopOnce + "M=!M\r\n"
		commandStr += "@SP\r\n" //sp 指针再次指向未分配的栈顶
		commandStr += "M=M+1\r\n"
	}

	cw.outputFile.Write([]byte(commandStr))
}

// WritePushPop 将给定的 command (类型为 C_PUSH 或 C_POP)
// 所对应的汇编代码写至输出
func (cw *CodeWriter) WritePushPop(command int, segment string, index int) {

	commandStr := ""

	if !isAllowSegment(segment) {
		println("error: cw1")
		println("error segment: " + segment)
		os.Exit(-1)
	}

	// 内存段参考
	// RAM[0] -> SP -> 栈顶指针
	// RAM[1] -> LCL -> 指向当前函数的local段地址
	// RAM[2] -> ARG -> 指向当前函数的argment段地址
	// RAM[3] -> THIS -> 指向this段基址(堆中) / pointer[0]
	// RAM[4] -> THAT -> 指向that段基址(堆中) / pointer[1]
	// RAM[5-12] -> ---- -> 保存temp段内容
	// RAM[13-15] -> ---- -> 可被VM实现用做通用寄存器
	// RAM[16-255] -> ---- -> VM程序的所有VM函数的静态变量
	// RAM[256-2047] -> ---- -> 栈
	// RAM[2048-16383] -> ---- -> 堆(用于存放对象和数组)
	// RAM[16384-24575] -> ---- -> 内存映像IO

	switch segment {
	case SEG_CONSTANT: //contant - 包含所有常数的伪段
		// 常数只有push没有pop
		switch command {
		case parser.C_PUSH:
			commandStr += "@" + strconv.Itoa(index) + "\r\n" // A 寄存器存入常数
			commandStr += "D=A\r\n"                          // D 寄存器复制 A 寄存器的内容(即常数)
			commandStr += "@SP\r\n"                          // A 寄存器存入sp地址
			commandStr += "A=M\r\n"                          // A 寄存器存入栈指针
			commandStr += "M=D\r\n"                          // 栈指针处存入 D寄存器内容(即常数)
			commandStr += "@SP\r\n"                          // A 寄存器存入sp地址
			commandStr += "M=M+1\r\n"                        // sp 内容(即栈指针)加1

			println("--constant start:--")
			println("constant: " + segment + strconv.Itoa(index))
			println("--constant end:--")
			println("")
		case parser.C_POP:
			println("error: cw3")
			os.Exit(-1)
		}
	case SEG_STATIC: //static - 存储同一个vm文件内共享的静态变量
		staticName := cw.fileName + "." + strconv.Itoa(index)

		switch command {
		case parser.C_PUSH:
			commandStr += "@" + staticName + "\r\n" // A 寄存器存入常量地址
			commandStr += "D=M\r\n"                 // D 寄存器存入常量内容
			commandStr += "@SP\r\n"                 // A 寄存器存入sp地址
			commandStr += "A=M\r\n"                 // A 寄存器存入栈地址
			commandStr += "M=D\r\n"                 // 栈地址存入常量内容
			commandStr += "@SP\r\n"                 // A 寄存器存入sp地址
			commandStr += "M=M+1\r\n"               // 栈地址加1
		case parser.C_POP:
			commandStr += "@SP\r\n"    //
			commandStr += "AM=M-1\r\n" //
			commandStr += "D=M\r\n"    //
			commandStr += "@" + staticName + "\r\n"
			commandStr += "M=D\r\n"
		}
	case SEG_POINTER: //pointer - 该段存储两个两个内存单元: 0-this 1-that
		switch index {
		case 0:
			segment = "THIS"
		case 1:
			segment = "THAT"
		}

		switch command {
		case parser.C_PUSH:
			commandStr += "@" + segment + "\r\n" // A 寄存器存入segment地址
			commandStr += "D=M\r\n"              // D 寄存器存入segment内容
			commandStr += "@SP\r\n"              // A 寄存器存入sp地址
			commandStr += "A=M\r\n"              // A 寄存器存入栈地址
			commandStr += "M=D\r\n"              // 栈地址存入常量内容
			commandStr += "@SP\r\n"              // A 寄存器存入sp地址
			commandStr += "M=M+1\r\n"            // 栈地址加1
		case parser.C_POP:
			commandStr += "@SP\r\n"              //
			commandStr += "AM=M-1\r\n"           //
			commandStr += "D=M\r\n"              //
			commandStr += "@" + segment + "\r\n" //
			commandStr += "M=D\r\n"              //
		}
	case SEG_THIS: //this
		switch command {
		case parser.C_PUSH:
			commandStr += "@THIS\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "A=D+A\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "M=M+1\r\n"                        //
		case parser.C_POP:
			commandStr += "@THIS\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "D=D+A\r\n"                        //
			commandStr += "@R13\r\n"                         //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "AM=M-1\r\n"                       //
			commandStr += "D=M\r\n"                          //
			commandStr += "@R13\r\n"                         //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
		}
	case SEG_THAT: //that
		switch command {
		case parser.C_PUSH:
			commandStr += "@THAT\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "A=D+A\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "M=M+1\r\n"                        //
		case parser.C_POP:
			commandStr += "@THAT\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "D=D+A\r\n"                        //
			commandStr += "@R13\r\n"                         //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "AM=M-1\r\n"                       //
			commandStr += "D=M\r\n"                          //
			commandStr += "@R13\r\n"                         //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
		}
	case SEG_TEMP: //temp - 固定的段,由8个规定的内存单元组成,用来保存临时变量
		switch command {
		case parser.C_PUSH:
			commandStr += "@R5\r\n"                          //
			commandStr += "D=A\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "A=D+A\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "M=M+1\r\n"                        //
		case parser.C_POP:
			commandStr += "@R5\r\n"                          //
			commandStr += "D=A\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "D=D+A\r\n"                        //
			commandStr += "@R13\r\n"                         //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "AM=M-1\r\n"                       //
			commandStr += "D=M\r\n"                          //
			commandStr += "@R13\r\n"                         //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
		}
	case SEG_LOCAL: //local - 存储函数的局部变量
		switch command {
		case parser.C_PUSH:
			commandStr += "@LCL\r\n"                         //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "A=D+A\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "M=M+1\r\n"                        //
		case parser.C_POP:
			commandStr += "@LCL\r\n"                         //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "D=D+A\r\n"                        //
			commandStr += "@R13\r\n"                         //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "AM=M-1\r\n"                       //
			commandStr += "D=M\r\n"                          //
			commandStr += "@R13\r\n"                         //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
		}
	case SEG_ARGUMENT: //argment - 存储函数的参数
		switch command {
		case parser.C_PUSH:
			commandStr += "@ARG\r\n"                         //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "A=D+A\r\n"                        //
			commandStr += "D=M\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "M=M+1\r\n"                        //
		case parser.C_POP:
			commandStr += "@ARG\r\n"                         //
			commandStr += "D=M\r\n"                          //
			commandStr += "@" + strconv.Itoa(index) + "\r\n" //
			commandStr += "D=D+A\r\n"                        //
			commandStr += "@R13\r\n"                         //
			commandStr += "M=D\r\n"                          //
			commandStr += "@SP\r\n"                          //
			commandStr += "AM=M-1\r\n"                       //
			commandStr += "D=M\r\n"                          //
			commandStr += "@R13\r\n"                         //
			commandStr += "A=M\r\n"                          //
			commandStr += "M=D\r\n"                          //
		}
	}

	_, e := cw.outputFile.Write([]byte(commandStr))
	if e != nil {
		println(e.Error())
		os.Exit(-1)
	}
}

// WriteInit 编写VM初始化的汇编代码，也称为引导程序
// 该代码必须被置于输出文件的开头
// SP内容为256
func (cw *CodeWriter) WriteInit() {
	commandStr := ""
	commandStr += "@256\r\n"
	commandStr += "D=A\r\n"
	commandStr += "@SP\r\n"
	commandStr += "M=D\r\n"
	cw.outputFile.Write([]byte(commandStr))

	cw.WriteCall("Sys.init", 0)
}

// WriteLabel 编写执行 label 命令的汇编代码
func (cw *CodeWriter) WriteLabel(label string) {
	commandStr := "(" + label + ")\r\n"
	cw.outputFile.Write([]byte(commandStr))
}

// WriteGoto 编写执行 goto 命令的汇编代码
func (cw *CodeWriter) WriteGoto(label string) {
	commandStr := ""
	commandStr += "@" + label + "\r\n"
	commandStr += "0;JMP\r\n"
	cw.outputFile.Write([]byte(commandStr))
}

// WriteIf 编写执行 if 命令的汇编代码
func (cw *CodeWriter) WriteIf(label string) {
	commandStr := ""
	commandStr += "@SP\r\n"
	commandStr += "AM=M-1\r\n"
	commandStr += "D=M\r\n"
	commandStr += "@" + label + "\r\n"
	commandStr += "D;JNE\r\n"
	cw.outputFile.Write([]byte(commandStr))
}

var callIndex = -1

// WriteCall 编写执行 call 命令的汇编代码
func (cw *CodeWriter) WriteCall(functionName string, numArgs int) {
	callIndex++

	var returnAddressFunc = func() string {
		return functionName + "RETURN_ADDRESS" + strconv.Itoa(callIndex)
	}

	commandStr := ""
	// 首先把执行完的返回地址压入栈
	commandStr += "@" + returnAddressFunc() + "\r\n"
	commandStr += "D=A\r\n"
	commandStr += "@SP\r\n"
	commandStr += "A=M\r\n"
	commandStr += "M=D\r\n"
	commandStr += "@SP\r\n"
	commandStr += "M=M+1\r\n"

	// 再把四个寄存器压入栈
	bePushedPointers := []string{"LCL", "ARG", "THIS", "THAT"}
	for _, v := range bePushedPointers {
		commandStr += "@" + v + "\r\n"
		commandStr += "D=M\r\n"
		commandStr += "@SP\r\n"
		commandStr += "A=M\r\n"
		commandStr += "M=D\r\n"
		commandStr += "@SP\r\n"
		commandStr += "M=M+1\r\n"
	}

	// 查找栈，跳过前面设置的参数和寄存器和返回地址 即 参数数量 + 4 + 1
	// 那么查找到堆栈中的地址就是参数值在堆栈中的起始处(call 之前会push数据)
	// 调用function并执行
	commandStr += "@" + strconv.Itoa(numArgs) + "\r\n"
	commandStr += "D=A\r\n"
	commandStr += "@5\r\n"
	commandStr += "D=A+D\r\n"
	commandStr += "@SP\r\n"
	commandStr += "D=M-D\r\n"
	commandStr += "@ARG\r\n"
	commandStr += "M=D\r\n"
	commandStr += "@SP\r\n"
	commandStr += "D=M\r\n"
	commandStr += "@LCL\r\n"
	commandStr += "M=D\r\n"
	commandStr += "@" + functionName + "\r\n"
	commandStr += "0;JMP\r\n"
	commandStr += "(" + returnAddressFunc() + ")\r\n"

	cw.outputFile.Write([]byte(commandStr))
}

// WriteReturn 编写执行 return 命令的汇编代码
func (cw *CodeWriter) WriteReturn() {
	commandStr := ""
	commandStr += "@LCL\r\n"
	commandStr += "D=M\r\n"
	commandStr += "@R13\r\n"
	commandStr += "M=D\r\n" // 以上 LCL 保存在 R13中

	commandStr += "@5\r\n"
	commandStr += "A=D-A\r\n"
	commandStr += "D=M\r\n"
	commandStr += "@R14\r\n"
	commandStr += "M=D\r\n" // 以上 LCL - 5 取得的返回地址保存在 R14 中

	commandStr += "@SP\r\n"
	commandStr += "AM=M-1\r\n"
	commandStr += "D=M\r\n" // sp减一 获取最近塞入的值 及返回值 放入D寄存器中

	commandStr += "@ARG\r\n"
	commandStr += "A=M\r\n"
	commandStr += "M=D\r\n" // 当前ARG在堆栈中的位置放入返回值 待会ARG就会被改写成调用之前的ARG了，这个位置就没用了
	commandStr += "@ARG\r\n"
	commandStr += "D=M+1\r\n"
	commandStr += "@SP\r\n"
	commandStr += "M=D\r\n" // 堆栈指针加1指向新的位置

	// 恢复之前保存的四个寄存器里面的值 ps:和保存顺序相反
	for _, v := range []string{"THAT", "THIS", "ARG", "LCL"} {
		commandStr += "@R13\r\n"
		commandStr += "D=M-1\r\n"
		commandStr += "AM=D\r\n"
		commandStr += "D=M\r\n"
		commandStr += "@" + v + "\r\n"
		commandStr += "M=D\r\n"
	}

	commandStr += "@R14\r\n"
	commandStr += "A=M\r\n"
	commandStr += "0;JMP\r\n" // 以上 跳转到之前保存的指令地址 (调用之前的下一条指令)

	cw.outputFile.Write([]byte(commandStr))
}

// WriteFunction 编写执行 function 命令的汇编代码
func (cw *CodeWriter) WriteFunction(functionName string, numLocals int) {
	commandStr := ""
	commandStr += "(" + functionName + ")\r\n"
	cw.outputFile.Write([]byte(commandStr))

	for numLocals > 0 {
		numLocals--
		cw.WritePushPop(parser.C_PUSH, SEG_CONSTANT, 0)
	}
}

// Close 关闭输出
func (cw *CodeWriter) Close() {
	cw.outputFile.Close()
}
