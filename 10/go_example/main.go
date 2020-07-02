package main

import (
	"go_example/CompilationEngine"
	"go_example/JackTokenizer"
	"os"
)

func main() {
	path := "."
	files := getFiles(path)
	for _, file := range files {
		fihandle, _ := os.Open(file)
		fohandle, _ := os.Open(jacktoxml(file))
		compileEngine := CompilationEngine.NewCompilationEngine(
			JackTokenizer.NewJackTokenizer(fihandle),
			fohandle,
		)
		compileEngine.CompileClass()
		fihandle.Close()
		fohandle.Close()
	}
}

func getFiles(path string) []string {
	return []string{}
}

func jacktoxml(filename string) string {
	return ""
}
