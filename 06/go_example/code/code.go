package code

import (
	"fmt"
	"os"
	"strings"
)

// Code Code 模块
type Code struct {
	JumpTable map[string]string
	CompTable map[string]string
}

// NewCode ...
func NewCode() *Code {
	return &Code{
		JumpTable: map[string]string{
			"JGT": "001",
			"JEQ": "010",
			"JGE": "011",
			"JLT": "100",
			"JNE": "101",
			"JLE": "110",
			"JMP": "111",
		},
		CompTable: map[string]string{
			"0":   "0101010",
			"1":   "0111111",
			"-1":  "0111010",
			"D":   "0001100",
			"A":   "0110000",
			"M":   "1110000",
			"!D":  "0001101",
			"!A":  "0110001",
			"!M":  "1110001",
			"-D":  "0001111",
			"-A":  "0110011",
			"-M":  "1110011",
			"D+1": "0011111",
			"A+1": "0110111",
			"M+1": "1110111",
			"D-1": "0001110",
			"A-1": "0110010",
			"M-1": "1110010",
			"D+A": "0000010",
			"D+M": "1000010",
			"D-A": "0010011",
			"D-M": "1010011",
			"A-D": "0000111",
			"M-D": "1000111",
			"D&A": "0000000",
			"D&M": "1000000",
			"D|A": "0010101",
			"D|M": "1010101",
		},
	}
}

// Dest 返回dest助记符的二进制码
func (c *Code) Dest(str string) string {
	rst := ""
	if strings.Index(str, "A") >= 0 {
		rst += "1"
	} else {
		rst += "0"
	}

	if strings.Index(str, "D") >= 0 {
		rst += "1"
	} else {
		rst += "0"
	}

	if strings.Index(str, "M") >= 0 {
		rst += "1"
	} else {
		rst += "0"
	}

	return rst
}

// Comp 返回comp助记符的二进制码
func (c *Code) Comp(str string) string {
	items := strings.Split(str, "+")
	if len(items) == 1 {
		val, valExist := c.CompTable[items[0]]
		if !valExist {
			fmt.Println("c1")
			os.Exit(-1)
			return ""
		}
		return val
	}

	if len(items) > 2 {
		fmt.Println("c2")
		os.Exit(-1)
		return ""
	}

	val1, val1Exist := c.CompTable[items[0]+"+"+items[1]]
	if val1Exist {
		return val1
	}

	val2, val2Exist := c.CompTable[items[1]+"+"+items[0]]
	if val2Exist {
		return val2
	}

	fmt.Println("c3")
	os.Exit(-1)
	return ""
}

// Jump 返回jump助记符的二进制码
func (c *Code) Jump(str string) string {
	val, valExist := c.JumpTable[str]
	if !valExist {
		val = "000"
	}
	return val
}
