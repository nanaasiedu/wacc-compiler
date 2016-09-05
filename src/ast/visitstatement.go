package ast

func containsDuplicateFunc(functionTable []*Function) bool {
	freqMap := make(map[Ident]int)
	for _, funcDec := range functionTable {
		freqMap[funcDec.Ident]++
	}
	for _, val := range freqMap {
		if val > 1 {
			return true
		}
	}
	return false
}

func (node If) checkIfReturn(context *Context, returnType Type) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		semanticErrors = append(semanticErrors, errorConditional(node.FileText, node.Pos))
	}
	thenSymTab := context.SymbolTable.New()
	context.SymbolTable.Children = append(context.SymbolTable.Children, thenSymTab)
	for _, thenstat := range node.ThenStat {
		switch thenstat.(type) {
		case Return:
			errs := thenstat.(Return).checkReturnReturn(context, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		case If:
			errs := thenstat.(If).checkIfReturn(context, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		default:
			errs := thenstat.visitStatement(context)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		}
	}
	elseSymTab := context.SymbolTable.New()
	context.SymbolTable.Children = append(context.SymbolTable.Children, elseSymTab)
	for _, elsestat := range node.ElseStat {
		switch elsestat.(type) {
		case Return:
			errs := elsestat.(Return).checkReturnReturn(context, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		case If:
			errs := elsestat.(If).checkIfReturn(context, returnType)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		default:
			errs := elsestat.visitStatement(context)
			if errs != nil {
				semanticErrors = append(semanticErrors, errs)
			}
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Return) checkReturnReturn(context *Context, returnType Type) errorSlice {
	var semanticErrors []error
	exprTyp, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	} else {
		if !typesMatch(exprTyp, returnType) {
			semanticErrors = append(semanticErrors, errorReturn(node.FileText, node.Pos))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}
func (function *Function) checkFunc(root *Program) errorSlice {
	var semanticErrors []error
	funcSymbolTable := function.SymbolTable
	context := &Context{root.FunctionList, funcSymbolTable, root.ClassList}
	for _, param := range function.ParameterTypes {
		funcSymbolTable.insert(param.Ident, param.ParamType)
	}
	for ind, stat := range function.StatList {
		if ind == len(function.StatList)-1 {
			switch stat.(type) {
			case If:
				errIf := stat.(If).checkIfReturn(context, function.ReturnType)
				if errIf != nil {
					semanticErrors = append(semanticErrors, errIf)
				}
			case Return:
				errRet := stat.(Return).checkReturnReturn(context, function.ReturnType)
				if errRet != nil {
					semanticErrors = append(semanticErrors, errRet)
				}
			}
		}
		errs := stat.visitStatement(context)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {

		return semanticErrors
	}

	return nil
}

func (root *Program) SemanticCheck() errorSlice {
	var semanticErrors []error
	context := &Context{root.FunctionList, root.SymbolTable, root.ClassList }
	if containsDuplicateFunc(root.FunctionList) {
		semanticErrors = append(semanticErrors, errorFuncRedef(root.FileText, root.Pos))
	}

	for _, class := range root.ClassList {
    classymbtab := root.SymbolTable.New()
		classymbtab.insert(Ident("this"), class.Ident)
		classContext := &Context{root.FunctionList, classymbtab , root.ClassList }
  	_, err :=	class.Eval(classContext)
  	if err != nil {
			semanticErrors = append(semanticErrors, errorClassError(root.FileText, root.Pos))
  	}

	}
	for _, functionProg := range root.FunctionList {
		funcErrs := functionProg.checkFunc(root)
		if funcErrs != nil {
			semanticErrors = append(semanticErrors, funcErrs)
		}
	}
	for _, stat := range root.StatList {
		errs := stat.visitStatement(context)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
		switch stat.(type) {
		case Return:
			semanticErrors = append(semanticErrors, errorReturnInMain(root.FileText, root.Pos))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Skip) visitStatement(context *Context) errorSlice {
	return nil
}

// If called function not defined : ERROR
// other call symantic checks
// return value does not need to be caught
func (node Call) visitStatement(context *Context) errorSlice {
	return nil
}

//So you need to remember all classes
//So when calling a class method you have to check if:
//The first <ident> (object variable) is of type ClassType (using a symbol table i presume.
//Then using a list of classes and the <class ident> inside the ClassType. You need to check if the class actually has this method
//Then you need to check if the correct number of arguments and their type are correct​ This logic follows very closely to how class semantics is done
func (node CallInstance) visitStatement(context *Context) errorSlice {
	return nil
}

func (node ThisInstance) visitStatement(context *Context) errorSlice {
	return nil
}

func (node Instance) visitStatement(context *Context) errorSlice {
	return nil
}

func (node NewObject) visitStatement(context *Context) errorSlice {
	return nil
}

func (node Declare) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	if context.SymbolTable.isDefinedInScope(node.Lhs) {
		semanticErrors = append(semanticErrors, errorDeclared(node.FileText, node.Pos))
	}
	exprType, errs := node.Rhs.Eval(context)
	if errs != nil {
		semanticErrors = append(semanticErrors, errs)
	}
	if exprType != node.DecType && exprType != nil {
		switch node.DecType.(type) {
		case PairType:
			pairTypeStruct := node.DecType.(PairType)
			if pairTypeStruct.FstType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				case ConstType:
					if exprType.(ConstType) != Pair {
						semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
					}
				default:
					semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
				}
			}
			if pairTypeStruct.SndType == Pair {
				switch exprType.(type) {
				case PairType:
					// Do nothing
				case ConstType:
					if exprType.(ConstType) != Pair {
						semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
					}
				default:
					semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
				}
			}
		default:
			semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
		}
	}
	context.SymbolTable.insert(node.Lhs, node.DecType)
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Assignment) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	lhsType, errl := node.Lhs.Eval(context)
	rhsType, errr := node.Rhs.Eval(context)
	if errl != nil {
		semanticErrors = append(semanticErrors, errl)
	}
	if errr != nil {
		semanticErrors = append(semanticErrors, errr)
	}
	if rhsType == nil && errl == nil && errr == nil {
		switch lhsType.(type) {
		case ArrayType:
			// Do nothing
		default:
			semanticErrors = append(semanticErrors, errorArray(node.FileText, node.Pos))
		}
	}
	if !typesMatch(lhsType, rhsType) && rhsType != nil && lhsType != nil {
		semanticErrors = append(semanticErrors, errorTypeMatch(node.FileText, node.Pos))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Read) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.AssignLHS.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	} else if exprTyp != Char && exprTyp != Int {
		semanticErrors = append(semanticErrors, errorRead(node.FileText, node.Pos))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Free) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	switch exprTyp.(type) {
	case PairType:
		// Do nothing
	default:
		if exprTyp != Pair {
			semanticErrors = append(semanticErrors, errorPair(node.FileText, node.Pos))
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Return) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}

	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Exit) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	exprTyp, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if exprTyp != Int {
		semanticErrors = append(semanticErrors, errorExit(node.FileText, node.Pos))
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Print) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Println) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	_, err := node.Expr.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node If) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		//		semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Conditional is not boolean expression"))
		semanticErrors = append(semanticErrors, errorConditional(node.FileText, node.Pos))
	}
	thenSymTab := context.SymbolTable.New()
	thenContext := &Context{context.FunctionTable, thenSymTab, context.ClassTable}
	context.SymbolTable.Children = append(context.SymbolTable.Children, thenSymTab)
	for _, thenstat := range node.ThenStat {
		errs := thenstat.visitStatement(thenContext)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	elseSymTab := context.SymbolTable.New()
	elseContext := &Context{context.FunctionTable, elseSymTab, context.ClassTable}
	context.SymbolTable.Children = append(context.SymbolTable.Children, elseSymTab)
	for _, elsestat := range node.ElseStat {
		errs := elsestat.visitStatement(elseContext)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node While) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	cond, err := node.Conditional.Eval(context)
	if err != nil {
		semanticErrors = append(semanticErrors, err)
	}
	if cond != Bool {
		//			semanticErrors = append(semanticErrors, errors.New("line:"+fmt.Sprint(node.Pos)+" :Conditional is not boolean expression"))
		semanticErrors = append(semanticErrors, errorConditional(node.FileText, node.Pos))
	}
	whileSymTab := context.SymbolTable.New()
	context.SymbolTable.Children = append(context.SymbolTable.Children, whileSymTab)
	newContext := &Context{context.FunctionTable, whileSymTab, context.ClassTable}
	for _, dostat := range node.DoStat {
		errs := dostat.visitStatement(newContext)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}

func (node Scope) visitStatement(context *Context) errorSlice {
	var semanticErrors errorSlice
	newSymTab := context.SymbolTable.New()
	context.SymbolTable.Children = append(context.SymbolTable.Children, newSymTab)
	newContext := &Context{context.FunctionTable, newSymTab, context.ClassTable}
	for _, stat := range node.StatList {
		errs := stat.visitStatement(newContext)
		if errs != nil {
			semanticErrors = append(semanticErrors, errs)
		}
	}
	if len(semanticErrors) > 0 {
		return semanticErrors
	}
	return nil
}
