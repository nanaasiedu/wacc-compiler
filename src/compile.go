package main

import (
	. "ast"
	cg "backend/codeGeneration"
	fw "backend/filewriter"
	"fmt"
	"io/ioutil"
	"os"
	"parser"
	"path/filepath"
)

const SYNTAX_ERROR = 100
const SEMANTIC_ERROR = 200

func main() {
	armList := &fw.ARMList{}

	file := os.Args[1] // index 1 is file path
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(b)
	root, err := parser.ParseFile(file, s)
	if err != nil {
		os.Exit(SYNTAX_ERROR)
	}
	fmt.Println(root)

	errs := root.SemanticCheck()
	if errs != nil {
		for _, str := range errs {
			fmt.Println(str)
		}
		os.Exit(SEMANTIC_ERROR)
	}
	root.SymbolTable.PrintChildren()

	filename := filepath.Base(file)
	ext := filepath.Ext(filename)
	fileARM := filename[0:len(filename)-len(ext)] + ".s"

	codeGen := cg.ConstructCodeGenerator(root, armList, root.SymbolTable)
	codeGen.GenerateCode()
	for _, instr := range *armList {
		fmt.Print(instr)
	}
	armList.WriteToFile(fileARM)

}
