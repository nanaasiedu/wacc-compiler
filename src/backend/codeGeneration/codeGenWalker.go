package codeGeneration

import (
	. "ast"
	"fmt"
	"strconv"
)

// CONSTANTS -------------------------------------------------------------------

// Type sizes in bytes
const (
	INT_SIZE     = 4
	ARRAY_SIZE   = 4
	BOOL_SIZE    = 1
	CHAR_SIZE    = 1
	STRING_SIZE  = 4
	PAIR_SIZE    = 4
	ADDRESS_SIZE = 4
	// Maximum offset of stack pointer that can be added or subtracted from SP
	STACK_SIZE_MAX = 1024
)

// Print format strings
const (
	INT_FORMAT     = "\"%d\\0\""
	STRING_FORMAT  = "\"%.*s\\0\""
	NEWLINE_MSG    = "\"\\0\""
	TRUE_MSG       = "\"true\\0\""
	FALSE_MSG      = "\"false\\0\""
	READ_INT       = "\"%d\\0\""
	READ_CHAR      = "\" %c\\0\""
	POINTER_FORMAT = "\"%p\\0\""
)

// error messages
const (
	NULL_REFERENCE       = "\"NullReferenceError: dereference a null reference\\n\\0\""
	ARRAY_INDEX_NEGATIVE = "\"ArrayIndexOutOfBoundsError: negative index\\n\\0\""
	ARRAY_INDEX_LARGE    = "\"ArrayIndexOutOfBoundsError: index too large\\n\\0\""
	OVERFLOW             = "\"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\\n\""
	DIVIDE_BY_ZERO       = "\"DivideByZeroError: divide or modulo by zero\\n\\0\""
)

// HELPER FUNCTIONS
// cgVisitReadStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitReadStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// Define the read function
		instructionSelection(InstructionSelection{cg.progFuncInstrs, funcName + ":", 0, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "MOV r1, r0", 1, 1})
		switch funcName {
		case "p_read_char":
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", READ_CHAR), 1, 1})
		case "p_read_int":
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", READ_INT), 1, 1})
		}
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "ADD r0, r0, #4", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL scanf", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
	}
}

// Null pointer dereference check
func (cg *CodeGenerator) dereferenceNullPointer() {
	if !cg.AddCheckProgName("p_check_null_pointer") {
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "p_check_null_pointer" + ":", 0, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r0, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDREQ r0, " + cg.getMsgLabel("", NULL_REFERENCE), 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BLEQ p_throw_runtime_error", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
	}
	cg.throwRunTimeError()
}

// Run time error check
func (cg *CodeGenerator) throwRunTimeError() {
	if !cg.AddCheckProgName("p_throw_runtime_error") {
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "p_throw_runtime_error" + ":", 0, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL p_print_string", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "MOV r0, #-1", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL exit", 1, 1})
		cg.cgVisitPrintStatFuncHelper("p_print_string")
	}
}

// Check array bounds check
func (cg *CodeGenerator) checkArrayBounds() {
	if !cg.AddCheckProgName("p_check_array_bounds") {
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "p_check_array_bounds" + ":", 0, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r0, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDRLT r0, " + cg.getMsgLabel("", ARRAY_INDEX_NEGATIVE), 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BLLT p_throw_runtime_error", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r1, [r1]", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r0, r1", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDRCS r0, " + cg.getMsgLabel("", ARRAY_INDEX_LARGE), 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BLCS p_throw_runtime_error", 1, 1})
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
	}
	cg.throwRunTimeError()
}

// Helper function to check the index of the array and remove duplication
func (cg CodeGenerator) arrayCheckIndex() {
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r5", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r1, r4", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_array_bounds", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r4, r4, #4", 1, 1})
}

// Helper function to remove code duplication (BETTER NAME?)
func (cg CodeGenerator) arrayCheckBoundsHelper(t Type) {
	switch t {
	case Bool, Char:
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r4, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDRSB r4, [r4]", 1, 1})
	default:
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r4, r4, r5, LSL #2", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [r4]", 1, 1})
	}
}

// Checks the bounds of an array
func (cg *CodeGenerator) arrayCheckBounds(array []Evaluation, ident Ident) {

	// Load the first index
	cg.evalRHS(array[0], "r5")

	// Set a register to point to the array
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [r4]", 1, 1})

	//Check the index of the array
	cg.arrayCheckIndex()

	// Point to the correct index
	var t = cg.eval(ident)
	switch t.(type) {
	case ArrayType:
		cg.arrayCheckBoundsHelper(t.(ArrayType).Type)
	case ConstType:
		cg.arrayCheckBoundsHelper(t.(ConstType))
	}

	// Load the second index if there is one
	if len(array) > 1 {

		// Load the second index
		cg.evalRHS(array[1], "r5")

		// Check the index of the array
		cg.arrayCheckIndex()

		// Point to the correct index
		cg.arrayCheckBoundsHelper(cg.eval(ident).(ArrayType).Type.(ArrayType).Type)

	}

	// Forgot what this MOV is for...
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})

}

// cgVisitFreeStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitFreeStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// Define the read function
		instructionSelection(InstructionSelection{cg.progFuncInstrs, funcName + ":", 0, 1})
		switch funcName {
		case "p_free_pair":
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r0, #0", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDREQ r0, " + cg.getMsgLabel("", NULL_REFERENCE), 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BEQ p_throw_runtime_error", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {r0}", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, [r0]", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL free", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, [sp]", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, [r0, #4]", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL free", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {r0}", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL free", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
		}
		cg.throwRunTimeError()
		cg.cgVisitPrintStatFuncHelper("p_print_string")
	}
}

