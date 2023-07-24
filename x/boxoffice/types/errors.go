package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/boxoffice module sentinel errors
var (
	ErrSample              = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidOwner        = sdkerrors.Register(ModuleName, 1101, "invalid owner address : %s")
	ErrInvalidVolume       = sdkerrors.Register(ModuleName, 1102, "invalid concert volume")
	ErrInvalidConcertIndex = sdkerrors.Register(ModuleName, 1103, "invalid concert index")
	ErrConcertNotFound     = sdkerrors.Register(ModuleName, 1104, "concert not found")
	ErrConcertSoldOut      = sdkerrors.Register(ModuleName, 1105, "concert sold out")
)
