package bind

import (
	"context"

	"github.com/trilitech/tzgo/contract"
	"github.com/trilitech/tzgo/micheline"
	"github.com/trilitech/tzgo/rpc"
	"github.com/trilitech/tzgo/tezos"
)

type Contract interface {
	Address() tezos.Address
	Call(ctx context.Context, args contract.CallArguments, opts *rpc.CallOptions) (*rpc.Receipt, error)
	RunView(ctx context.Context, name string, args micheline.Prim) (micheline.Prim, error)
}

type RPC interface {
	GetContractStorage(ctx context.Context, addr tezos.Address, id rpc.BlockID) (micheline.Prim, error)
	GetBigmapValue(ctx context.Context, bigmap int64, hash tezos.ExprHash, id rpc.BlockID) (micheline.Prim, error)
}

var (
	_ Contract = &contract.Contract{}
	_ RPC      = &rpc.Client{}
)
