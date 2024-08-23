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

var _ alpha.TaskBuilder = (*UndelegateTask)(nil)

func init() {
	alpha.RegisterTask("undelegate", NewUndelegateTask)
}

type UndelegateTask struct {
	BaseTask
}

func NewUndelegateTask() alpha.TaskBuilder {
	return &UndelegateTask{}
}

func (t *UndelegateTask) Type() string {
	return "undelegate"
}

func (t *UndelegateTask) Build(ctx compose.Context, task alpha.Task) (*codec.Op, *rpc.CallOptions, error) {
	if err := t.parse(ctx, task); err != nil {
		return nil, nil, errors.Wrap(err, "parse")
	}
	opts := rpc.NewCallOptions()
	opts.Signer = signer.NewFromKey(t.Key)
	opts.IgnoreLimits = true
	op := codec.NewOp().
		WithSource(t.Source).
		WithUndelegation().
		WithLimits([]tezos.Limits{rpc.DefaultDelegationLimitsEOA}, 0)
	return op, opts, nil
}

func (t *UndelegateTask) Validate(ctx compose.Context, task alpha.Task) error {
	return t.parse(ctx, task)
}

func (t *UndelegateTask) parse(ctx compose.Context, task alpha.Task) error {
	return t.BaseTask.parse(ctx, task)
}
