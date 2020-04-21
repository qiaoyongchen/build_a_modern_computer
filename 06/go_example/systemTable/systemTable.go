package systemTable

import "os"

// SystemTable 模块
type SystemTable struct {
	Tabs map[string]int
}

// Construct 创建空的符号表
func Construct() *SystemTable {
	return &SystemTable{
		Tabs: map[string]int{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"R0":     0,
			"R1":     1,
			"R2":     2,
			"R3":     3,
			"R4":     4,
			"R5":     5,
			"R6":     6,
			"R7":     7,
			"R8":     8,
			"R9":     9,
			"R10":    10,
			"R11":    11,
			"R12":    12,
			"R13":    13,
			"R14":    14,
			"R15":    15,
			"SCREEN": 16384,
			"KBD":    24576,
		},
	}
}

// AddEntry 将(symbol, address)对加入符号表
func (t *SystemTable) AddEntry(symbol string, address int) {
	t.Tabs[symbol] = address
}

// Contains 符号表是否包含了 symbol ?
func (t *SystemTable) Contains(symbol string) bool {
	_, exist := t.Tabs[symbol]
	return exist
}

// GetAddress 返回与 symbol 关联的地址
func (t *SystemTable) GetAddress(symbol string) int {
	val, exist := t.Tabs[symbol]
	if !exist {
		os.Exit(-1)
		return 0
	}
	return val
}
