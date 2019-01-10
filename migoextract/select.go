package migoextract

import (
	"github.com/nickng/migo/v3"
	"golang.org/x/tools/go/ssa"
)

// Select keeps track of select statement and its branches.
type Select struct {
	Instr    *ssa.Select           // Select SSA instruction.
	MigoStmt *migo.SelectStatement // Select statement in MiGo.
	Index    Instance              // Index (extracted from Select instruction).
}
