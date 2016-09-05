
//line waccparser.y:2
package parser
import __yyfmt__ "fmt"
//line waccparser.y:2

import (
. "ast"
)


//line waccparser.y:10
type parserSymType struct{
	yys int
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



var parserToknames = []string{
	"BEGIN",
	"END",
	"CLASS",
	"OPEN",
	"CLOSE",
	"NEW",
	"DOT",
	"THIS",
	"IS",
	"SKIP",
	"READ",
	"FREE",
	"RETURN",
	"EXIT",
	"PRINT",
	"PRINTLN",
	"IF",
	"THEN",
	"ELSE",
	"FI",
	"WHILE",
	"DO",
	"DONE",
	"NEWPAIR",
	"CALL",
	"FST",
	"SND",
	"INT",
	"BOOL",
	"CHAR",
	"STRING",
	"PAIR",
	"NOT",
	"NEG",
	"LEN",
	"ORD",
	"CHR",
	"MUL",
	"DIV",
	"MOD",
	"PLUS",
	"SUB",
	"AND",
	"OR",
	"GT",
	"GTE",
	"LT",
	"LTE",
	"EQ",
	"NEQ",
	"POSITIVE",
	"NEGATIVE",
	"TRUE",
	"FALSE",
	"NULL",
	"OPENSQUARE",
	"OPENROUND",
	"CLOSESQUARE",
	"CLOSEROUND",
	"ASSIGNMENT",
	"COMMA",
	"SEMICOLON",
	"ERROR",
	"FOR",
	"STRINGCONST",
	"IDENTIFIER",
	"INTEGER",
	"CHARACTER",
}
var parserStatenames = []string{}

const parserEofCode = 1
const parserErrCode = 2
const parserMaxDepth = 200

//line waccparser.y:300


//line yacctab:1
var parserExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 28,
	69, 116,
	-2, 15,
	-1, 46,
	69, 116,
	-2, 70,
}

const parserNprod = 124
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 693

