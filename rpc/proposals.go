// Copyright (c) 2020-2021 Blockwatch Data Inc.
// Author: alex@blockwatch.cc

package rpc

import (
	"github.com/tulpenhaendler/tzgo/tezos"
)

// Ensure Proposals implements the TypedOperation interface.
var _ TypedOperation = (*Proposals)(nil)

// Proposals represents a proposal operation
type Proposals struct {
	Generic
	Source    tezos.Address        `json:"source"`
	Period    int                  `json:"period"`
	Proposals []tezos.ProtocolHash `json:"proposals"`
}