// cgVisitPrintStat helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitPrintStatFuncHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// funcLabel:
		instructionSelection(InstructionSelection{cg.progFuncInstrs, funcName + ":", 0, 1})
		// push {lr} to save the caller address
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})

		switch funcName {
		case "p_print_int":
			// r1 = int value
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "MOV r1, r0", 1, 1})
			// r0 = int format string
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", INT_FORMAT), 1, 1})

		case "p_print_bool":
			// Check bool value - 0
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r0, #0", 1, 1})
			// If bool = true then r0 = "true"
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDRNE r0, " + cg.getMsgLabel("", TRUE_MSG), 1, 1})
			// If bool = false then r0 = "false"
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDREQ r0, " + cg.getMsgLabel("", FALSE_MSG), 1, 1})

		case "p_print_string":
			// r1 = string value
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r1, [r0]", 1, 1})
			// r2 = r0 + 4
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "ADD r2, r0, #4", 1, 1})
			// r0 = string format string
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", STRING_FORMAT), 1, 1})

		case "p_print_ln":
			// r0 = new line string
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", NEWLINE_MSG), 1, 1})

		case "p_print_reference":
			// r1 = int value
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "MOV r1, r0", 1, 1})
			// r0 = pointer format string
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", POINTER_FORMAT), 1, 1})
		}

		//
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "ADD r0, r0, #4", 1, 1})
		// calls printf or puts
		if funcName == "p_print_ln" {
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL puts", 1, 1})
		} else {
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL printf", 1, 1})
		}
		// Sets fflush argument
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "MOV r0, #0", 1, 1})
		// calls fflush
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL fflush", 1, 1})

		// pop {pc} to restore the caller address as the next instruction
		instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
	}
}

// Because the maximum amount we can add or subtract the stack pointer by is 1024
// These helper functions allocate and deallocate space on the stack for us

func (cg *CodeGenerator) createStackSpace(stackSize int) {
	if stackSize > STACK_SIZE_MAX {
		instructionSelection(InstructionSelection{cg.currInstrs(), "SUB sp, sp, #" + strconv.Itoa(STACK_SIZE_MAX), 1, 1})
		cg.createStackSpace(stackSize - STACK_SIZE_MAX)
	} else {
		instructionSelection(InstructionSelection{cg.currInstrs(), "SUB sp, sp, #" + strconv.Itoa(stackSize), 1, 1})
	}
}

// This cleans the stack
func (cg *CodeGenerator) removeStackSpace(stackSize int) {
	if stackSize > STACK_SIZE_MAX {
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD sp, sp, #" + strconv.Itoa(STACK_SIZE_MAX), 1, 1})
		cg.removeStackSpace(stackSize - STACK_SIZE_MAX)
	} else {
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD sp, sp, #" + strconv.Itoa(stackSize), 1, 1})
	}
}

// EVALUATION FUNCTIONS

// Evalutes the RHS of an expression
func (cg *CodeGenerator) evalRHS(t Evaluation, srcReg string) {

	switch t.(type) {
	// Literals
	case Integer:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", =" + strconv.Itoa(int(t.(Integer))), 1, 1})
	case Boolean:
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", #" + boolInt(bool(t.(Boolean))), 1, 1})
	case Character:
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", #" + string(t.(Character)), 1, 1})
	case Str:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", " + cg.getMsgLabel("", string(t.(Str))), 1, 1})
	case PairLiter:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", =0", 1, 1})
	case Ident:
		var offset, resType = cg.getIdentOffset(t.(Ident))
		switch resType.(type) {
		case ConstType:
			switch resType.(ConstType) {
			case Bool, Char:
				instructionSelection(InstructionSelection{cg.currInstrs(), "LDRSB " + srcReg + ", [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
			case Int, String:
				instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
			}
		default:
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
		}
	case ArrayElem:
		cg.evalArrayElem(t)
	case Unop:
		cg.cgVisitUnopExpr(t.(Unop))
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", r0", 1, 1})
	case Binop:
		cg.cgVisitBinopExpr(t.(Binop))
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", r0", 1, 1})
	case PairElem:
		cg.evalPairElem(t.(PairElem), srcReg)
	case Call:
		cg.cgVisitCallStat(t.(Call).Ident, t.(Call).ParamList, srcReg)
	case CallInstance:
		cg.cgVisitCallMethod(t.(CallInstance).Func,
			t.(CallInstance).Class, t.(CallInstance).ParamList, srcReg)
	case NewObject:
		cg.evalNewObject(t.(NewObject).Init, srcReg)
	case FieldAccess:
		classTyp := cg.eval(t.(FieldAccess).ObjectName).(ClassType)
		objClass := cg.getClass(classTyp)

		cg.evalField(t.(FieldAccess).ObjectName, t.(FieldAccess).Field, *objClass, srcReg)
	case ThisInstance:
		field := t.(ThisInstance).Field
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [sp, #" + strconv.Itoa(cg.getObjectOffset()) + "]", 1, 1})
		// Field Accumulator
		acc := 0
		for _, currField := range *cg.currStack.fieldList {
			if field == currField.Ident {
				//Stores the current initialiser value on the heap for the object in the correct offset
				instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [r5, #" + strconv.Itoa(acc) + "]", 1, 1})
				break
			}
			acc += sizeOf(currField.FieldType)
		}
	default:
		fmt.Println("ERROR: Expression can not be evaluated")
	}
}

