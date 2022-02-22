package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	ErrSample                = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrResourceIdNotFound    = sdkerrors.Register(ModuleName, 1101, "resourceId not found")
	ErrBalanceNotEnough      = sdkerrors.Register(ModuleName, 1102, "balance not enough")
	ErrChainIdNotSupport     = sdkerrors.Register(ModuleName, 1103, "chainId not support")
	ErrUnKnownResourceIdType = sdkerrors.Register(ModuleName, 1104, "unknown resource id type")
)
