package keeper

import (
	"github.com/bermuell/boxoffice/x/boxoffice/types"
)

var _ types.QueryServer = Keeper{}