// Evalute a pair element
func (cg *CodeGenerator) evalPairElem(t PairElem, srcReg string) {

	//Load the address of the pair from the stack
	var offset, _ = cg.getIdentOffset(t.Expr.(Ident))
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [sp, #" + strconv.Itoa(offset) + "]", 1, 1})

	//Check for null pointer deference
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, " + srcReg, 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_null_pointer", 1, 1})
	cg.dereferenceNullPointer()
	cg.cgVisitPrintStatFuncHelper("p_print_string")

	//Depending on fst or snd , load the address
	switch t.Fsnd {
	case Fst:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [" + srcReg + "]", 1, 1})
	case Snd:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [" + srcReg + ", #4]", 1, 1})
	}

	//Double deference
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [" + srcReg + "]", 1, 1})
}

// Helper to reduce code duplication
func (cg *CodeGenerator) evalNewPairHelper(pair Evaluation, reg1 string, reg2 string) {

	// Get the type
	var pairType = cg.eval(pair)
	var pairSize = sizeOf(pairType)

	// Load the element into a register to be stored
	cg.evalRHS(pair, reg1)

	//Allocate memory for the an element
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r0, =" + strconv.Itoa(pairSize), 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL malloc", 1, 1})

	//Store the element to the newly allocated memory onto the heap
	switch pairType.(type) {
	case ConstType:

		switch pairType.(ConstType) {
		case Bool, Char:
			instructionSelection(InstructionSelection{cg.currInstrs(), "STRB " + reg1 + ", [r0]", 1, 1})
		case Int, String, Pair:
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR " + reg1 + ", [r0]", 1, 1})
		}

	default:
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR " + reg1 + ", [r0]", 1, 1})
	}

}

// Evalutes a pair of elements onto the stack
func (cg *CodeGenerator) evalNewPair(fst Evaluation, snd Evaluation, reg1 string, reg2 string) {
	// First allocate memory to store two addresses (8-bytes)
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r0, =" + strconv.Itoa(ADDRESS_SIZE*2), 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL malloc", 1, 1})

	// Store the address in the free register
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + reg2 + ", r0", 1, 1})

	// Evalute the first element
	cg.evalNewPairHelper(fst, reg1, reg2)

	//Store the address of allocated memory block of the element on the heap
	instructionSelection(InstructionSelection{cg.currInstrs(), "STR r0, [" + reg2 + "]", 1, 1})

	// Evalute the second element
	cg.evalNewPairHelper(snd, reg1, reg2)

	//Store the address of allocated memory block of the element on the heap
	instructionSelection(InstructionSelection{cg.currInstrs(), "STR r0, [" + reg2 + ", #4]", 1, 1})
}

// Evaluates array literals
func (cg *CodeGenerator) evalArrayLiter(typeNode Type, rhs Evaluation, srcReg string, dstReg string) {

	//Calculate the amount of storage space required for the array
	// = ((arrayLength(array) * sizeOf(arrayType)) + INT_SIZE
	var arrayStorage = (len(rhs.(ArrayLiter).Exprs) * sizeOf(typeNode.(ArrayType).Type)) + INT_SIZE

	//Allocate memory for the array
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r0, =" + strconv.Itoa(arrayStorage), 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL malloc", 1, 1})

	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + dstReg + ", r0", 1, 1})

	//Start loading each element in the array onto the stack
	cg.evalArray(rhs.(ArrayLiter).Exprs, srcReg, dstReg, typeNode)

}

// Evalutes an array (where t is the type of the array)
func (cg *CodeGenerator) evalArray(array []Evaluation, srcReg string, dstReg string, t Type) {

	var arraySize = len(array)
	// Loop through the array pushing it onto the stack
	for i := 0; i < arraySize; i++ {
		// Array of pairs,ints,bools,chars,strings
		switch t.(type) {
		case ArrayType:
			cg.evalRHS(array[i], srcReg)
			switch t.(ArrayType).Type {
			case Int, String:
				instructionSelection(InstructionSelection{cg.currInstrs(), "STR " + srcReg + ", [" + dstReg + ", #" + strconv.Itoa(ARRAY_SIZE+sizeOf(t.(ArrayType).Type)*i) + "]", 1, 1})
			case Bool, Char:
				instructionSelection(InstructionSelection{cg.currInstrs(), "STRB " + srcReg + ", [" + dstReg + ", #" + strconv.Itoa(ARRAY_SIZE+sizeOf(t.(ArrayType).Type)*i) + "]", 1, 1})
			default:
				var offset, _ = cg.getIdentOffset(array[i].(Ident))
				switch t.(ArrayType).Type.(type) {
				case ArrayType:
					instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
				}
				instructionSelection(InstructionSelection{cg.currInstrs(), "STR " + srcReg + ", [r4, #" + strconv.Itoa(sizeOf(t.(ArrayType).Type)+sizeOf(t)*i) + "]", 1, 1})
			}
		}
	}
	// Put the size of the array onto the stack
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + srcReg + ", =" + strconv.Itoa(arraySize), 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "STR " + srcReg + ", [" + dstReg + "]", 1, 1})
}

