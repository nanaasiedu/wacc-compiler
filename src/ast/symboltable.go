package ast

import "fmt"

// SymbolTable constructor
type SymbolTable struct {
	Parent     *SymbolTable
	Table      map[Ident]Type
	OffsetVals map[string]int
	Children   []*SymbolTable
}

func (symbolTable *SymbolTable) PrintChildren() {
	if len(symbolTable.Children) == 0 {
		return
	}
	fmt.Println(symbolTable.Children)
	for _, sym := range symbolTable.Children {
		sym.PrintChildren()
	}
}

// New Constructor creates new instance of a symbolTable with pointer to its parent
func (symbolTable *SymbolTable) New() *SymbolTable {
	return &SymbolTable{Parent: symbolTable, Table: make(map[Ident]Type), OffsetVals: make(map[string]int), Children: []*SymbolTable{}}
}

// New Constructor creates new instance of a symbolTable with pointer to its parent
func NewInstance() *SymbolTable {
	return &SymbolTable{Table: make(map[Ident]Type), OffsetVals: make(map[string]int), Children: []*SymbolTable{}}
}

// Inserts a given key and value into the symbol table
func (symbolTable *SymbolTable) insert(key Ident, value Type) {
	symbolTable.Table[key] = value
}

// Inserts the offset of a given key into symbol table
func (symbolTable *SymbolTable) InsertOffset(key string, offset int) {
	symbolTable.OffsetVals[key] = offset
}

// Returns the offset of a given key
func (symbolTable *SymbolTable) GetOffset(key string) int {
	return symbolTable.OffsetVals[key]
}

// Checks if the key is already declared in the offset map
func (symbolTable *SymbolTable) IsOffsetDefined(key Ident) bool {
	if symbolTable == nil {
		return false
	}

	_, inMap := symbolTable.OffsetVals[string(key)]

	return inMap
}

// Checks if the key is already declared in the symbol table
func (symbolTable *SymbolTable) isDefined(key Ident) bool {
	if symbolTable == nil {
		return false
	}
	if symbolTable.contains(key) {
		return true
	} else {
		if symbolTable.Parent == nil {
			return false
		}
		return symbolTable.Parent.isDefined(key)
	}
}

// Checks if the key is already declared in the symbol table
func (symbolTable *SymbolTable) isDefinedInScope(key Ident) bool {
	return symbolTable.contains(key)
}

// Helper function which return a boolean depending on
// whether or not the given key is in the symbol table
func (symbolTable *SymbolTable) contains(key Ident) bool {
	_, ok := symbolTable.Table[key]
	return ok
}

// Given an Ident key, returns the slice of Type tokens
func (symbolTable *SymbolTable) getTypeOfIdent(key Ident) Type {
	if symbolTable == nil {
		return nil
	}
	if symbolTable.contains(key) {
		k := symbolTable.Table[key]
		return k
	} else {
		return symbolTable.Parent.getTypeOfIdent(key)
	}
}

func (symbolTable *SymbolTable) GetTypeOfIdent(key Ident) Type {
	return symbolTable.getTypeOfIdent(key)
}

// Removes the head of the children list
func (symbolTable *SymbolTable) RemoveChild() {
	if len(symbolTable.Children) == 0 {
		return
	}
	symbolTable.Children = symbolTable.Children[1:]
}

// Returns head of the children list
func (symbolTable *SymbolTable) GetFrontChild() *SymbolTable {
	return symbolTable.Children[0]
}
