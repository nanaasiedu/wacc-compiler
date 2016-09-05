package codeGeneration

import (
	. "ast"
	. "backend/filewriter"
	"fmt"
	"strconv"
)

// CodeGenerator: Using a constructor that provides an address of the place holder
// for the list of assembly instructions for the program (And the AST with the symbol table),
// the code generator can use the "GenerateCode" function to fill the assembly list
// instrs with the assembly produced by traversing the tree
type CodeGenerator struct {
	root              *Program          // Root of the AST
	instrs            *ARMList          // List of assembly instructions for the program
	msgInstrs         *ARMList          // List of assembly instructions to create msg labels
	symTable          *SymbolTable      // Used to map variable identifiers to their types
	funcSymTables     []*SymbolTable    // Used to map function variable indentifier to ther types
	classes           []*Class          // List of classes defined within the program
	functionList      []*Function       // List of functions defined within the program
	funcInstrs        *ARMList          // List of assembly instructions that define functions and their labels
	progFuncInstrs    *ARMList          // List of assembly instructions that define program generated functions e.g. p_print_string
	progFuncNames     *[]string         // List of program defined function names. Used to avoid program redefinitions
	globalStack       *scopeData        // Stack data for the global scope
	currStack         *scopeData        // Stack data for the current scope
	msgMap            map[string]string // Maps string values to their msg labels
	messages          []string          // A slice of all messages used within the program
	currentLabelIndex int               // The index of the current generic label (used in control flow functions)
	functionDefs      []Ident           // A list of function identifiers that have been defined
}

// Constructor for the code generator.
func ConstructCodeGenerator(cRoot *Program, cInstrs *ARMList, cSymTable *SymbolTable) CodeGenerator {
	cg := CodeGenerator{root: cRoot, instrs: cInstrs, msgInstrs: new(ARMList),
		funcInstrs: new(ARMList), progFuncInstrs: new(ARMList), progFuncNames: new([]string),
		symTable: cSymTable, globalStack: &scopeData{identMsgLabelMap: make(map[Ident]string)}}

	// The program starts off with the current scope as the global scope
	cg.currStack = cg.globalStack

	cg.msgMap = make(map[string]string)
	return cg
}

// Evaluates the evaluation using the code generator
func (cg *CodeGenerator) eval(e Evaluation) Type {
	var eType Type

	if cg.currStack.isFunc {
		context1 := &Context{cg.root.FunctionList, cg.getFuncSymTable(), cg.classes}
		eType, _ = e.Eval(context1)
		//		eType, _ = e.Eval(cg.root.FunctionList, cg.getFuncSymTable())
	} else {
		context2 := &Context{cg.root.FunctionList, cg.symTable, cg.classes}
		eType, _ = e.Eval(context2)
	}

	return eType
}

// Provides information about the stack in relation to a specific scope
type scopeData struct {
	currP            int              // the current position of the pointer to the stack
	size             int              // size of the variable stack scope space in bytes
	parentScope      *scopeData       // Address of the parent scope
	isFunc           bool             // true iff the scope data is used for a function scope
	paramList        *[]Param         // List of parameters for a function if the scope is a function scope
	isMethod         bool             // true iff the scope data is used for a class method scope
	fieldList        *[]Field         // List of fields for a class object if the scope is a class scope
	identMsgLabelMap map[Ident]string // Map of string idents within this scope to their message labels
	extraOffset      int              // Extra offset used when stack is used to store intermediate values
}

// Creates new scope data for a new scope.
func (cg *CodeGenerator) setNewScope(varSpaceSize int) {
	if varSpaceSize > 0 {
		appendAssembly(cg.currInstrs(), "SUB sp, sp, #"+strconv.Itoa(varSpaceSize), 1, 1)
	}

	newScope := &scopeData{}
	newScope.currP = varSpaceSize
	newScope.size = varSpaceSize
	newScope.parentScope = cg.currStack
	newScope.isFunc = cg.currStack.isFunc
	newScope.identMsgLabelMap = make(map[Ident]string)

	if newScope.isFunc {
		newScope.paramList = cg.currStack.paramList
	}

	cg.currStack = newScope
	if cg.currStack.isFunc {
		cg.funcSymTables[len(cg.funcSymTables)-1] = cg.funcSymTables[len(cg.funcSymTables)-1].GetFrontChild()

	} else {
		cg.symTable = cg.symTable.GetFrontChild()
	}
}

// returns current function symbol table
func (cg *CodeGenerator) getFuncSymTable() *SymbolTable {
	if len(cg.funcSymTables) == 0 {
		return nil
	}

	return cg.funcSymTables[len(cg.funcSymTables)-1]
}