// Evalutes array elements
func (cg *CodeGenerator) evalArrayElem(t Evaluation) {

	// Define error labels and formatting
	cg.getMsgLabel("", ARRAY_INDEX_NEGATIVE)
	cg.getMsgLabel("", ARRAY_INDEX_LARGE)
	cg.getMsgLabel("", STRING_FORMAT)

	// Store the address at the next space in the stack
	var offset, _ = cg.getIdentOffset(t.(ArrayElem).Ident)
	instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r4, sp, #" + strconv.Itoa(offset), 1, 1})

	//Check the bounds of the array
	cg.arrayCheckBounds(t.(ArrayElem).Exprs, t.(ArrayElem).Ident)

	// Add message labels
	cg.checkArrayBounds()
	cg.cgVisitPrintStatFuncHelper("p_print_string")

}

// Evalutes a ord
func (cg *CodeGenerator) evalOrd(node Unop) {
	switch node.Expr.(type) {
	case Ident:
		//If it's an ident
		var offset, _ = cg.getIdentOffset(node.Expr.(Ident))
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDRSB r0, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
	case ArrayElem:
		fmt.Println("ArrayElem not done for ord")
	case Character:
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, #" + string(node.Expr.(Character)), 1, 1})
	}
}

// Places an instance of a class of type classType within the register (as an address to the object)
func (cg *CodeGenerator) evalNewObject(initValues []Evaluation, dstReg string) {
	fieldSize := cg.evalSize(initValues)
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r0, =" + strconv.Itoa(fieldSize), 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL malloc", 1, 1})

	// dstReg = object address
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + dstReg + ", r0", 1, 1})

	// Accumulator for variable sizes
	acc := 0
	for _, initValue := range initValues {
		// Store the evaluation of the initial value into r6
		cg.evalRHS(initValue, "r6")
		//Stores the current initialiser value on the heap for the object in the correct offset
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR r6, [" + dstReg + ", #" + strconv.Itoa(acc) + "]", 1, 1})
		// Add to Accumulator
		acc += sizeOf(cg.eval(initValue))
	}

}

// Uses the object stored in srcReg to retrieve the value of the field ident and stores its value in dstReg
func (cg *CodeGenerator) evalField(objIdent Ident, field Ident, class Class, dstReg string) {
	// Stores object reference in r5
	cg.evalRHS(objIdent, "r5")

	// Field Accumulator
	acc := 0
	for _, currField := range class.FieldList {
		if field == currField.Ident {
			//Stores the current initialiser value on the heap for the object in the correct offset
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR " + dstReg + ", [r5, #" + strconv.Itoa(acc) + "]", 1, 1})
			break
		}
		acc += sizeOf(currField.FieldType)
	}

}

// VISIT FUNCTIONS -------------------------------------------------------------

// Visit Program
func (cg *CodeGenerator) cgVisitProgram(node *Program) {
	// Set properties of global scope
	cg.currStack.size = getScopeVarSize(node.StatList)
	cg.currStack.currP = cg.currStack.size
	cg.currStack.isFunc = false
	cg.classes = node.ClassList
	cg.functionList = node.FunctionList

	// .text
	instructionSelection(InstructionSelection{cg.funcInstrs, ".text", 0, 2})

	// .global main
	instructionSelection(InstructionSelection{cg.funcInstrs, ".global main", 0, 1})

	// main:
	instructionSelection(InstructionSelection{cg.currInstrs(), "main:", 0, 1})

	// Push {lr} to save the caller address
	instructionSelection(InstructionSelection{cg.currInstrs(), "PUSH {lr}", 1, 1})

	// sub sp, sp, #n to create variable space
	if cg.currStack.size > 0 {
		cg.createStackSpace(cg.globalStack.size)
	}

	// Traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}

	// add sp, sp, #n to remove variable space
	if cg.currStack.size > 0 {
		cg.removeStackSpace(cg.globalStack.size)
	}

	// ldr r0, =0 to return 0 as the main return
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r0, =0", 1, 1})

	// pop {pc} to restore the caller address as the next instruction
	instructionSelection(InstructionSelection{cg.currInstrs(), "POP {pc}", 1, 1})

	// .ltorg
	instructionSelection(InstructionSelection{cg.currInstrs(), ".ltorg", 1, 1})

	// Adds functions that were called
	for _, function := range cg.functionList {
		if cg.isFunctionDefined(function.Ident) {
			cg.cgVisitFunction(*function)
		}
	}

	// Adds methods that were called
	var methodName Ident
	for _, class := range cg.classes {
		for _, function := range class.FunctionList {
			methodName = Ident(string(class.Ident) + "." + string(function.Ident))
			if cg.isFunctionDefined(methodName) {
				cg.cgVisitMethod(*function, *class)
			}
		}
	}
}

// Evaluate a statement
func (cg *CodeGenerator) cgEvalStat(stat interface{}) {
	switch stat.(type) {
	case Declare:
		cg.cgVisitDeclareStat(stat.(Declare))
	case Assignment:
		cg.cgVisitAssignmentStat(stat.(Assignment))
	case Read:
		cg.cgVisitReadStat(stat.(Read))
	case Free:
		cg.cgVisitFreeStat(stat.(Free))
	case Return:
		cg.cgVisitReturnStat(stat.(Return))
	case Exit:
		cg.cgVisitExitStat(stat.(Exit))
	case Print:
		cg.cgVisitPrintStat(stat.(Print))
	case Println:
		cg.cgVisitPrintlnStat(stat.(Println))
	case If:
		cg.cgVisitIfStat(stat.(If))
	case While:
		cg.cgVisitWhileStat(stat.(While))
	case Scope:
		cg.cgVisitScopeStat(stat.(Scope))
	case Call:
		cg.cgVisitCallStat(stat.(Call).Ident, stat.(Call).ParamList, "r0")
	case CallInstance:
		cg.cgVisitCallMethod(stat.(CallInstance).Func,
			stat.(CallInstance).Class, stat.(CallInstance).ParamList, "r0")
	}
}

// Visit Declare node
func (cg *CodeGenerator) cgVisitDeclareStat(node Declare) {
	rhs := node.Rhs
	switch node.DecType.(type) {
	case ConstType:
		switch node.DecType.(ConstType) {
		case Bool, Char:
			cg.evalRHS(rhs, "r4")
			// Using STRB, store it on the stack
			instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [sp, #" + cg.subCurrP(sizeOf(node.DecType.(ConstType))) + "]", 1, 1})
		case Int:
			cg.evalRHS(rhs, "r4")
			// Store the value of declaration to stack
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(INT_SIZE) + "]", 1, 1})
		case String:
			// Store the address onto the stack
			switch rhs.(type) {
			case Str:
				instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, " + cg.getMsgLabel(node.Lhs, string(rhs.(Str))), 1, 1})
			default:
				cg.evalRHS(rhs, "r4")
			}
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(STRING_SIZE) + "]", 1, 1})
		}
	case PairType:
		switch rhs.(type) {
		case NewPair:
			cg.evalNewPair(rhs.(NewPair).FstExpr, rhs.(NewPair).SndExpr, "r5", "r4")
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
		case Ident:
			cg.evalRHS(rhs.(Ident), "r4")
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
		case PairLiter:
			//Can only be Null
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, =0", 1, 1})
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
		case Call:
			cg.evalRHS(rhs.(Call), "r4")
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(sizeOf(cg.eval(rhs.(Call)))) + "]", 1, 1})
		case PairElem:
			cg.evalPairElem(rhs.(PairElem), "r4")
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
		case ArrayElem:
			cg.evalArrayElem(rhs.(ArrayElem))
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
		}
	case ArrayType:
		switch rhs.(type) {
		case ArrayLiter:
			// Evalute an array
			cg.evalArrayLiter(node.DecType, rhs, "r5", "r4")
		case Ident:
			cg.evalRHS(rhs.(Ident), "r4")
		case PairElem:
			cg.evalPairElem(rhs.(PairElem), "r4")
		}
		// Now store the address of the array onto the stack
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
	case ClassType:
		cg.evalRHS(rhs, "r4")
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + cg.subCurrP(ADDRESS_SIZE) + "]", 1, 1})
	default:
		fmt.Println("ERROR: UNKNOWN DECLARE TYPE")
	}
	// Saves Idents offset in the symbol tables offset map
	cg.currSymTable().InsertOffset(string(node.Lhs), cg.currStack.currP)
}

