// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc, abdul@blockwatch.cc

package task

import (
	"github.com/tulpenhaendler/tzgo/codec"
	"github.com/tulpenhaendler/tzgo/internal/compose"
	"github.com/tulpenhaendler/tzgo/internal/compose/alpha"
	"github.com/tulpenhaendler/tzgo/rpc"
	"github.com/tulpenhaendler/tzgo/signer"
	"github.com/tulpenhaendler/tzgo/tezos"

	"github.com/pkg/errors"
)

var _ alpha.TaskBuilder = (*DelegateTask)(nil)

func init() {
	alpha.RegisterTask("delegate", NewDelegateTask)
}

type DelegateTask struct {
	TargetTask
}

func NewDelegateTask() alpha.TaskBuilder {
	return &DelegateTask{}
}

func (t *DelegateTask) Type() string {
	return "delegate"
}

func (t *DelegateTask) Build(ctx compose.Context, task alpha.Task) (*codec.Op, *rpc.CallOptions, error) {
	if err := t.parse(ctx, task); err != nil {
		return nil, nil, errors.Wrap(err, "parse")
	}
	opts := rpc.NewCallOptions()
	opts.Signer = signer.NewFromKey(t.Key)
	opts.IgnoreLimits = true
	op := codec.NewOp().
		WithSource(t.Source).
		WithDelegation(t.Destination).
		WithLimits([]tezos.Limits{rpc.DefaultDelegationLimitsEOA}, 0)
	return op, opts, nil
}

func (t *DelegateTask) Validate(ctx compose.Context, task alpha.Task) error {
	return t.parse(ctx, task)
}

func (t *DelegateTask) parse(ctx compose.Context, task alpha.Task) error {
	return t.TargetTask.parse(ctx, task)
}
