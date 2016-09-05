package armarchitecture

import (
	"fmt"
	"reflect"
)

// Operand of ARM instruction set
type Operand int

// Operand enum

// CHECK WHICH OPERANDS WE DON@T USE AND REMOVE THEM
const (
	ADDS Operand = iota
	B
	BEQ
	BL
	BLCS
	BLEQ
	BLLT
	BLNE
	BLVS
	BNE
	CMP
	DIV
	EOR
	LDR
	LDRCS
	LDREQ
	LDRLT
	LDRNE
	LDRSB
	MOV
	MOVEQ
	MOVGE
	MOVGT
	MOVLE
	MOVLT
	MOVNE
	MUL
	POP
	POPEQ
	PUSH
	RSBS
	SMULL
	STR
	STRB
	SUB
	SUBS
)

// Address of ARM instruction set:
// * [value]
// * [value, #offset]
type Address struct {
	value  interface{}
	offset int
}

// Returns string representation of Address
func (a Address) String() string {
	if a.offset == 0 {
		return fmt.Sprint("[", a.value, "]")
	}
	return fmt.Sprint("[", a.value, ", ", a.offset, "]")
}

// Branch of ARM instruction set
type Branch struct {
	branch int
}

// Returns string representation of Branch
func (b Branch) String() string {
	return fmt.Sprint("L", b.branch, ":")
}

// Constant of ARM instruction set
type Constant struct {
	value string // value can be string or int. If int then convert to string in caller
}

// Returns string representation of Constant
func (c Constant) String() string {
	return "=" + c.value
}

// Directive of ARM instuction set
type Directive struct {
	directive string
}

// Returns the string representation of a Directive
func (d Directive) String() string {
	return "\t." + d.directive
}

// Immediate value of ARM instruction set
type Immediate struct {
	value string // can be int or string, if int convert to string in caller
}

// Returns the string representation of Immediate value
func (i Immediate) String() string {
	return "#" + i.value
}

// Label of ARM instruction set
type Label struct {
	label string
}

// Returns the string representation of Label
func (l Label) String() string {
	return l.label
}

// Operation of ARM instruction set
type Operation struct {
	operand  Operand
	listVars []interface{} // can be up to at least 3 vars
	// ISUE WILL OCCURE WHEN WE HAVE A MIXTURE OF REGISTERS AND INTS IN THE SLICE
	// CAN't have a SLICE WITH DIFFERENT TYPES
}

// Returns string representation of Operation
func (o Operation) String() string {
	result := "\t" + fmt.Sprint(o.operand) + " "

	list := reflect.ValueOf(o.listVars)
	for i := 0; i < list.Len(); i++ {
		result += fmt.Sprint(list.Index(i)) // HOW TO CONVERT TO STRING WHEN INT // REGISTER
		if i+1 != list.Len() {
			result += ", "
		}
	}
	return result
}

// Register of ARM instruction set
type Register struct {
	value  int
	offset int
	inUse  bool
}

// CheckInUse returns a boolean indicating if a given register is in use
func (r Register) CheckInUse() bool {
	return r.inUse
}

// SetInUse sets a current register in use
func (r Register) SetInUse() {
	r.inUse = true
}

func (r Register) String() string {
	offsetString := ""
	if r.offset != 0 {
		offsetString = ", #" + fmt.Sprint(r.offset)
	}

	//TODO: FIX THESE MAGIC NUMBERS
	switch r.value {
	case 13:
		return "sp" + offsetString
	case 14:
		return "lr" + offsetString
	case 15:
		return "pc" + offsetString
	default:
		return "r" + fmt.Sprint(r.value) + offsetString
	}
}

// Variable interface is a superclass of the following:
// * register
// * constant
// * address
// * empty_variable
type Variable interface {
	String() string
}
