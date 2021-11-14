package xform

import (
	"github.com/rj45/nanogo/ir"
)

//go:generate enumer -type=Pass

type Pass int

const (
	Elaboration Pass = iota
	Simplification
	Lowering
	Legalize
	CleanUp
)

var passes [][]func(*ir.Value) int

func addToPass(pass Pass, fn func(*ir.Value) int) int {
	for int(pass) >= len(passes) {
		passes = append(passes, nil)
	}

	passes[pass] = append(passes[pass], fn)
	return 0
}

func Transform(pass Pass, fn *ir.Func) {
	changes := 1
	tries := 0
nextchange:
	for changes > 0 {
		changes = 0
		tries++
		if tries > 1000 {
			panic("too many tries")
		}
		for _, blk := range fn.Blocks() {
			for i := 0; i < blk.NumInstrs(); i++ {
				instr := blk.Instr(i)

				for _, xform := range passes[pass] {
					changes += xform(instr)
					if changes > 0 {
						continue nextchange
					}
				}
			}
		}
	}
}
