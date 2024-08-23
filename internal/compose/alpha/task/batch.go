// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc, abdul@blockwatch.cc

package task

import (
	"fmt"

	"github.com/tulpenhaendler/tzgo/codec"
	"github.com/tulpenhaendler/tzgo/internal/compose"
	"github.com/tulpenhaendler/tzgo/internal/compose/alpha"
	"github.com/tulpenhaendler/tzgo/rpc"
	"github.com/tulpenhaendler/tzgo/signer"

	"github.com/pkg/errors"
)

var _ alpha.TaskBuilder = (*BatchTask)(nil)

func init() {
	alpha.RegisterTask("batch", NewBatchTask)
}

type BatchTask struct {
	BaseTask
}

func NewBatchTask() alpha.TaskBuilder {
	return &BatchTask{}
}

func (t *BatchTask) Type() string {
	return "batch"
}

func (t *BatchTask) Build(ctx compose.Context, task alpha.Task) (*codec.Op, *rpc.CallOptions, error) {
	if err := t.parse(ctx, task); err != nil {
		return nil, nil, errors.Wrap(err, "parse")
	}
	opts := rpc.NewCallOptions()
	opts.Signer = signer.NewFromKey(t.Key)
	op := codec.NewOp().WithSource(t.Source)
	for i, ct := range task.Contents {
		// use common source
		ct.Source = task.Source
		childTask, err := alpha.NewTask(ct.Type)
		if err != nil {
			return nil, nil, fmt.Errorf("batch[%d]: %v", i, err)
		}
		childOp, _, err := childTask.Build(ctx, ct)
		if err != nil {
			return nil, nil, fmt.Errorf("batch[%d]: %v", i, err)
		}
		for _, o := range childOp.Contents {
			op.WithContents(o)
		}
	}
	return op, opts, nil
}

func (t *BatchTask) Validate(ctx compose.Context, task alpha.Task) error {
	return t.parse(ctx, task)
}

func (t *BatchTask) parse(ctx compose.Context, task alpha.Task) (err error) {
	if err = t.BaseTask.parse(ctx, task); err != nil {
		return err
	}
	return
}