// Visit Assignment node
func (cg *CodeGenerator) cgVisitAssignmentStat(node Assignment) {

	var rhs = node.Rhs
	var lhs = node.Lhs
	cg.evalRHS(rhs, "r4")

	// lhs can be
	// IDENT , ARRAY-ELEM , PAIR-ELEM
	switch lhs.(type) {
	case Ident:
		ident := lhs.(Ident)
		typeIdent := cg.eval(ident)
		switch typeIdent.(type) {
		case PairType:
			var offset, _ = cg.getIdentOffset(lhs.(Ident))
			switch rhs.(type) {
			case NewPair:
				cg.evalNewPair(rhs.(NewPair).FstExpr, rhs.(NewPair).SndExpr, "r5", "r4")
			}
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
		case ConstType:
			offset, _ := cg.getIdentOffset(ident)
			switch typeIdent.(ConstType) {
			case Bool, Char:
				// Using STRB, store it on the stack
				instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
			case Int, String:
				// Store the value of declaration to stack
				instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
			}
		case ClassType:
			offset, _ := cg.getIdentOffset(ident)
			instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
		}

	case ArrayElem:
		var offset, _ = cg.getIdentOffset(lhs.(ArrayElem).Ident)

		//Have a register point to the start of the array
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r5, sp, #" + strconv.Itoa(offset), 1, 1})

		//Load the index
		cg.evalRHS(lhs.(ArrayElem).Exprs[0], "r6")
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [r5]", 1, 1})

		//r6 = Index
		//r5 = Address of array
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r6", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r1, r5", 1, 1})

		//Branch
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_array_bounds", 1, 1})
		cg.checkArrayBounds()

		//Point to the element to be changed
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r5, r5, #4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r5, r5, r6, LSL #2", 1, 1})

		//Get the type of the RHS
		switch rhs.(type) {
		case Boolean, Character:
			instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [r5]", 1, 1})
		default:
			instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [r5]", 1, 1})
		}

	case PairElem:
		// Load the address of the pair into a register
		var offset, _ = cg.getIdentOffset(lhs.(PairElem).Expr.(Ident))
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [sp, #" + strconv.Itoa(offset) + "]", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r5", 1, 1})
		// Jump
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_null_pointer", 1, 1})
		cg.dereferenceNullPointer()
		// Check if it's the fst or snd
		switch lhs.(PairElem).Fsnd {
		case Fst:
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [r5]", 1, 1})
		case Snd:
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [r5, #4]", 1, 1})
		}
		// Store the value into the pair
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [r5]", 1, 1})
	case FieldAccess:
		objIdent := lhs.(FieldAccess).ObjectName
		field := lhs.(FieldAccess).Field
		classTyp := cg.eval(objIdent).(ClassType)
		class := cg.getClass(classTyp)

		cg.evalRHS(objIdent, "r5")

		// Field Accumulator
		acc := 0
		for _, currField := range class.FieldList {
			if field == currField.Ident {
				//Stores the current initialiser value on the heap for the object in the correct offset
				instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [r5, #" + strconv.Itoa(acc) + "]", 1, 1})
				break
			}
			acc += sizeOf(currField.FieldType)
		}
	case ThisInstance:
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r5, [sp, #" + strconv.Itoa(cg.getObjectOffset()) + "]", 1, 1})
		field := lhs.(ThisInstance).Field

		// Field Accumulator
		acc := 0
		for _, currField := range *cg.currStack.fieldList {
			if field == currField.Ident {
				//Stores the current initialiser value on the heap for the object in the correct offset
				instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [r5, #" + strconv.Itoa(acc) + "]", 1, 1})
				break
			}
			acc += sizeOf(currField.FieldType)
		}
	}
}

// Helper to remove duplication
func (cg *CodeGenerator) cgVisitReadStatHelper() {
	instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [sp]", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_null_pointer", 1, 1})
	cg.dereferenceNullPointer()
}

// Visit Read node
func (cg *CodeGenerator) cgVisitReadStat(node Read) {
	switch node.AssignLHS.(type) {
	case Ident:
		constType := cg.eval(node.AssignLHS.(Ident))
		offset, _ := cg.getIdentOffset(node.AssignLHS.(Ident))
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADD r0, sp, #" + strconv.Itoa(offset), 1, 1})
		switch constType {
		case Char:
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_read_char", 1, 1})
			cg.cgVisitReadStatFuncHelper("p_read_char")
		case Int:
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_read_int", 1, 1})
			cg.cgVisitReadStatFuncHelper("p_read_int")
		}
	case ArrayElem:
		//Complete
	case PairElem:
		cg.cgVisitReadStatHelper()
		switch node.AssignLHS.(PairElem).Fsnd {
		case Fst:
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [r4]", 1, 1})
		case Snd:
			instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [r4, #4]", 1, 1})
		}
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		switch cg.eval(node.AssignLHS.(PairElem)) {
		case Char:
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_read_char", 1, 1})
			cg.cgVisitReadStatFuncHelper("p_read_char")
		case Int:
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_read_int", 1, 1})
			cg.cgVisitReadStatFuncHelper("p_read_int")
		}
	}
}

// Visit Free node
func (cg *CodeGenerator) cgVisitFreeStat(node Free) {
	cg.evalRHS(node.Expr, "r4")
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_free_pair", 1, 1})
	cg.cgVisitFreeStatFuncHelper("p_free_pair")
}

// Visit Return node
func (cg *CodeGenerator) cgVisitReturnStat(node Return) {
	cg.evalRHS(node.Expr, "r0")
	funcVarUsed := funcVarSize(cg.currStack)
	if funcVarUsed > 0 {
		cg.removeStackSpace(funcVarUsed)
	}
	instructionSelection(InstructionSelection{cg.currInstrs(), "POP {pc}", 1, 1})
}

// Visit Exit node
func (cg *CodeGenerator) cgVisitExitStat(node Exit) {
	cg.evalRHS(node.Expr, "r4")
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL exit", 1, 1})
}

// Visit Print node
func (cg *CodeGenerator) cgVisitPrintStat(node Print) {

	// Get value of expr into dstReg
	cg.evalRHS(node.Expr, "r0")
	exprType := cg.eval(node.Expr)

	switch exprType.(type) {
	case ConstType:
		switch exprType.(ConstType) {
		case String:
			// BL p_print_string
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_string", 1, 1})
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_string")

		case Int:
			// BL p_print_int
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_int", 1, 1})
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_int")

		case Char:
			// BL putchar
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL putchar", 1, 1})

		case Bool:
			// BL p_print_bool
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_bool", 1, 1})
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_bool")
		case Pair:
			// BL p_print_reference
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_reference", 1, 1})
			// Define relevant print function definition (iff it hasnt been defined)
			cg.cgVisitPrintStatFuncHelper("p_print_reference")
		}
	case PairType, ArrayType, ClassType:

		switch exprType.(type) {
		case ArrayType:

			if exprType.(ArrayType).Type == Char {
				// BL p_print_string
				instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_string", 1, 1})
				// Define relevant print function definition (iff it hasnt been defined)
				cg.cgVisitPrintStatFuncHelper("p_print_string")
				return
			}
		}

		// BL p_print_reference
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_reference", 1, 1})
		// Define relevant print function definition (iff it hasnt been defined)
		cg.cgVisitPrintStatFuncHelper("p_print_reference")

	default:
		fmt.Println("Error: type not implemented")
	}

}

// Visit Println node
func (cg *CodeGenerator) cgVisitPrintlnStat(node Println) {

	cg.cgVisitPrintStat(Print{Expr: node.Expr})
	// BL p_print_ln
	instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_print_ln", 1, 1})
	// Define relevant print function definition (iff it hasnt been defined)
	cg.cgVisitPrintStatFuncHelper("p_print_ln")

}

// Visit If node
func (cg *CodeGenerator) cgVisitIfStat(node If) {
	fstLabel, sndLabel := cg.getNewLabel(), cg.getNewLabel()
	cg.evalRHS(node.Conditional, "r0")
	instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r0, #0", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BEQ " + fstLabel, 1, 1})

	cg.cgVisitScopeStat(Scope{StatList: node.ThenStat})

	instructionSelection(InstructionSelection{cg.currInstrs(), "B " + sndLabel, 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), fstLabel + ":", 0, 1})

	cg.cgVisitScopeStat(Scope{StatList: node.ElseStat})

	instructionSelection(InstructionSelection{cg.currInstrs(), sndLabel + ":", 0, 1})

}

// Visit While node
func (cg *CodeGenerator) cgVisitWhileStat(node While) {
	fstLabel, sndLabel := cg.getNewLabel(), cg.getNewLabel()
	instructionSelection(InstructionSelection{cg.currInstrs(), "B " + fstLabel, 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), sndLabel + ":", 0, 1})
	cg.cgVisitScopeStat(Scope{StatList: node.DoStat})
	instructionSelection(InstructionSelection{cg.currInstrs(), fstLabel + ":", 0, 1})
	cg.evalRHS(node.Conditional, "r0")
	instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r0, #1", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "BEQ " + sndLabel, 1, 1})
}

