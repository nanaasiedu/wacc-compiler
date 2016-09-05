package ast

import (
	"errors"
	"fmt"
)

func (value CallInstance) Eval(context *Context) (Type, error) {
	return nil, nil
}

func (value Class) Eval(context *Context) (Type, error) {
	for _, field := range value.FieldList {
		switch field.FieldType.(type) {
		case ClassType:
			for _, classname := range context.ClassTable {
				if field.FieldType == classname.Ident {
					return field.FieldType, nil
				}
			}
			return nil, errorFieldUndefined(value.FileText, value.Pos)
		case ConstType:
			return field.FieldType, nil
		case ArrayType:
			return field.FieldType, nil
		case PairType:
			return field.FieldType, nil
		}
	}
	return nil,  errorFieldUndefined(value.FileText, value.Pos)
}

func (value ThisInstance) Eval(context *Context) (Type, error) {

	if context.SymbolTable.isDefined(Ident("this")) {
		classType := context.SymbolTable.getTypeOfIdent(Ident("this"))
		for _, class := range context.ClassTable {
			if classType == class.Ident {
				for _, field := range class.FieldList {
					if field.Ident == value.Field {
						return field.FieldType, nil
					}
				}
			}
		}
	}
	return nil, errorFieldUndefined(value.FileText, value.Pos)
}

func (value NewObject) Eval(context *Context) (Type, error) {
	classType := value.Class
	for _, class := range context.ClassTable {
		if classType == class.Ident {
			for ind, ev := range value.Init {
				t, err := ev.Eval(context)
				if err != nil {
					return nil, err
				}
				if t != class.FieldList[ind].FieldType {
					errorMisMatchingTypeInConstructor(value.FileText, value.Pos)
				}
			}
		}
	}
	return classType,nil
}
func (value FieldAccess) Eval(context *Context) (Type, error) {
	classType, err := value.ObjectName.Eval(context)
	if err != nil {
		return nil, errorFieldUndefined(value.FileText, value.Pos)
	}
	for _, k := range context.ClassTable {
		if k.Ident == classType {
			for _, a := range k.FieldList {
				if a.Ident == value.Field {
					return a.FieldType, nil
				}
			}
		}
	}
	return nil, errorFieldUndefined(value.FileText, value.Pos)
}

func (value Call) Eval(context *Context) (Type, error) {
	for _, function := range context.FunctionTable {
		if value.Ident == function.Ident {
			if len(value.ParamList) != len(function.ParameterTypes) {
				return nil, errorCallParam(value.FileText, value.Pos)
			}
			for ind := range value.ParamList {
				exprType, err := value.ParamList[ind].Eval(context)
				if err != nil {
					return nil, err
				}
				if exprType != function.ParameterTypes[ind].ParamType {
					return nil, errorCallParam(value.FileText, value.Pos)
				}
			}
			return function.ReturnType, nil
		}
	}
	return nil, errorNoFunction(value.FileText, value.Pos)
}

func (value ArrayLiter) Eval(context *Context) (Type, error) {
	var currType Type
	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].Eval(context)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.Eval(context)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, errorBadArrayLiter(value.FileText, value.Pos)
			}
		}
		return ArrayType{Type: currType}, nil
	}
	return nil, nil
}

func (value NewPair) Eval(context *Context) (Type, error) {
	fstTyp, err1 := value.FstExpr.Eval(context)
	sndTyp, err2 := value.SndExpr.Eval(context)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	return PairType{FstType: fstTyp, SndType: sndTyp}, nil
}

func (value PairElem) Eval(context *Context) (Type, error) {
	fstsnd := value.Fsnd
	exprTyp, err := value.Expr.Eval(context)
	if err != nil {
		return nil, err
	}
	switch exprTyp.(type) {
	case PairType:
		switch fstsnd {
		case Fst:
			return exprTyp.(PairType).FstType, nil
		case Snd:
			return exprTyp.(PairType).SndType, nil
		}
	}
	return nil, errorBadElemPair(value.FileText, value.Pos)
}

func (value Integer) Eval(context *Context) (Type, error) {
	return Int, nil
}

func (value PairLiter) Eval(context *Context) (Type, error) {
	return Pair, nil
}

func (value Str) Eval(context *Context) (Type, error) {
	return String, nil
}

func (value Character) Eval(context *Context) (Type, error) {
	return Char, nil
}

func (value Boolean) Eval(context *Context) (Type, error) {
	return Bool, nil
}

