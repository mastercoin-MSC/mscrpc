package mscrpc

import (
	"net/rpc"
	"fmt"
)

// Simplistic argument structure
// Validates amount of arguments and throws
// errors if it's unable to satisfy.
type Args struct {
	Vals []interface{}
}
func NewArgs() *Args {
	return &Args{}
}
func (args *Args) Add(v interface{}) *Args {
	args.Vals = append(args.Vals, v)

	// Returns args so we can chain
	return args
}

func (args *Args) Require(num int) error {
	l := len(args.Vals)
	if l != num {
		return fmt.Errorf("invalid args %d for %d", l, num)
	}

	return nil
}

type SimpleSendPackage struct {}
func (p *SimpleSendPackage) ListTxs(args *Args, reply *int) error {
	return nil
}

func (p *SimpleSendPackage) CreateTx(args *Args, reply *string) error {
	// Require 2: receiver, amount (In that order)
	err := args.Require(2)
	if err != nil {
		return err
	}

	*reply = "Test"

	return nil
}

// Registers all RPC packages
func RegisterPackagesRpcPackages() {
	rpc.Register(new(SimpleSendPackage))
}
