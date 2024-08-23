// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc, abdul@blockwatch.cc

package task

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tulpenhaendler/tzgo/codec"
	"github.com/tulpenhaendler/tzgo/internal/compose"
	"github.com/tulpenhaendler/tzgo/internal/compose/alpha"
	"github.com/tulpenhaendler/tzgo/rpc"
	"github.com/tulpenhaendler/tzgo/signer"
)

var _ alpha.TaskBuilder = (*DeployTask)(nil)

func init() {
	alpha.RegisterTask("deploy", NewDeployTask)
}

type DeployTask struct {
	TargetTask
	Amount int64
}

func NewDeployTask() alpha.TaskBuilder {
	return &DeployTask{}
}

func (c *DeployTask) Type() string {
	return "deploy"
}

func (t *DeployTask) Build(ctx compose.Context, task alpha.Task) (*codec.Op, *rpc.CallOptions, error) {
	if err := t.parse(ctx, task); err != nil {
		return nil, nil, errors.Wrap(err, "parse")
	}
	script, err := alpha.ParseScript(ctx, task)
	if err != nil {
		return nil, nil, errors.Wrap(err, "script")
	}
	opts := rpc.NewCallOptions()
	opts.Signer = signer.NewFromKey(t.Key)
	op := codec.NewOp().
		WithSource(t.Source).
		WithOriginationExt(*script, t.Destination, t.Amount)
	return op, opts, nil
}

func (t *DeployTask) Validate(ctx compose.Context, task alpha.Task) error {
	if err := t.parse(ctx, task); err != nil {
		return err
	}
	if _, err := alpha.ParseScript(ctx, task); err != nil {
		return errors.Wrap(err, "script")
	}
	return nil
}

func (t *DeployTask) parse(ctx compose.Context, task alpha.Task) (err error) {
	if err = t.BaseTask.parse(ctx, task); err != nil {
		return
	}
	t.Amount = int64(task.Amount)
	if task.Alias == "" {
		return fmt.Errorf("alias name required")
	}
	return
}
