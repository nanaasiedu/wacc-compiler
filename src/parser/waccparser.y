%{
package parser

import (
. "ast"
)

%}

%union{
str         string
stringconst Str
number      int
pos         int
integer     Integer
ident       Ident
character   Character
boolean     Boolean
fieldaccess Evaluation
functions      []*Function
function       *Function
classes        []*Class
class          *Class
stmt           Statement
stmts          []Statement
assignrhs      Evaluation
assignlhs      Evaluation
expr           Evaluation
exprs          []Evaluation
params         []Param
param          Param
fields         []Field
field          Field
bracketed      []Evaluation
pairliter      Evaluation
arrayliter     ArrayLiter
pairelem       PairElem
arrayelem      ArrayElem
typedefinition Type
pairelemtype   Type
}

%start program

%token BEGIN END                                    // Program delimiters
%token CLASS OPEN CLOSE NEW                         // Class delimiters
%token DOT
%token THIS
%token IS
%token <number> SKIP
%token READ FREE RETURN EXIT PRINT PRINTLN
%token IF THEN ELSE FI                              // If statement
%token WHILE DO DONE                                // While statement
%token NEWPAIR
%token CALL
%token FST SND
%token INT BOOL CHAR STRING PAIR
%token NOT NEG LEN ORD CHR                         // Unary ops
%token MUL DIV MOD PLUS SUB AND OR GT GTE LT LTE EQ NEQ // Binary ops
%token POSITIVE NEGATIVE
%token <boolean> TRUE FALSE                         // Booleans
%token NULL
%token OPENSQUARE OPENROUND CLOSESQUARE CLOSEROUND
%token ASSIGNMENT
%token COMMA SEMICOLON
%token ERROR
%token FOR


%token <stringconst> STRINGCONST
%token <ident> IDENTIFIER
%token <integer> INTEGER
%token <character> CHARACTER

%type <prog> program
%type <classes> classes
%type <class> class
%type <functions> functions
%type <function> function
%type <stmt> statement
%type <stmts> statements
%type <assignrhs> assignrhs
%type <assignlhs> assignlhs
%type	<expr>	expr
%type <exprs> exprlist
%type <params> paramlist
%type <param> param
%type <fields> fieldlist
%type <field> field
%type <pairelem> pairelem
%type <arrayliter> arrayliter
%type <arrayelem> arrayelem
%type <exprs> bracketed
%type <pairliter> pairliter
%type <typedefinition> basetype typeDef arraytype pairtype
%type <pairelemtype> pairelemtype
%type <stmt>         assignment
%type <fieldaccess> fieldaccess

%left OR
%left AND
%left EQ NEQ
%left PLUS SUB
%left MUL DIV MOD
%left LT GT LTE GTE
%right NOT NEG LEN ORD CHR

%%