// Creates new scope data for a new function scope. Sets isFunc to true which
// set the code generator into function mode (So statements evaluate for functions not main)
func (cg *CodeGenerator) setNewFuncScope(varSpaceSize int, paramList *[]Param, funcSymTable *SymbolTable) {
	newScope := &scopeData{}
	newScope.currP = varSpaceSize
	newScope.size = varSpaceSize
	newScope.parentScope = cg.currStack
	newScope.isFunc = true
	newScope.paramList = paramList
	newScope.identMsgLabelMap = make(map[Ident]string)

	cg.currStack = newScope

	cg.funcSymTables = append(cg.funcSymTables, funcSymTable)
}

// Creates new scope data for a new function scope. Sets isFunc to true which
// set the code generator into function mode (So statements evaluate for functions not main)
func (cg *CodeGenerator) setNewMethodScope(varSpaceSize int, paramList *[]Param, funcSymTable *SymbolTable, fieldList *[]Field) {
	newScope := &scopeData{}
	newScope.currP = varSpaceSize
	newScope.size = varSpaceSize
	newScope.parentScope = cg.currStack
	newScope.isFunc = true
	newScope.paramList = paramList
	newScope.identMsgLabelMap = make(map[Ident]string)
	newScope.isMethod = true
	newScope.fieldList = fieldList

	cg.currStack = newScope

	cg.funcSymTables = append(cg.funcSymTables, funcSymTable)
}

// Removes current scope and replaces it with the parent scope
func (cg *CodeGenerator) removeCurrScope() {
	// add sp, sp, #n to remove variable space
	if cg.currStack.size > 0 {
		appendAssembly(cg.currInstrs(), "ADD sp, sp, #"+strconv.Itoa(cg.currStack.size), 1, 1)
	}

	cg.currStack = cg.currStack.parentScope
	if cg.currStack.isFunc {
		cg.funcSymTables[len(cg.funcSymTables)-1] = cg.funcSymTables[len(cg.funcSymTables)-1].Parent
	} else {
		cg.symTable = cg.symTable.Parent
	}
	if cg.currStack.isFunc {
		if cg.getFuncSymTable() != nil {
			cg.getFuncSymTable().RemoveChild()
		}
	} else {
		if cg.symTable != nil {
			cg.symTable.RemoveChild()
		}
	}
}

// Removes current function scope
func (cg *CodeGenerator) removeFuncScope() {
	cg.currStack = cg.currStack.parentScope
	cg.funcSymTables = cg.funcSymTables[:len(cg.funcSymTables)]
}

// Used to add extra offset to the current scope when intermediate values are stored on the stack
func (cg *CodeGenerator) addExtraOffset(n int) {
	cg.currStack.extraOffset += n
}

// Used to sub extra offset to the current scope when intermediate values are stored on the stack
func (cg *CodeGenerator) subExtraOffset(n int) {
	cg.currStack.extraOffset -= n
}

// Returns cg.funcInstrs iff the current scope is a function scope. cg.instrs otherwise
func (cg *CodeGenerator) currInstrs() *ARMList {
	if cg.currStack.isFunc {
		return cg.funcInstrs
	} else {
		return cg.instrs
	}
}

// Returns cg.funcSymbolTable iff the current scope is a function scope. cg.symbolTable otherwise
func (cg *CodeGenerator) currSymTable() *SymbolTable {
	if cg.currStack.isFunc {
		return cg.getFuncSymTable()
	} else {
		return cg.symTable
	}
}

// Decreases current pointer to the stack by n
// Returns new currP as a string (Does not have to be used)
func (cg *CodeGenerator) subCurrP(n int) string {
	cg.currStack.currP = cg.currStack.currP - n
	return strconv.Itoa(cg.currStack.currP)
}

// Using the ARMList pointer provided in the constructor,
// this function will fill the slice with an array of assembly instructions
// based on the provided AST
func (cg *CodeGenerator) GenerateCode() {
	cg.cgVisitProgram(cg.root)
	cg.buildFullInstr()
}

// Using all the code generators ARMList, instrs is modified to include
// all ARMList instructions in the correct order
func (cg *CodeGenerator) buildFullInstr() {
	*cg.instrs = append(*cg.funcInstrs, (*cg.instrs)...)
	*cg.instrs = append(*cg.msgInstrs, (*cg.instrs)...)
	*cg.instrs = append(*cg.instrs, *cg.progFuncInstrs...)
}