func (value ArrayElem) Eval(context *Context) (Type, error) {
	var currType Type

	if len(value.Exprs) > 0 {
		fstType, err := value.Exprs[0].Eval(context)
		currType = fstType
		if err != nil {
			return nil, err
		}
		for _, exprs := range value.Exprs {
			currType2, err2 := exprs.Eval(context)
			if err2 != nil {
				return nil, err2
			}
			if currType2 != currType {
				return nil, errorBadArrayLiter(value.FileText, value.Pos)
			}
		}
	}
	if currType != Int {
		return nil, errorArrayAccessExpr(value.FileText, value.Pos)
	}
	if !context.SymbolTable.isDefined(value.Ident) {
		return nil, errorArrayNotDefined(value.FileText, value.Pos)
	}
	arrayTyp := context.SymbolTable.getTypeOfIdent(value.Ident)
	for _ = range value.Exprs {
		switch arrayTyp.(type) {
		case ArrayType:
			arrayTyp = (arrayTyp.(ArrayType)).Type
		case ConstType:
			if arrayTyp.(ConstType) == String && len(value.Exprs) == 1 {
				return Char, nil
			}
		default:
			return nil, errorArrayTooMuchNesting(value.FileText, value.Pos)
		}
	}
	return arrayTyp, nil
}

// Recursive function
// Base case : No more dots
// Keep recursing while there are dots

// Get ident before first dot, look up in symboltable and get type (should be of type ClassType and in the []*Class)
// get next ident, check if Class/field/method is defined in class
func (value Ident) Eval(context *Context) (Type, error) {
	/*	valueString := string(value)
		if strings.Contains(valueString, ".") {
			index := strings.Index(valueString, ".")
			///item := valueString[:index]
			rest := valueString[index:]

			if context.symbolTable.isDefined(Ident(item)) {
				resType, err := Ident(rest).Eval(context)

				//Check if class type
				switch resType.(type) {
				case ClassType:

				}
			}

		} else { */
	if context.SymbolTable.isDefined(value) {
		return context.SymbolTable.getTypeOfIdent(value), nil
		//		}
	}
	return nil, errors.New(" :Cannot find " + string(value) + "in symbol table") // AWWHHHW
}

func (binop Binop) Eval(context *Context) (Type, error) {
	typl, err := binop.Left.Eval(context)
	typr, err2 := binop.Right.Eval(context)
	if err != nil {
		return nil, err
	}
	if err2 != nil {
		return nil, err2
	}
	switch binop.Binary {
	case PLUS, SUB, MUL, DIV, MOD:
		if typl != Int {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		if typr != Int {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		return Int, nil
	case AND, OR:
		if typl != Bool {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		if typr != Bool {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		return Bool, nil
	case LT, LTE, GT, GTE:
		if typl != Int && typl != Char {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		if typr != Int && typr != Char {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		if !typesMatch(typl, typr) {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		return Bool, nil
	case EQ, NEQ:
		if !typesMatch(typl, typr) {
			return nil, errorBadBinary(binop.FileText, binop.Pos)
		}
		return Bool, nil
	default:
		return nil, errorBadBinary(binop.FileText, binop.Pos)
	}
}

func (value Unop) Eval(context *Context) (Type, error) {
	typExpr, err := value.Expr.Eval(context)
	if err != nil {
		return nil, err
	}
	switch value.Unary {
	case NOT:
		if typExpr != Bool {
			return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Cannot negate a non Bool expression")
		}
		return Bool, nil
	case SUB:
		if typExpr != Int {
			return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Cannot sub a non Int expression")
		}
		return Int, nil
	case LEN:
		switch typExpr.(type) {
		case ArrayType:
			return Int, nil
		}
		return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Cannot perform len on non Array expression")
	case ORD:
		if typExpr != Char {
			return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Cannot perform ord on non Char expression")
		}
		return Int, nil
	case CHR:
		if typExpr != Int {
			return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Cannot perform len on non Int expression")
		}
		return Char, nil
	default:
		return nil, errors.New("line: " + fmt.Sprint(value.Pos) + " :Unary operation not recognised")
	}
}

func typesMatch(type1 Type, type2 Type) bool {
	switch type1.(type) {
	case PairType:
		if type2 == Pair {
			return true
		}
	}
	switch type2.(type) {
	case PairType:
		if type1 == Pair {
			return true
		}
	}
	return type1 == type2
}