var parserAct = []int{

	212, 206, 7, 180, 211, 37, 172, 158, 4, 250,
	265, 45, 240, 231, 13, 221, 70, 72, 73, 74,
	75, 76, 77, 52, 27, 78, 26, 25, 30, 202,
	178, 147, 31, 32, 33, 34, 35, 141, 69, 132,
	105, 106, 79, 110, 84, 66, 30, 11, 85, 41,
	253, 230, 274, 44, 271, 228, 125, 126, 127, 128,
	129, 130, 131, 241, 54, 278, 170, 30, 267, 36,
	182, 31, 32, 33, 34, 35, 38, 269, 272, 36,
	244, 38, 137, 256, 140, 244, 133, 36, 169, 31,
	32, 33, 34, 35, 39, 40, 43, 157, 159, 39,
	40, 242, 204, 161, 43, 264, 43, 43, 36, 182,
	45, 81, 43, 186, 187, 188, 189, 190, 191, 192,
	193, 194, 195, 196, 197, 198, 184, 182, 176, 82,
	174, 173, 83, 42, 185, 30, 203, 43, 159, 68,
	43, 209, 210, 161, 43, 208, 183, 171, 213, 214,
	181, 215, 248, 216, 217, 244, 218, 219, 255, 245,
	244, 244, 30, 30, 243, 223, 244, 224, 222, 225,
	226, 80, 227, 238, 167, 239, 36, 165, 154, 207,
	151, 168, 136, 149, 159, 137, 155, 152, 229, 161,
	97, 232, 166, 43, 276, 146, 164, 145, 150, 23,
	153, 22, 148, 36, 36, 234, 235, 144, 38, 89,
	12, 14, 15, 16, 17, 18, 19, 20, 89, 136,
	247, 21, 103, 201, 143, 24, 39, 40, 31, 32,
	33, 34, 35, 135, 252, 249, 254, 257, 88, 87,
	258, 260, 89, 89, 261, 262, 31, 32, 33, 34,
	175, 86, 263, 176, 266, 174, 173, 30, 259, 200,
	156, 30, 270, 233, 10, 30, 28, 237, 142, 9,
	29, 134, 251, 181, 93, 92, 94, 90, 91, 277,
	104, 107, 207, 30, 177, 67, 6, 63, 275, 2,
	53, 96, 96, 118, 120, 117, 119, 30, 36, 95,
	30, 160, 36, 64, 62, 179, 36, 31, 32, 33,
	34, 35, 55, 108, 56, 57, 58, 205, 8, 5,
	60, 59, 3, 1, 36, 101, 100, 102, 98, 99,
	0, 0, 48, 49, 65, 162, 61, 63, 36, 0,
	0, 36, 0, 0, 51, 46, 47, 50, 0, 0,
	114, 116, 115, 64, 62, 39, 40, 118, 120, 117,
	119, 0, 55, 0, 56, 57, 58, 0, 0, 0,
	60, 59, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 49, 65, 163, 61, 0, 0, 23,
	0, 22, 0, 0, 51, 71, 47, 50, 38, 0,
	12, 14, 15, 16, 17, 18, 19, 20, 0, 0,
	0, 21, 63, 0, 0, 24, 39, 40, 31, 32,
	33, 34, 35, 114, 116, 115, 112, 113, 64, 62,
	118, 120, 117, 119, 121, 122, 0, 55, 0, 56,
	57, 58, 0, 0, 0, 60, 59, 0, 0, 0,
	0, 0, 0, 0, 109, 0, 28, 48, 49, 65,
	0, 61, 0, 0, 0, 0, 0, 0, 0, 51,
	71, 47, 50, 114, 116, 115, 112, 113, 123, 124,
	118, 120, 117, 119, 121, 122, 114, 116, 115, 112,
	113, 0, 0, 118, 120, 117, 119, 273, 114, 116,
	115, 112, 113, 123, 124, 118, 120, 117, 119, 121,
	122, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 111, 114, 116, 115, 112, 113, 123, 124,
	118, 120, 117, 119, 121, 122, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 236, 114, 116, 115,
	112, 113, 123, 124, 118, 120, 117, 119, 121, 122,
	0, 0, 0, 0, 0, 0, 0, 0, 268, 114,
	116, 115, 112, 113, 123, 124, 118, 120, 117, 119,
	121, 122, 0, 0, 0, 0, 0, 0, 0, 0,
	199, 114, 116, 115, 112, 113, 123, 124, 118, 120,
	117, 119, 121, 122, 0, 0, 0, 0, 0, 0,
	0, 246, 114, 116, 115, 112, 113, 123, 124, 118,
	120, 117, 119, 121, 122, 139, 0, 0, 0, 0,
	0, 0, 220, 0, 138, 0, 0, 0, 0, 0,
	0, 114, 116, 115, 112, 113, 123, 124, 118, 120,
	117, 119, 121, 122, 114, 116, 115, 112, 113, 123,
	124, 118, 120, 117, 119, 121, 122, 114, 116, 115,
	112, 113, 123, 124, 118, 120, 117, 119, 121, 122,
	114, 116, 115, 112, 113, 123, 0, 118, 120, 117,
	119, 121, 122,
}
var parserPact = []int{

	285, -1000, -1000, 280, 197, -1000, -20, 128, -1000, -1000,
	276, -24, -1000, -1000, 70, 401, 401, 401, 401, 401,
	401, 401, 197, 106, -25, 192, 180, 179, 233, 127,
	284, -1000, -1000, -1000, -1000, 162, -1000, -1000, 270, 401,
	401, 274, -1000, 387, -26, 457, 232, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 401, 401, 401, 401, 401,
	401, 401, -30, 261, 173, -1000, 122, -1000, 232, -1000,
	626, 232, 626, 626, 626, 626, 613, 600, 79, -32,
	-1000, -1000, -1000, -1000, 208, 164, 146, 136, 134, -38,
	139, 135, 124, 137, 123, 201, 401, 326, 133, 129,
	118, 25, 84, 215, -39, 626, 626, 58, -1000, 401,
	83, 65, 401, 401, 401, 401, 401, 401, 401, 401,
	401, 401, 401, 401, 401, -1000, -1000, -1000, -1000, 309,
	309, 528, 199, 163, -40, 401, 40, 326, 197, 197,
	-1000, 19, 401, 401, -1000, -1000, -1000, -1000, 401, -1000,
	401, -1000, 401, 401, -1000, 401, 401, 571, -1000, 626,
	-1000, -1000, -54, 401, 401, -1000, 401, -1000, 401, 401,
	-1000, 401, -9, 192, 180, 162, 179, -1000, -1000, -13,
	-1000, -56, -1000, 326, 238, 233, 309, 309, 245, 245,
	245, -1000, -1000, -1000, -1000, 445, 445, 382, 639, -1000,
	401, 401, -1000, 482, 255, 111, -1000, -57, -1000, 41,
	75, 102, 626, 97, 626, 626, 626, 626, 626, 550,
	-1000, 160, 91, 626, 626, 626, 626, 626, 215, 1,
	58, -1000, -15, 197, 96, 21, 401, 197, 246, 58,
	-1000, 197, -1000, -1000, 401, -1000, -1000, 401, -1000, 43,
	-1000, -59, -1000, 401, 42, -1000, -1000, 506, 72, 197,
	-1000, 31, 626, 16, -1000, 159, 432, -1000, -1000, -1000,
	47, -1000, -1000, 65, -1000, 169, 197, 39, -1000,
}
var parserPgo = []int{

	0, 323, 322, 319, 8, 318, 269, 2, 7, 270,
	0, 4, 317, 1, 305, 3, 5, 301, 64, 299,
	290, 27, 42, 26, 24, 6, 14, 23,
}
var parserR1 = []int{

	0, 1, 2, 2, 3, 14, 14, 15, 4, 4,
	5, 5, 12, 12, 13, 9, 9, 9, 9, 9,
	27, 8, 8, 8, 8, 7, 7, 7, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, 6, 6, 6, 6, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 17, 11, 11,
	11, 18, 19, 19, 20, 16, 16, 24, 25, 25,
	25, 25, 25, 22, 22, 22, 22, 21, 21, 21,
	21, 23, 23, 23,
}
var parserR2 = []int{

	0, 5, 2, 0, 6, 3, 1, 2, 2, 0,
	7, 8, 3, 1, 2, 1, 1, 1, 1, 3,
	3, 1, 1, 1, 5, 3, 1, 12, 1, 4,
	1, 2, 2, 2, 2, 2, 2, 7, 7, 5,
	3, 2, 2, 2, 2, 5, 5, 3, 4, 4,
	4, 4, 4, 3, 3, 3, 4, 4, 4, 4,
	4, 3, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 2, 2, 2, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 5, 3, 6, 3, 3, 1,
	0, 2, 4, 3, 1, 2, 2, 6, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 3,
}
var parserChk = []int{

	-1000, -1, 4, -2, -4, -3, 6, -7, -5, -6,
	67, -22, 13, -26, 14, 15, 16, 17, 18, 19,
	20, 24, 4, 2, 28, -21, -23, -24, 69, -9,
	-27, 31, 32, 33, 34, 35, -18, -16, 11, 29,
	30, 69, 5, 65, -22, -10, 69, 70, 56, 57,
	71, 68, -27, -20, -18, 36, 38, 39, 40, 45,
	44, 60, 28, 11, 27, 58, 69, -9, 69, -27,
	-10, 69, -10, -10, -10, -10, -10, -10, -7, -22,
	65, 5, 23, 26, 69, -27, 59, 59, 59, 10,
	44, 45, 42, 41, 43, -19, 59, 63, 44, 45,
	42, 41, 43, 60, 10, -10, -10, 7, -6, 67,
	69, 65, 44, 45, 41, 43, 42, 50, 48, 51,
	49, 52, 53, 46, 47, -10, -10, -10, -10, -10,
	-10, -10, 69, -27, 10, 60, 60, 63, 21, 25,
	5, 69, 60, 60, 61, 61, 61, 69, 63, 44,
	63, 45, 63, 63, 41, 63, 59, -10, -8, -10,
	-17, -16, 9, 59, 63, 44, 63, 45, 63, 63,
	41, 63, -25, -21, -23, 35, -24, 69, 69, -14,
	-15, -22, 69, 63, -26, 69, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, 62,
	60, 60, 69, -10, 62, -12, -13, -22, -8, -7,
	-7, -11, -10, -11, -10, -10, -10, -10, -10, -10,
	61, 69, -11, -10, -10, -10, -10, -10, 64, -4,
	64, 69, -8, 25, -11, -11, 64, 12, 62, 64,
	69, 22, 26, 62, 64, 62, 61, 60, 61, -25,
	8, -22, -15, 65, -7, 62, 62, -10, -7, 12,
	-13, -7, -10, -11, 62, 69, -10, 26, 62, 5,
	-7, 23, 62, 65, 5, -26, 25, -7, 26,
}
var parserDef = []int{

	0, -2, 3, 9, 0, 2, 0, 0, 8, 26,
	0, 0, 28, 30, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 113, 114, 115, -2, 0,
	18, 117, 118, 119, 120, 0, 16, 17, 0, 0,
	0, 0, 1, 0, 0, 0, -2, 64, 65, 66,
	67, 68, 69, 71, 72, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 104, 0, 31, 15, 18,
	32, 70, 33, 34, 35, 36, 0, 0, 0, 0,
	41, 42, 43, 44, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 105, 106, 0, 25, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 73, 74, 75, 76, 77,
	78, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	40, 0, 100, 100, 121, 123, 122, 20, 0, 53,
	0, 54, 0, 0, 55, 0, 0, 0, 47, 21,
	22, 23, 0, 100, 0, 61, 0, 62, 0, 0,
	63, 0, 0, 108, 109, 110, 111, 112, 19, 9,
	6, 0, 116, 0, 0, 15, 79, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 91, 92,
	100, 100, 95, 0, 0, 0, 13, 0, 29, 0,
	0, 0, 99, 0, 48, 49, 50, 51, 52, 0,
	103, 0, 0, 56, 57, 58, 59, 60, 0, 0,
	0, 7, 0, 0, 0, 0, 0, 0, 0, 0,
	14, 0, 39, 45, 0, 46, 102, 100, 97, 0,
	4, 0, 5, 0, 0, 93, 94, 0, 0, 0,
	12, 0, 98, 0, 107, 0, 0, 38, 96, 10,
	0, 37, 24, 0, 11, 0, 0, 0, 27,
}
var parserTok1 = []int{

	1,
}
var parserTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
}
var parserTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var parserDebug = 0

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

