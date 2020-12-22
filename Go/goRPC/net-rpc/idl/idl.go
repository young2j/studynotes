package idl

import "errors"

// Args ...
type Args struct {
	A, B int
}

//Return ...
type Return struct {
	Quo, Rem int
}

// Arith ...
type Arith int


// Multiply ...
func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide ...
func (a *Arith) Divide(args *Args, ret *Return) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	ret.Quo = args.A / args.B
	ret.Rem = args.A % args.B
	return nil
}