// Visit Scope node
func (cg *CodeGenerator) cgVisitScopeStat(node Scope) {
	// Amount of bytes on the stack the scope takes up for variables
	varSpaceSize := getScopeVarSize(node.StatList)
	cg.setNewScope(varSpaceSize)
	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}
	cg.removeCurrScope()
}

// Create a map of params to offsets
// make sure visitFunction can take a pointer to this map

// ONLY VISIT FUNCTION IF IT IS CALLED
// IE WE ONLY PUSH ONTO STACK FUNC VARIABLES WHEN A FUNCTION IS CALLED
// but
// WE EXECUTE WHAT IS INSIDE THE FUNCTION REGARDLESS OF WHETHER IT IS CALLED OR NOT
func (cg *CodeGenerator) cgVisitCallStat(ident Ident, paramList []Evaluation, srcReg string) {
	for _, function := range cg.functionList {
		if function.Ident == ident {
			for i := len(paramList) - 1; i >= 0; i-- {
				cg.cgVisitParameter(paramList[i])
			}
			instructionSelection(InstructionSelection{cg.currInstrs(), "BL f_" + string(function.Ident), 1, 1})
			offset := cg.cgGetParamSize(paramList)
			cg.subExtraOffset(offset)

			instructionSelection(InstructionSelection{cg.currInstrs(), "ADD sp, sp, #" + strconv.Itoa(offset), 1, 1})
			instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", r0", 1, 1})
			if !cg.isFunctionDefined(function.Ident) {
				cg.addFunctionDef(function.Ident)
			}
			break
		}
	}
}