program : BEGIN classes functions statements END {
                                                  parserlex.(*Lexer).prog = &Program{ClassList : $2 , FunctionList : $3 , StatList : $4 , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
                                                 }


classes : classes class  { $$ = append($1, $2)}
        |                { $$ = []*Class{} }


class : CLASS IDENTIFIER OPEN fieldlist functions CLOSE { if !checkClassIdent($2) {
                                                         	parserlex.Error("Invalid class name")
                                                     }
                                                     $$ = &Class{ Pos : $<pos>1, FileText :&parserlex.(*Lexer).input, Ident : ClassType($2), FieldList : $4 , FunctionList : $5}
                                                   }

fieldlist : fieldlist COMMA field { $$ = append($1, $3)}
          | field                 { $$ = []Field{ $1 } }
      //  |                       { $$ = []Field{}     }

field : typeDef IDENTIFIER { $$ = Field{FieldType : $1, Ident : $2} }

functions : functions function  { $$ = append($1, $2)}
          |                     { $$ = []*Function{} }

function : typeDef IDENTIFIER OPENROUND CLOSEROUND IS statements END
           { if !checkStats($6) {
          	parserlex.Error("Missing return statement")
           }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $6, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
           }
         | typeDef IDENTIFIER OPENROUND paramlist CLOSEROUND IS statements END
           { if !checkStats($7) {
            	parserlex.Error("Missing return statement")
            }
             $$ = &Function{Ident : $2, ReturnType : $1, StatList : $7, ParameterTypes : $4, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
           }

paramlist : paramlist COMMA param { $$ = append($1, $3)}
          | param                 { $$ = []Param{ $1 } }

param : typeDef IDENTIFIER { $$ = Param{ParamType : $1, Ident : $2} }

assignlhs : IDENTIFIER    {$$ = $1}
          | arrayelem     {$$ = $1}
          | pairelem      {$$ = $1}
          | fieldaccess   { $$ = $1}
          | THIS DOT IDENTIFIER                  { $$ = ThisInstance{&parserlex.(*Lexer).input, $<pos>1, $3} }

fieldaccess : IDENTIFIER DOT IDENTIFIER {$$ = FieldAccess{ &parserlex.(*Lexer).input, $<pos>1, $1, $3, } }

assignrhs : expr                                           {$$ = $1}
          | arrayliter                                     {$$ = $1}
          | pairelem                                       {$$ = $1}
          | NEW IDENTIFIER OPENROUND exprlist CLOSEROUND   { $$ = NewObject{Class : ClassType($2) , Init : $4 , Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}}

statements : statements SEMICOLON statement                { $$ = append($1,$3)   }
           | statement                                     { $$ = []Statement{$1} }
        //   |                                               { $$ = []Statement{} }
           | FOR typeDef IDENTIFIER ASSIGNMENT assignrhs SEMICOLON expr SEMICOLON assignment DO statements DONE {
                                                                                                                 stats := append($11, $9)
                                                                                                                 w := While{Conditional : $7, DoStat : stats, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}
                                                                                                                 d := Declare{DecType : $2, Lhs : $3, Rhs : $5, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input }
                                                                                                                 $$ = []Statement{d,w}
                                                                                                                }

statement : SKIP                                        { $$ = Skip{Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input } }
          | typeDef IDENTIFIER ASSIGNMENT assignrhs     { $$ = Declare{DecType : $1, Lhs : $2, Rhs : $4, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input } }
          | assignment                                  { $$ = $1 }
          | READ assignlhs                              { $$ = Read{ &parserlex.(*Lexer).input, $<pos>1 , $2, } }
          | FREE expr                                   { $$ = Free{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | RETURN expr                                 { $$ = Return{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | EXIT expr                                   { $$ = Exit{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | PRINT expr                                  { $$ = Print{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | PRINTLN expr                                { $$ = Println{&parserlex.(*Lexer).input, $<pos>1, $2} }
          | IF expr THEN statements ELSE statements FI  { $$ = If{Conditional : $2, ThenStat : $4, ElseStat : $6, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }
          | FOR expr SEMICOLON assignment DO statements DONE {
                                                              stats := append($6, $4)
                                                              $$ = While{Conditional : $2, DoStat : stats, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}
                                                             }
          | WHILE expr DO statements DONE               { $$ = While{Conditional : $2, DoStat : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input} }
          | BEGIN statements END                        { $$ = Scope{StatList : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }
          | error SEMICOLON                             {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error END                                   { parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error FI                                     {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | error DONE                                   {
                                                          parserlex.Error("Syntax error : Invalid statement")
                                                          $$ = nil
                                                        }
          | CALL IDENTIFIER OPENROUND exprlist CLOSEROUND    { $$ = Call{Ident : $2, ParamList : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
          | CALL fieldaccess OPENROUND exprlist CLOSEROUND    { $$ = CallInstance{Class : ($2.(FieldAccess)).ObjectName, Func: ($2.(FieldAccess)).Field, ParamList : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }

assignment :  assignlhs ASSIGNMENT assignrhs               { $$ = Assignment{Lhs : $1, Rhs : $3, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER PLUS ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : PLUS, Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER SUB  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : SUB , Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER DIV  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : DIV,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER MUL  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MUL,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER MOD  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MOD,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER PLUS PLUS                        { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : PLUS, Right : Integer(1), Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER SUB  SUB                         { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : SUB,  Right : Integer(1), Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | IDENTIFIER MUL  MUL                         { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MUL,  Right : $1,         Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }

             | fieldaccess PLUS ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : PLUS, Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess SUB  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : SUB , Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess DIV  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : DIV,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess MUL  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MUL,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess MOD  ASSIGNMENT expr             { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MOD,  Right : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess PLUS PLUS                        { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : PLUS, Right : Integer(1), Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess SUB  SUB                         { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : SUB,  Right : Integer(1), Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }
             | fieldaccess MUL  MUL                         { $$ = Assignment{Lhs : $1, Rhs : Binop{Left : $1, Binary : MUL,  Right : $1,         Pos : $<pos>1, FileText :&parserlex.(*Lexer).input}, Pos : $<pos>1 ,FileText :&parserlex.(*Lexer).input} }

expr : INTEGER        { $$ =  $1 }
     | TRUE           { $$ =  $1 }
     | FALSE          { $$ =  $1 }
     | CHARACTER      { $$ =  $1 }
     | STRINGCONST    { $$ =  $1 }
     | fieldaccess    { $$ = $1 }
     | IDENTIFIER     { $$ = $1}
     | pairliter      { $$ =  $1 }
     | arrayelem      { $$ =  $1 }
     | NOT expr       { $$ = Unop{Unary : NOT, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | LEN expr       { $$ = Unop{Unary : LEN, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | ORD expr       { $$ = Unop{Unary : ORD, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | CHR expr       { $$ = Unop{Unary : CHR, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | SUB expr       { $$ = Unop{Unary : SUB, Expr : $2, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | PLUS expr      { $$ = $2 }
     | expr PLUS expr { $$ = Binop{Left : $1, Binary : PLUS, Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr SUB expr  { $$ = Binop{Left : $1, Binary : SUB,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr MUL expr  { $$ = Binop{Left : $1, Binary : MUL,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr MOD expr  { $$ = Binop{Left : $1, Binary : MOD,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr DIV expr  { $$ = Binop{Left : $1, Binary : DIV,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr LT expr   { $$ = Binop{Left : $1, Binary : LT,   Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr GT expr   { $$ = Binop{Left : $1, Binary : GT,   Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr LTE expr  { $$ = Binop{Left : $1, Binary : LTE,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr GTE expr  { $$ = Binop{Left : $1, Binary : GTE,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr EQ expr   { $$ = Binop{Left : $1, Binary : EQ,   Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr NEQ expr  { $$ = Binop{Left : $1, Binary : NEQ,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr AND expr  { $$ = Binop{Left : $1, Binary : AND,  Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | expr OR expr   { $$ = Binop{Left : $1, Binary : OR,   Right : $3, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | OPENROUND expr CLOSEROUND  { $$ = $2 }
     | CALL IDENTIFIER OPENROUND exprlist CLOSEROUND    { $$ = Call{Ident : $2, ParamList : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | CALL fieldaccess OPENROUND exprlist CLOSEROUND   { $$ = CallInstance{Class : ($2.(FieldAccess)).ObjectName, Func: ($2.(FieldAccess)).Field, ParamList : $4, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input  } }
     | THIS DOT IDENTIFIER                                 { $$ = ThisInstance{&parserlex.(*Lexer).input, $<pos>1, $3} }
     | NEWPAIR OPENROUND expr COMMA expr CLOSEROUND   { $$ = NewPair{FstExpr : $3, SndExpr : $5, Pos : $<pos>1, FileText :&parserlex.(*Lexer).input } }

arrayliter : OPENSQUARE exprlist CLOSESQUARE { $$ = ArrayLiter{&parserlex.(*Lexer).input, $<pos>1, $2 } }

exprlist : exprlist COMMA expr {$$ = append($1, $3)}
         | expr                {$$ = []Evaluation{$1}}
         |                     {$$ = []Evaluation{}}

arrayelem : IDENTIFIER bracketed {$$ = ArrayElem{Ident: $1, Exprs : $2, Pos : $<pos>1,FileText :&parserlex.(*Lexer).input  } }

bracketed : bracketed OPENSQUARE expr CLOSESQUARE {$$ = append($1, $3)}
          | OPENSQUARE expr CLOSESQUARE {$$ = []Evaluation{$2}}

pairliter : NULL    { $$ =  PairLiter{} }

pairelem : FST expr { $$ = PairElem{Fsnd: Fst, Expr : $2, Pos : $<pos>1  } }
         | SND expr { $$ = PairElem{Fsnd: Snd, Expr : $2, Pos : $<pos>1  } }

pairtype : PAIR OPENROUND pairelemtype COMMA pairelemtype CLOSEROUND  { $$ = PairType{FstType : $3, SndType : $5} }

pairelemtype : basetype   { $$ = $1 }
             | arraytype  { $$ = $1 }
             | PAIR       { $$ = Pair}
             | pairtype   { $$ = $1}
             | IDENTIFIER { $$ = ClassType($1)}

typeDef : basetype   { $$ =  $1 }
        | arraytype  { $$ =  $1 }
        | pairtype   { $$ =  $1 }
        | IDENTIFIER { $$ = ClassType($1) }

basetype : INT      { $$ =  Int }
         | BOOL     { $$ =  Bool }
         | CHAR     { $$ =  Char }
         | STRING   { $$ =  String }

arraytype : basetype OPENSQUARE CLOSESQUARE  { $$ = ArrayType{Type : $1} }
          | pairtype OPENSQUARE CLOSESQUARE  { $$ = ArrayType{Type : $1} }
          | arraytype OPENSQUARE CLOSESQUARE { $$ = ArrayType{Type : $1} }

%%