const parserFlag = -1000

func parserTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(parserToknames) {
		if parserToknames[c-4] != "" {
			return parserToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserlex1(lex parserLexer, lval *parserSymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		c = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			c = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		c = parserTok3[i+0]
		if c == char {
			c = parserTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(c), uint(char))
	}
	return c
}

func parserParse(parserlex parserLexer) int {
	var parsern int
	var parserlval parserSymType
	var parserVAL parserSymType
	parserS := make([]parserSymType, parserMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserchar := -1
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parserchar), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserchar < 0 {
		parserchar = parserlex1(parserlex, &parserlval)
	}
	parsern += parserchar
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parserchar { /* valid shift */
		parserchar = -1
		parserVAL = parserlval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserchar < 0 {
			parserchar = parserlex1(parserlex, &parserlval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parserchar {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error("syntax error")
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parserchar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parserchar))
			}
			if parserchar == parserEofCode {
				goto ret1
			}
			parserchar = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		//line waccparser.y:110
		{
	                                                  parserlex.(*Lexer).prog = &Program{ClassList : parserS[parserpt-3].classes , FunctionList : parserS[parserpt-2].functions , StatList : parserS[parserpt-1].stmts , SymbolTable : NewInstance(), FileText :&parserlex.(*Lexer).input}
	                                                 }
	case 2:
		//line waccparser.y:115
		{ parserVAL.classes = append(parserS[parserpt-1].classes, parserS[parserpt-0].class)}
	case 3:
		//line waccparser.y:116
		{ parserVAL.classes = []*Class{} }
	case 4:
		//line waccparser.y:119
		{ if !checkClassIdent(parserS[parserpt-4].ident) {
	                                                         	parserlex.Error("Invalid class name")
	                                                     }
	                                                     parserVAL.class = &Class{ Pos : parserS[parserpt-5].pos, FileText :&parserlex.(*Lexer).input, Ident : ClassType(parserS[parserpt-4].ident), FieldList : parserS[parserpt-2].fields , FunctionList : parserS[parserpt-1].functions}
	                                                   }
	case 5:
		//line waccparser.y:125
		{ parserVAL.fields = append(parserS[parserpt-2].fields, parserS[parserpt-0].field)}
	case 6:
		//line waccparser.y:126
		{ parserVAL.fields = []Field{ parserS[parserpt-0].field } }
	case 7:
		//line waccparser.y:129
		{ parserVAL.field = Field{FieldType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 8:
		//line waccparser.y:131
		{ parserVAL.functions = append(parserS[parserpt-1].functions, parserS[parserpt-0].function)}
	case 9:
		//line waccparser.y:132
		{ parserVAL.functions = []*Function{} }
	case 10:
		//line waccparser.y:135
		{ if !checkStats(parserS[parserpt-1].stmts) {
	          	parserlex.Error("Missing return statement")
	           }
	             parserVAL.function = &Function{Ident : parserS[parserpt-5].ident, ReturnType : parserS[parserpt-6].typedefinition, StatList : parserS[parserpt-1].stmts, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 11:
		//line waccparser.y:141
		{ if !checkStats(parserS[parserpt-1].stmts) {
	            	parserlex.Error("Missing return statement")
	            }
	             parserVAL.function = &Function{Ident : parserS[parserpt-6].ident, ReturnType : parserS[parserpt-7].typedefinition, StatList : parserS[parserpt-1].stmts, ParameterTypes : parserS[parserpt-4].params, SymbolTable: NewInstance(), FileText :&parserlex.(*Lexer).input}
	           }
	case 12:
		//line waccparser.y:147
		{ parserVAL.params = append(parserS[parserpt-2].params, parserS[parserpt-0].param)}
	case 13:
		//line waccparser.y:148
		{ parserVAL.params = []Param{ parserS[parserpt-0].param } }
	case 14:
		//line waccparser.y:150
		{ parserVAL.param = Param{ParamType : parserS[parserpt-1].typedefinition, Ident : parserS[parserpt-0].ident} }
	case 15:
		//line waccparser.y:152
		{parserVAL.assignlhs = parserS[parserpt-0].ident}
	case 16:
		//line waccparser.y:153
		{parserVAL.assignlhs = parserS[parserpt-0].arrayelem}
	case 17:
		//line waccparser.y:154
		{parserVAL.assignlhs = parserS[parserpt-0].pairelem}
	case 18:
		//line waccparser.y:155
		{ parserVAL.assignlhs = parserS[parserpt-0].fieldaccess}
	case 19:
		//line waccparser.y:156
		{ parserVAL.assignlhs = ThisInstance{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-0].ident} }
	case 20:
		//line waccparser.y:158
		{parserVAL.fieldaccess = FieldAccess{ &parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-2].ident, parserS[parserpt-0].ident, } }
	case 21:
		//line waccparser.y:160
		{parserVAL.assignrhs = parserS[parserpt-0].expr}
	case 22:
		//line waccparser.y:161
		{parserVAL.assignrhs = parserS[parserpt-0].arrayliter}
	case 23:
		//line waccparser.y:162
		{parserVAL.assignrhs = parserS[parserpt-0].pairelem}
	case 24:
		//line waccparser.y:163
		{ parserVAL.assignrhs = NewObject{Class : ClassType(parserS[parserpt-3].ident) , Init : parserS[parserpt-1].exprs , Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input}}
	case 25:
		//line waccparser.y:165
		{ parserVAL.stmts = append(parserS[parserpt-2].stmts,parserS[parserpt-0].stmt)   }
	case 26:
		//line waccparser.y:166
		{ parserVAL.stmts = []Statement{parserS[parserpt-0].stmt} }
	case 27:
		//line waccparser.y:168
		{
	                                                                                                                 stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                                                                                 w := While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-11].pos, FileText :&parserlex.(*Lexer).input}
	                                                                                                                 d := Declare{DecType : parserS[parserpt-10].typedefinition, Lhs : parserS[parserpt-9].ident, Rhs : parserS[parserpt-7].assignrhs, Pos : parserS[parserpt-11].pos ,FileText :&parserlex.(*Lexer).input }
	                                                                                                                 parserVAL.stmts = []Statement{d,w}
	                                                                                                                }
	case 28:
		//line waccparser.y:175
		{ parserVAL.stmt = Skip{Pos : parserS[parserpt-0].pos ,FileText :&parserlex.(*Lexer).input } }
	case 29:
		//line waccparser.y:176
		{ parserVAL.stmt = Declare{DecType : parserS[parserpt-3].typedefinition, Lhs : parserS[parserpt-2].ident, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input } }
	case 30:
		//line waccparser.y:177
		{ parserVAL.stmt = parserS[parserpt-0].stmt }
	case 31:
		//line waccparser.y:178
		{ parserVAL.stmt = Read{ &parserlex.(*Lexer).input, parserS[parserpt-1].pos , parserS[parserpt-0].assignlhs, } }
	case 32:
		//line waccparser.y:179
		{ parserVAL.stmt = Free{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 33:
		//line waccparser.y:180
		{ parserVAL.stmt = Return{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 34:
		//line waccparser.y:181
		{ parserVAL.stmt = Exit{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 35:
		//line waccparser.y:182
		{ parserVAL.stmt = Print{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 36:
		//line waccparser.y:183
		{ parserVAL.stmt = Println{&parserlex.(*Lexer).input, parserS[parserpt-1].pos, parserS[parserpt-0].expr} }
	case 37:
		//line waccparser.y:184
		{ parserVAL.stmt = If{Conditional : parserS[parserpt-5].expr, ThenStat : parserS[parserpt-3].stmts, ElseStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input } }
	case 38:
		//line waccparser.y:185
		{
	                                                              stats := append(parserS[parserpt-1].stmts, parserS[parserpt-3].stmt)
	                                                              parserVAL.stmt = While{Conditional : parserS[parserpt-5].expr, DoStat : stats, Pos : parserS[parserpt-6].pos, FileText :&parserlex.(*Lexer).input}
	                                                             }
	case 39:
		//line waccparser.y:189
		{ parserVAL.stmt = While{Conditional : parserS[parserpt-3].expr, DoStat : parserS[parserpt-1].stmts, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input} }
	case 40:
		//line waccparser.y:190
		{ parserVAL.stmt = Scope{StatList : parserS[parserpt-1].stmts, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input } }
	case 41:
		//line waccparser.y:191
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 42:
		//line waccparser.y:195
		{ parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 43:
		//line waccparser.y:198
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 44:
		//line waccparser.y:202
		{
	                                                          parserlex.Error("Syntax error : Invalid statement")
	                                                          parserVAL.stmt = nil
	                                                        }
	case 45:
		//line waccparser.y:206
		{ parserVAL.stmt = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 46:
		//line waccparser.y:207
		{ parserVAL.stmt = CallInstance{Class : (parserS[parserpt-3].fieldaccess.(FieldAccess)).ObjectName, Func: (parserS[parserpt-3].fieldaccess.(FieldAccess)).Field, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 47:
		//line waccparser.y:209
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].assignlhs, Rhs : parserS[parserpt-0].assignrhs, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 48:
		//line waccparser.y:210
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 49:
		//line waccparser.y:211
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 50:
		//line waccparser.y:212
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 51:
		//line waccparser.y:213
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 52:
		//line waccparser.y:214
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].ident, Rhs : Binop{Left : parserS[parserpt-3].ident, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 53:
		//line waccparser.y:215
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 54:
		//line waccparser.y:216
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 55:
		//line waccparser.y:217
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].ident, Rhs : Binop{Left : parserS[parserpt-2].ident, Binary : MUL,  Right : parserS[parserpt-2].ident,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 56:
		//line waccparser.y:219
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 57:
		//line waccparser.y:220
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : SUB , Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 58:
		//line waccparser.y:221
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 59:
		//line waccparser.y:222
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 60:
		//line waccparser.y:223
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-3].fieldaccess, Rhs : Binop{Left : parserS[parserpt-3].fieldaccess, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-3].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-3].pos ,FileText :&parserlex.(*Lexer).input} }
	case 61:
		//line waccparser.y:224
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : PLUS, Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 62:
		//line waccparser.y:225
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : SUB,  Right : Integer(1), Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 63:
		//line waccparser.y:226
		{ parserVAL.stmt = Assignment{Lhs : parserS[parserpt-2].fieldaccess, Rhs : Binop{Left : parserS[parserpt-2].fieldaccess, Binary : MUL,  Right : parserS[parserpt-2].fieldaccess,         Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input}, Pos : parserS[parserpt-2].pos ,FileText :&parserlex.(*Lexer).input} }
	case 64:
		//line waccparser.y:228
		{ parserVAL.expr =  parserS[parserpt-0].integer }
	case 65:
		//line waccparser.y:229
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 66:
		//line waccparser.y:230
		{ parserVAL.expr =  parserS[parserpt-0].boolean }
	case 67:
		//line waccparser.y:231
		{ parserVAL.expr =  parserS[parserpt-0].character }
	case 68:
		//line waccparser.y:232
		{ parserVAL.expr =  parserS[parserpt-0].stringconst }
	case 69:
		//line waccparser.y:233
		{ parserVAL.expr = parserS[parserpt-0].fieldaccess }
	case 70:
		//line waccparser.y:234
		{ parserVAL.expr = parserS[parserpt-0].ident}
	case 71:
		//line waccparser.y:235
		{ parserVAL.expr =  parserS[parserpt-0].pairliter }
	case 72:
		//line waccparser.y:236
		{ parserVAL.expr =  parserS[parserpt-0].arrayelem }
	case 73:
		//line waccparser.y:237
		{ parserVAL.expr = Unop{Unary : NOT, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 74:
		//line waccparser.y:238
		{ parserVAL.expr = Unop{Unary : LEN, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 75:
		//line waccparser.y:239
		{ parserVAL.expr = Unop{Unary : ORD, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 76:
		//line waccparser.y:240
		{ parserVAL.expr = Unop{Unary : CHR, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 77:
		//line waccparser.y:241
		{ parserVAL.expr = Unop{Unary : SUB, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos, FileText :&parserlex.(*Lexer).input  } }
	case 78:
		//line waccparser.y:242
		{ parserVAL.expr = parserS[parserpt-0].expr }
	case 79:
		//line waccparser.y:243
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : PLUS, Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 80:
		//line waccparser.y:244
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : SUB,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 81:
		//line waccparser.y:245
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MUL,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 82:
		//line waccparser.y:246
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : MOD,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 83:
		//line waccparser.y:247
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : DIV,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 84:
		//line waccparser.y:248
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 85:
		//line waccparser.y:249
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GT,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 86:
		//line waccparser.y:250
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : LTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 87:
		//line waccparser.y:251
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : GTE,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 88:
		//line waccparser.y:252
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : EQ,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 89:
		//line waccparser.y:253
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : NEQ,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 90:
		//line waccparser.y:254
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : AND,  Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 91:
		//line waccparser.y:255
		{ parserVAL.expr = Binop{Left : parserS[parserpt-2].expr, Binary : OR,   Right : parserS[parserpt-0].expr, Pos : parserS[parserpt-2].pos, FileText :&parserlex.(*Lexer).input  } }
	case 92:
		//line waccparser.y:256
		{ parserVAL.expr = parserS[parserpt-1].expr }
	case 93:
		//line waccparser.y:257
		{ parserVAL.expr = Call{Ident : parserS[parserpt-3].ident, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 94:
		//line waccparser.y:258
		{ parserVAL.expr = CallInstance{Class : (parserS[parserpt-3].fieldaccess.(FieldAccess)).ObjectName, Func: (parserS[parserpt-3].fieldaccess.(FieldAccess)).Field, ParamList : parserS[parserpt-1].exprs, Pos : parserS[parserpt-4].pos, FileText :&parserlex.(*Lexer).input  } }
	case 95:
		//line waccparser.y:259
		{ parserVAL.expr = ThisInstance{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-0].ident} }
	case 96:
		//line waccparser.y:260
		{ parserVAL.expr = NewPair{FstExpr : parserS[parserpt-3].expr, SndExpr : parserS[parserpt-1].expr, Pos : parserS[parserpt-5].pos, FileText :&parserlex.(*Lexer).input } }
	case 97:
		//line waccparser.y:262
		{ parserVAL.arrayliter = ArrayLiter{&parserlex.(*Lexer).input, parserS[parserpt-2].pos, parserS[parserpt-1].exprs } }
	case 98:
		//line waccparser.y:264
		{parserVAL.exprs = append(parserS[parserpt-2].exprs, parserS[parserpt-0].expr)}
	case 99:
		//line waccparser.y:265
		{parserVAL.exprs = []Evaluation{parserS[parserpt-0].expr}}
	case 100:
		//line waccparser.y:266
		{parserVAL.exprs = []Evaluation{}}
	case 101:
		//line waccparser.y:268
		{parserVAL.arrayelem = ArrayElem{Ident: parserS[parserpt-1].ident, Exprs : parserS[parserpt-0].exprs, Pos : parserS[parserpt-1].pos,FileText :&parserlex.(*Lexer).input  } }
	case 102:
		//line waccparser.y:270
		{parserVAL.exprs = append(parserS[parserpt-3].exprs, parserS[parserpt-1].expr)}
	case 103:
		//line waccparser.y:271
		{parserVAL.exprs = []Evaluation{parserS[parserpt-1].expr}}
	case 104:
		//line waccparser.y:273
		{ parserVAL.pairliter =  PairLiter{} }
	case 105:
		//line waccparser.y:275
		{ parserVAL.pairelem = PairElem{Fsnd: Fst, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 106:
		//line waccparser.y:276
		{ parserVAL.pairelem = PairElem{Fsnd: Snd, Expr : parserS[parserpt-0].expr, Pos : parserS[parserpt-1].pos  } }
	case 107:
		//line waccparser.y:278
		{ parserVAL.typedefinition = PairType{FstType : parserS[parserpt-3].pairelemtype, SndType : parserS[parserpt-1].pairelemtype} }
	case 108:
		//line waccparser.y:280
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 109:
		//line waccparser.y:281
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition }
	case 110:
		//line waccparser.y:282
		{ parserVAL.pairelemtype = Pair}
	case 111:
		//line waccparser.y:283
		{ parserVAL.pairelemtype = parserS[parserpt-0].typedefinition}
	case 112:
		//line waccparser.y:284
		{ parserVAL.pairelemtype = ClassType(parserS[parserpt-0].ident)}
	case 113:
		//line waccparser.y:286
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 114:
		//line waccparser.y:287
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 115:
		//line waccparser.y:288
		{ parserVAL.typedefinition =  parserS[parserpt-0].typedefinition }
	case 116:
		//line waccparser.y:289
		{ parserVAL.typedefinition = ClassType(parserS[parserpt-0].ident) }
	case 117:
		//line waccparser.y:291
		{ parserVAL.typedefinition =  Int }
	case 118:
		//line waccparser.y:292
		{ parserVAL.typedefinition =  Bool }
	case 119:
		//line waccparser.y:293
		{ parserVAL.typedefinition =  Char }
	case 120:
		//line waccparser.y:294
		{ parserVAL.typedefinition =  String }
	case 121:
		//line waccparser.y:296
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 122:
		//line waccparser.y:297
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	case 123:
		//line waccparser.y:298
		{ parserVAL.typedefinition = ArrayType{Type : parserS[parserpt-2].typedefinition} }
	}
	goto parserstack /* stack new state and value */
}