func (cg *CodeGenerator) cgVisitCallMethod(ident Ident, objIdent Ident, paramList []Evaluation, srcReg string) {
	for _, class := range cg.classes {
		classTyp := cg.eval(objIdent).(ClassType)
		if class.Ident == classTyp {

			for _, function := range class.FunctionList {
				if function.Ident == ident {
					// Store object onto the stack
					cg.evalRHS(objIdent, "r4")
					instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [sp, #" + strconv.Itoa(-ADDRESS_SIZE) + "]!", 1, 1})

					for i := len(paramList) - 1; i >= 0; i-- {
						cg.cgVisitParameter(paramList[i])
					}
					instructionSelection(InstructionSelection{cg.currInstrs(), "BL f_" + string(classTyp) + "." + string(function.Ident), 1, 1})
					offset := cg.cgGetParamSize(paramList) + ADDRESS_SIZE
					//cg.subExtraOffset(offset)

					instructionSelection(InstructionSelection{cg.currInstrs(), "ADD sp, sp, #" + strconv.Itoa(offset), 1, 1})
					instructionSelection(InstructionSelection{cg.currInstrs(), "MOV " + srcReg + ", r0", 1, 1})

					methodName := Ident(string(classTyp) + "." + string(function.Ident))
					if !cg.isFunctionDefined(methodName) {
						cg.addFunctionDef(methodName)
					}
					break
				}

			}
		}
	}
}

func (cg *CodeGenerator) cgGetParamSize(paramList []Evaluation) int {
	totalCount := 0
	for _, param := range paramList {
		totalCount += sizeOf(cg.eval(param))
	}
	return totalCount
}

func (cg *CodeGenerator) cgVisitFunction(node Function) {
	varSpaceSize := getScopeVarSize(node.StatList)
	cg.setNewFuncScope(varSpaceSize, &node.ParameterTypes, node.SymbolTable)
	// f_funcName:
	instructionSelection(InstructionSelection{cg.currInstrs(), "f_" + string(node.Ident) + ":", 0, 1})
	// push {lr} to save the caller address
	instructionSelection(InstructionSelection{cg.currInstrs(), "PUSH {lr}", 1, 1})
	if varSpaceSize > 0 {
		cg.createStackSpace(varSpaceSize)
	}
	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}
	instructionSelection(InstructionSelection{cg.currInstrs(), ".ltorg", 1, 2})
	cg.removeFuncScope()
}

