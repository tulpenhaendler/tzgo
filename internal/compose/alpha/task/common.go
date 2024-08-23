// Copyright (c) 2023 Blockwatch Data Inc.
// Author: alex@blockwatch.cc, abdul@blockwatch.cc

package task

import (
	"github.com/tulpenhaendler/tzgo/internal/compose"
	"github.com/tulpenhaendler/tzgo/internal/compose/alpha"
	"github.com/tulpenhaendler/tzgo/tezos"

	"github.com/pkg/errors"
)

type BaseTask struct {
	Source tezos.Address
	Key    tezos.PrivateKey
}

func (t *BaseTask) parse(ctx compose.Context, task alpha.Task) (err error) {
	if task.Source != "" {
		if t.Source, err = ctx.ResolveAddress(task.Source); err != nil {
			err = errors.Wrap(err, "source")
			return
		}
		if t.Key, err = ctx.ResolvePrivateKey(task.Source); err != nil {
			err = errors.Wrap(err, "key")
			return
		}
	} else {
		t.Source = ctx.BaseAccount.Address
		t.Key = ctx.BaseAccount.PrivateKey
	}
	return
}

type TargetTask struct {
	BaseTask
	Destination tezos.Address
}

func (t *TargetTask) parse(ctx compose.Context, task alpha.Task) (err error) {
	if err = t.BaseTask.parse(ctx, task); err != nil {
		return
	}
	if t.Destination, err = ctx.ResolveAddress(task.Destination); err != nil {
		err = errors.Wrap(err, "destination")
		return
	}
	return
}
