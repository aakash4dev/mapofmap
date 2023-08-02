package keeper

import (
	"mapofmap/x/mapofmap/types"
)

var _ types.QueryServer = Keeper{}
