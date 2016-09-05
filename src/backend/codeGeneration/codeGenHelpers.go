package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"strconv"
)

var zeroInt Integer = 0
var zeroCharater Character = "0"
var zeroString Str = "0"
var zeroBool Boolean = false

// Contains the size in bytes of all print format strings
var mapPrintFormatToSize = map[string]int{
	INT_FORMAT:     3,
	STRING_FORMAT:  5,
	NEWLINE_MSG:    1,
	TRUE_MSG:       5,
	FALSE_MSG:      6,
	POINTER_FORMAT: 3,
}

//Size in bytes for all the variables in the current scope
func getScopeVarSize(statList []Statement) int {
	var scopeSize = 0
	for _, stat := range statList {
		switch stat.(type) {
		case Declare:
			scopeSize += sizeOf(stat.(Declare).DecType)
		}
	}
	return scopeSize
}

// Calcuates the size of a type
func sizeOf(t Type) int {
	var size = 0
	switch t.(type) {
	case ConstType:
		switch t.(ConstType) {
		case Int:
			size = INT_SIZE
		case Bool:
			size = BOOL_SIZE
		case Char:
			size = CHAR_SIZE
		case String:
			size = ADDRESS_SIZE
		case Pair:
			size = PAIR_SIZE
		}
	default: //PairType + ArrayType
		size = ADDRESS_SIZE
		/*	case PairType:
				size = ADDRESS_SIZE
			case ArrayType:
				size = ADDRESS_SIZE*/
	}
	return size
}

type Generic interface{}

func instructionSelection(is InstructionSelection) {
	appendAssembly(is.instrs, is.code, is.numTabs, is.numNewLines)
}

// Adds the string code to the list of instructions instrs.
// numTabs  \t will be added before the string and
// numNewLines \n will be added after the string
func appendAssembly(instrs *ARMList, code string, numTabs int, numNewLines int) {
	const default_num_tabs = 1
	var str string = ""

	for i := 0; i < numTabs+default_num_tabs; i++ {
		str += "\t"
	}

	str += code

	for i := 0; i < numNewLines; i++ {
		str += "\n"
	}

	*instrs = append(*instrs, str)
}

// Adds a msg label definition for the strValue using the msg label to the
// list of assembly instructions msgInstrs
func addMsgLabel(msgInstrs *ARMList, label string, strValue string) {
	if len(*msgInstrs) == 0 {
		appendAssembly(msgInstrs, ".data", 0, 2)
	}

	appendAssembly(msgInstrs, label+":", 0, 1)

	// size of strValue in bytes
	wordSize := calculateWordSize(strValue)

	appendAssembly(msgInstrs, ".word "+strconv.Itoa(wordSize), 1, 1)
	appendAssembly(msgInstrs, ".ascii "+strValue, 1, 2)
}

// Returns true iff the string ident as been previously assigned a label
func strIdentPrevDefined(ident Ident, scope *scopeData) bool {
	if scope == nil {
		return false
	}

	_, inMap := scope.identMsgLabelMap[ident]
	if inMap {
		return true
	} else {
		return strIdentPrevDefined(ident, scope.parentScope)
	}
}

// Returns the label of a ident. if it cant be found within the current scope,
// the parent is searched
func findLabel(ident Ident, scope *scopeData) string {
	if scope == nil {
		return "ERROR LABEL NOT FOUND"
	}

	label, inMap := scope.identMsgLabelMap[ident]
	if inMap {
		return label
	} else {
		return findLabel(ident, scope.parentScope)
	}
}

// Calculates the size of strValue in bytes
func calculateWordSize(strValue string) int {
	size, contained := mapPrintFormatToSize[strValue]

	// if strValue is a format string then
	if contained {
		return size
	} else {
		quotation := 2
		var escapedChars int
		backSlashSeen := false

		for _, c := range strValue {
			if c == '\\' && !backSlashSeen {
				escapedChars++
				backSlashSeen = true
			} else {
				backSlashSeen = false
			}
		}

		return len(strValue) - escapedChars - quotation
	}
}

// Returns true iff ident is within in the paramList
func isParamInList(ident Ident, paramList *[]Param) bool {
	for _, param := range *paramList {
		if ident == param.Ident {
			return true
		}
	}
	return false
}

// Returns total variable space a function is currently using from a starting inner scope
func funcVarSize(scope *scopeData) int {
	if !scope.isFunc {
		return 0
	}
	return scope.size + funcVarSize(scope.parentScope)
}

// returns total size of a list of params
func paramListSize(paramList []Param) int {
	totalCount := 0
	for _, param := range paramList {
		totalCount += sizeOf(param.ParamType)
	}
	return totalCount
}

// Returns the negative offset of the parameter
func getParamOffset(ident Ident, paramList *[]Param) (int, Type) {
	accOffset := 0

	for _, param := range *paramList {
		if ident == param.Ident {
			return -accOffset, param.ParamType //- sizeOf(param.ParamType)
		}

		accOffset += sizeOf(param.ParamType)
	}

	return 0, Int // ERROR
}

// Returns "1" iff b = true. "0" otherwise
func boolInt(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}