func (cg *CodeGenerator) cgVisitMethod(node Function, class Class) {
	varSpaceSize := getScopeVarSize(node.StatList)
	cg.setNewMethodScope(varSpaceSize, &node.ParameterTypes, node.SymbolTable, &class.FieldList)
	// f_funcName:
	instructionSelection(InstructionSelection{cg.currInstrs(), "f_" + string(class.Ident) + "." + string(node.Ident) + ":", 0, 1})
	// push {lr} to save the caller address
	instructionSelection(InstructionSelection{cg.currInstrs(), "PUSH {lr}", 1, 1})
	if varSpaceSize > 0 {
		cg.createStackSpace(varSpaceSize)
	}
	// traverse all statements by switching on statement type
	for _, stat := range node.StatList {
		cg.cgEvalStat(stat)
	}
	instructionSelection(InstructionSelection{cg.currInstrs(), ".ltorg", 1, 2})
	cg.removeFuncScope()
}

// VISIT STATEMENT -------------------------------------------------------------

func (cg *CodeGenerator) cgVisitParameter(node Evaluation) {
	resType := cg.eval(node)
	cg.evalRHS(node, "r4")
	switch resType {
	case Bool, Char:
		instructionSelection(InstructionSelection{cg.currInstrs(), "STRB r4, [sp, #" + strconv.Itoa(-sizeOf(resType)) + "]!", 1, 1})
	default:
		instructionSelection(InstructionSelection{cg.currInstrs(), "STR r4, [sp, #" + strconv.Itoa(-sizeOf(resType)) + "]!", 1, 1})
	}
	cg.addExtraOffset(sizeOf(resType))
}

// VISIT EXPRESSIONS -----------------------------------------------------------

// Visit Unop node
func (cg *CodeGenerator) cgVisitUnopExpr(node Unop) {
	switch node.Unary {
	case SUB:
		cg.evalRHS(node.Expr, "r4")
		//Negate the result in the register
		instructionSelection(InstructionSelection{cg.currInstrs(), "RSBS r4, r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1})
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case LEN:
		cg.evalRHS(node.Expr, "r4")
		//Assume the RHS is always an array
		//So the RHS type is an Ident
		//Load the length of the array into the register
		instructionSelection(InstructionSelection{cg.currInstrs(), "LDR r4, [r4]", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case ORD:
		cg.evalOrd(node)
	case CHR:
		cg.evalRHS(node.Expr, "r4")
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case NOT:
		cg.evalRHS(node.Expr, "r4")
		instructionSelection(InstructionSelection{cg.currInstrs(), "EOR r4, r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	}
}

// Visit Binop node
func (cg *CodeGenerator) cgVisitBinopExpr(node Binop) {
	cg.evalRHS(node.Left, "r4")
	instructionSelection(InstructionSelection{cg.currInstrs(), "PUSH {r4}", 1, 1})
	cg.addExtraOffset(4)
	cg.evalRHS(node.Right, "r4")
	instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r5, r4", 1, 1})
	instructionSelection(InstructionSelection{cg.currInstrs(), "POP {r4}", 1, 1})
	cg.subExtraOffset(4)
	switch node.Binary {
	case PLUS:
		instructionSelection(InstructionSelection{cg.currInstrs(), "ADDS r4, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case SUB:
		instructionSelection(InstructionSelection{cg.currInstrs(), "SUBS r4, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BLVS p_throw_overflow_error", 1, 1})
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case MUL:
		instructionSelection(InstructionSelection{cg.currInstrs(), "SMULL r4, r5, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r5, r4, ASR #31", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BLNE p_throw_overflow_error", 1, 1})
		cg.cgVisitBinopExprHelper("p_throw_overflow_error")
	case DIV:
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r1, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_divide_by_zero", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL __aeabi_idiv", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r4, r0", 1, 1})
		cg.cgVisitBinopExprHelper("p_check_divide_by_zero")
	case MOD:
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r1, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL p_check_divide_by_zero", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "BL __aeabi_idivmod", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r1", 1, 1})
		cg.cgVisitBinopExprHelper("p_check_divide_by_zero")
	case AND:
		instructionSelection(InstructionSelection{cg.currInstrs(), "AND r4, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case OR:
		instructionSelection(InstructionSelection{cg.currInstrs(), "ORR r4, r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case LT:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVLT r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVGE r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case GT:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVGT r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVLE r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case LTE:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVLE r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVGT r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case GTE:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVGE r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVLT r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case EQ:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVEQ r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVNE r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	case NEQ:
		instructionSelection(InstructionSelection{cg.currInstrs(), "CMP r4, r5", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVNE r4, #1", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOVEQ r4, #0", 1, 1})
		instructionSelection(InstructionSelection{cg.currInstrs(), "MOV r0, r4", 1, 1})
	}
}

// cgVisitBinopExpr helper function
// Adds a function definition to the progFuncInstrs ARMList depending on the
// function name provided
func (cg *CodeGenerator) cgVisitBinopExprHelper(funcName string) {
	if !cg.AddCheckProgName(funcName) {
		// funcLabel:
		instructionSelection(InstructionSelection{cg.progFuncInstrs, funcName + ":", 0, 1})
		switch funcName {
		case "p_throw_overflow_error":
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDR r0, " + cg.getMsgLabel("", OVERFLOW), 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BL p_throw_runtime_error", 1, 1})
		case "p_throw_runtime_error":
			//
		case "p_check_divide_by_zero":
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "PUSH {lr}", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "CMP r1, #0", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "LDREQ r0, " + cg.getMsgLabel("", DIVIDE_BY_ZERO), 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "BLEQ p_throw_runtime_error", 1, 1})
			instructionSelection(InstructionSelection{cg.progFuncInstrs, "POP {pc}", 1, 1})
		default:
		}
		cg.throwRunTimeError()
	}
}