// Returns a msg label value for the strValue using msgMap
// If strValue is not contained in the map then it will be added to the map
// with a new msg label value (which will be returned)
// e.g. =msg_0
func (cg *CodeGenerator) getMsgLabel(ident Ident, strValue string) string {
	/*msgLabel, contained := cg.msgMap[strValue]

	if contained {
		return "=" + msgLabel
	}

	cg.msgMap[strValue] = "msg_" + strconv.Itoa(len(cg.msgMap))
	addMsgLabel(cg.msgInstrs, cg.msgMap[strValue], strValue)
	*/

	// For string constants
	if ident == "" {
		for i, message := range cg.messages {
			if strValue == message {
				return "=msg_" + strconv.Itoa(i)
			}
		}
		newIndex := len(cg.messages)
		cg.messages = append(cg.messages, strValue)
		newLabel := "msg_" + strconv.Itoa(newIndex)
		addMsgLabel(cg.msgInstrs, newLabel, strValue)
		return "=" + newLabel
	}

	// If the ident hasnt been previously defined in a msg
	if !strIdentPrevDefined(ident, cg.currStack) {
		newIndex := len(cg.messages)
		cg.messages = append(cg.messages, strValue)
		newLabel := "msg_" + strconv.Itoa(newIndex)
		addMsgLabel(cg.msgInstrs, newLabel, strValue)

		cg.currStack.identMsgLabelMap[ident] = newLabel

		return "=" + newLabel
	}

	// Find the idents label using all scopes
	label := findLabel(ident, cg.currStack)
	return "=" + label
}

// Adds the function name to cg.progFuncNames iff it isnt already in the list
// Returns true iff funcName is already in the list
func (cg *CodeGenerator) AddCheckProgName(progName string) bool {
	for _, s := range *cg.progFuncNames {
		if s == progName {
			// if progName has already been defined return true
			return true
		}
	}

	// else add progName to the list
	*cg.progFuncNames = append(*cg.progFuncNames, progName)

	return false
}

// Using symbol tables, a offset to the sp is returned so the ident value can
// be executed
func (cg *CodeGenerator) getIdentOffset(ident Ident) (int, Type) {
	return cg.findIdentOffset(ident, cg.currSymTable(), cg.currStack, 0)
}

// Checks if the ident is in the given symbol table. If not the parents are searched
// The function assumes an offset will be found eventually (semantically correct)
func (cg *CodeGenerator) findIdentOffset(ident Ident, symTable *SymbolTable,
	scope *scopeData, accOffset int) (int, Type) {
	if symTable == nil {
		fmt.Println("ERROR: incorrect symbol table")
		return 0, Int
	}

	if scope.isFunc && isParamInList(ident, scope.paramList) {
		offset, typ := getParamOffset(ident, scope.paramList)
		offset = -offset
		return offset + scope.extraOffset + ADDRESS_SIZE + funcVarSize(scope), typ //-scope.currP
	}

	/*fmt.Println("Ident: ", ident, "  table: ", symTable, " accOffset: ", accOffset)
	fmt.Println("Scope: ", scope)
	fmt.Println("Defined?:", symTable.IsOffsetDefined(ident), "\n")*/
	if !symTable.IsOffsetDefined(ident) {
		return cg.findIdentOffset(ident, symTable.Parent, scope.parentScope, accOffset+scope.size+scope.extraOffset)
	}

	return symTable.GetOffset(string(ident)) + accOffset + scope.extraOffset, symTable.GetTypeOfIdent(ident)

}

func (cg *CodeGenerator) getNewLabel() string {
	newLabel := "L" + strconv.Itoa(cg.currentLabelIndex)
	cg.currentLabelIndex++
	return newLabel
}

func (cg *CodeGenerator) isFunctionDefined(ident Ident) bool {
	for _, def := range cg.functionDefs {
		if string(ident) == string(def) {
			return true
		}
	}
	return false
}

func (cg *CodeGenerator) addFunctionDef(ident Ident) {
	cg.functionDefs = append(cg.functionDefs, ident)
}

// Returns total variable size of a list of evaluations
func (cg *CodeGenerator) evalSize(es []Evaluation) int {
	total := 0
	var eType Type
	for _, e := range es {
		eType = cg.eval(e)
		total += sizeOf(eType)
	}
	return total
}

// Returns class with the given class identifier
func (cg *CodeGenerator) getClass(classIdent ClassType) *Class {
	for _, class := range cg.classes {
		if classIdent == class.Ident {
			return class
		}
	}
	return nil
}

// Returns the offset to a object on the stack. Can only be used when within a method scope
func (cg *CodeGenerator) getObjectOffset() int {
	return ADDRESS_SIZE + paramListSize(*cg.currStack.paramList)
}
