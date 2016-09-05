package codeGeneration

import (
	. "backend/filewriter"
)

type InstructionSelection struct {
	instrs      *ARMList
	code        string
	numTabs     int
	numNewLines int
}
