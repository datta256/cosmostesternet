package keeper

import (
	"testernet/x/identity/types"
)

var _ types.QueryServer = Keeper{}
