package ast

type errorSlice []error

func (e errorSlice) Error() string {
	var result string
	for _, errs := range e {
		result += errs.Error() + "\n"
	}
	return result
}

type ConstType int
type FSND int

type Type interface {
	typeString() string
}

type Context struct {
	FunctionTable []*Function
	SymbolTable   *SymbolTable
	ClassTable    []*Class
}

type Evaluation interface {
	Eval(context *Context) (Type, error)
}

type Statement interface {
	visitStatement(context *Context) errorSlice
}

type Ident string

type ClassType string

type Integer int
type Character string
type Str string
type Boolean bool
type PairLiter struct {
}

func (t ConstType) String() string {
	return t.typeString()
}

func (t ArrayType) String() string {
	return t.typeString()
}

func (t PairType) String() string {
	return t.typeString()
}

func (x ConstType) typeString() string {
	switch x {
	case Int:
		return "Int"
	case Bool:
		return "Bool"
	case Char:
		return "Char"
	case String:
		return "String"
	case Pair:
		return "Pair"
	default:
		return "Non-type"
	}
}

func (x ArrayType) typeString() string {
	return "Array:" + x.Type.typeString()
}

func (x ClassType) typeString() string {
	return "Class:" + string(x)
}

func (x PairType) typeString() string {
	return "Pair:" + x.FstType.typeString() + "/" + x.SndType.typeString()
}

func (x Ident) typeString() string {
	return "Fix this"
}

type Skip struct {
	FileText *string
	Pos      int
}

const (
	Int ConstType = iota
	Bool
	Char
	String
	Pair
)

// ArrayType struct
type ArrayType struct {
	Type Type
}

// PairType struct
type PairType struct {
	FstType Type
	SndType Type
}

const (
	Fst FSND = iota
	Snd
)

type Function struct {
	FileText       *string
	ClassDef       Ident
	Ident          Ident
	ReturnType     Type
	ParameterTypes []Param
	StatList       []Statement
	SymbolTable    *SymbolTable
}

type Param struct {
	Ident     Ident
	ParamType Type
}

// Program ..
type Program struct {
	FileText     *string
	Pos          int
	ClassList    []*Class
	FunctionList []*Function
	StatList     []Statement
	SymbolTable  *SymbolTable
}

type NewObject struct {
	FileText *string
	Pos      int
	Class    ClassType
	Init     []Evaluation
}

type FieldAccess struct {
	FileText *string
	Pos      int
	ObjectName Ident
	Field      Ident
}

type Class struct {
	FileText  *string
	Pos       int
	Ident        ClassType
	FieldList    []Field
	FunctionList []*Function
}

type Field struct {
	Ident     Ident
	FieldType Type
}

type CallInstance struct {
	FileText  *string
	Pos       int
	Class     Ident
	Func      Ident
	ParamList []Evaluation
}

type ThisInstance struct {
	FileText  *string
	Pos       int
	Field Ident
}

type Instance struct {
	IdentLHS Ident
	IdentRHS Ident
}

// Binop struct
type Binop struct {
	FileText *string
	Pos      int
	Binary   int
	Left     Evaluation
	Right    Evaluation
}

// Unop struct
type Unop struct {
	FileText *string
	Pos      int
	Unary    int
	Expr     Evaluation
}

type NewPair struct {
	FileText *string
	Pos      int
	FstExpr  Evaluation
	SndExpr  Evaluation
}

// Declare struct
type Declare struct {
	FileText *string
	Pos      int
	DecType  Type
	Lhs      Ident
	Rhs      Evaluation
}

// Assignment struct
type Assignment struct {
	FileText *string
	Pos      int
	Lhs      Evaluation
	Rhs      Evaluation
}

// If struct
type If struct {
	FileText    *string
	Pos         int
	Conditional Evaluation
	ThenStat    []Statement
	ElseStat    []Statement
}

// While struct
type While struct {
	FileText    *string
	Pos         int
	Conditional Evaluation
	DoStat      []Statement
}

type Scope struct {
	FileText *string
	Pos      int
	StatList []Statement
}

// Read struct
type Read struct {
	FileText  *string
	Pos       int
	AssignLHS Evaluation // should be an assignLHS
}

// Free struct
type Free struct {
	FileText *string
	Pos      int
	Expr     Evaluation
}

// Return struct
type Return struct {
	FileText *string
	Pos      int
	Expr     Evaluation
}

// Exit struct
type Exit struct {
	FileText *string
	Pos      int
	Expr     Evaluation
}

// Print struct
type Print struct {
	FileText *string
	Pos      int
	Expr     Evaluation
}

// Println struct
type Println struct {
	FileText *string
	Pos      int
	Expr     Evaluation
}

type Call struct {
	FileText  *string
	Pos       int
	Ident     Ident
	ParamList []Evaluation
}

/*
type Ident struct {
	Name string
}
*/
type PairElem struct {
	FileText *string
	Pos      int
	Fsnd     FSND
	Expr     Evaluation
}

type ArrayLiter struct {
	FileText *string
	Pos      int
	Exprs    []Evaluation
}

type ArrayElem struct {
	FileText *string
	Pos      int
	Ident    Ident
	Exprs    []Evaluation
}
