package ast

import (
  "errors"
  "fmt"
  "strings"
)

func errorFuncRedef(file *string, pos int) error {
//  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint("Program has function redefinitions")
  return errors.New(msg)
}

func errorReturnInMain(file *string, pos int) error {
//  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint("Program has function redefinitions")
  return errors.New(msg)
}

func errorConditional(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "invalid conditional:", getLine(file, pos))
  return errors.New(msg)
}

func errorClassError(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Error in Class:", getLine(file, pos))
  return errors.New(msg)
}

func errorMisMatchingTypeInConstructor(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Mismatching types in constructor:", getLine(file, pos))
  return errors.New(msg)
}

func errorExit(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "exit value must be type int:", getLine(file, pos))
  return errors.New(msg)
}

func errorPair(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Cannot free non pair type:", getLine(file, pos))
  return errors.New(msg)
}

func errorRead(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Cannot read non Char or Int type:", getLine(file, pos))
  return errors.New(msg)
}

func errorArray(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "LHS is not of type Array:", getLine(file, pos))
  return errors.New(msg)
}

func errorTypeMatch(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Types do not match:", getLine(file, pos))
  return errors.New(msg)
}

func errorDeclared(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Variable already declared:", getLine(file, pos))
  return errors.New(msg)
}

func errorReturn(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", "Return type does not match:", getLine(file, pos))
  return errors.New(msg)
}

func errorCallParam(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Parameters of call and defintion do not match", getLine(file, pos))
  return errors.New(msg)
}

func errorNoFunction(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :No such function defined", getLine(file, pos))
  return errors.New(msg)
}

func errorBadArrayLiter(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Array has mixed types", getLine(file, pos))
  return errors.New(msg)
}

func errorBadElemPair(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Cannot get elem of non pair type", getLine(file, pos))
  return errors.New(msg)
}

func errorArrayAccessExpr(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Array cannot have non int expr", getLine(file, pos))
  return errors.New(msg)
}

func errorArrayNotDefined(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Array not defined", getLine(file, pos))
  return errors.New(msg)
}

func errorArrayTooMuchNesting(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Array not defined", getLine(file, pos))
  return errors.New(msg)
}

func errorIdentNotDeclared(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Identifier not declared", getLine(file, pos))
  return errors.New(msg)
}

func errorBadBinary(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Invalid binary expression", getLine(file, pos))
  return errors.New(msg)
}

func errorFieldUndefined(file *string, pos int) error {
  line, col := LineAndCol(file, pos)
  msg := fmt.Sprint(line, ":", col,"  ", " :Field undefined", getLine(file, pos))
  return errors.New(msg)
}

//Returns line number and column number of current lexing item in .wacc file
func LineAndCol(file *string, pos int) (line int, col int) {
  str := *file
	line = 1 + strings.Count(str[:pos], "\n")
	col = pos - strings.LastIndex(str[:pos], "\n")
	return
}

//Returns line number and column number of current lexing item in .wacc file
func getLine(file *string, pos int) string {
  str := *file
  start := strings.LastIndex(str[:pos], "\n")
  end := pos + strings.Index(str[pos:], "\n")
  result :=  str[start+1:end]
  return result
}
