package idl

import "errors"

// 入参
type Args struct {
	A, B int
}

// 出参
type Return struct {
	Quo, Rem int
}

// 主体-用于实现方法，或者叫过程
type Arith int

// Arith.Multiply
func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Arith.Divide
func (a *Arith) Divide(args *Args, ret *Return) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	ret.Quo = args.A / args.B
	ret.Rem = args.A % args.B
	return nil
}
