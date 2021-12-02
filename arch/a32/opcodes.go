package a32

import (
	"fmt"

	"github.com/rj45/nanogo/codegen/asm"
	"github.com/rj45/nanogo/ir/op"
)

type Opcode int

//go:generate enumer -type=Opcode

const (
	// Natively implemented instructions
	NOP Opcode = iota
	BRK
	HLT
	ERR

	ADD
	SUB
	ADDC
	SUBB
	AND
	OR
	XOR
	SHL
	ASR
	LSR

	LD
	ST
	LD8
	ST8
	LD16
	ST16

	BR_E
	BR_NE
	BR_U_L
	BR_U_LE
	BR_U_GE
	BR_U_G
	BR_S_L
	BR_S_LE
	BR_S_GE
	BR_S_G
	BRA

	CALL
	RET

	CMP
	NEG
	NEGB
	NOT

	MOV
	SWP

	NumOps
)

func (op Opcode) Fmt() asm.Fmt {
	return opDefs[op].fmt
}

func (op Opcode) IsMove() bool {
	return op == MOV || op == SWP
}

func (op Opcode) IsCall() bool {
	return op == CALL
}

type def struct {
	fmt Fmt
	op  op.Op
}

var opDefs = [...]def{
	NOP:     {fmt: NoFmt},
	BRK:     {fmt: NoFmt},
	HLT:     {fmt: NoFmt},
	ERR:     {fmt: NoFmt},
	ADD:     {fmt: BinaryFmt, op: op.Add},
	SUB:     {fmt: BinaryFmt, op: op.Sub},
	ADDC:    {fmt: BinaryFmt},
	SUBB:    {fmt: BinaryFmt},
	AND:     {fmt: BinaryFmt, op: op.And},
	OR:      {fmt: BinaryFmt, op: op.Or},
	XOR:     {fmt: BinaryFmt, op: op.Xor},
	SHL:     {fmt: BinaryFmt, op: op.ShiftLeft},
	ASR:     {fmt: BinaryFmt},
	LSR:     {fmt: BinaryFmt, op: op.ShiftRight},
	LD:      {fmt: LoadFmt, op: op.Load},
	ST:      {fmt: StoreFmt, op: op.Store},
	LD8:     {fmt: LoadFmt},
	ST8:     {fmt: StoreFmt},
	LD16:    {fmt: LoadFmt},
	ST16:    {fmt: StoreFmt},
	BR_E:    {fmt: CallFmt},
	BR_NE:   {fmt: CallFmt},
	BR_U_L:  {fmt: CallFmt},
	BR_U_LE: {fmt: CallFmt},
	BR_U_GE: {fmt: CallFmt},
	BR_U_G:  {fmt: CallFmt},
	BR_S_L:  {fmt: CallFmt},
	BR_S_LE: {fmt: CallFmt},
	BR_S_GE: {fmt: CallFmt},
	BR_S_G:  {fmt: CallFmt},
	BRA:     {fmt: CallFmt},
	CALL:    {fmt: CallFmt, op: op.Call},
	RET:     {fmt: NoFmt},
	CMP:     {fmt: CompareFmt},
	NEG:     {fmt: UnaryFmt, op: op.Negate},
	NEGB:    {fmt: UnaryFmt},
	NOT:     {fmt: UnaryFmt, op: op.Invert},
	MOV:     {fmt: MoveFmt, op: op.Copy},
	SWP:     {fmt: CompareFmt, op: op.SwapIn},
}

var translations []Opcode

func init() {
	translations = make([]Opcode, op.NumOps)
	for i := NOP; i < NumOps; i++ {
		if opDefs[i].fmt == BadFmt {
			panic(fmt.Sprintf("missing opDef for %s", i))
		}
		translations[opDefs[i].op] = i
	}
}
