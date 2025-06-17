package keeper

import (
	"testernet/x/assetidentity/types"
)

var _ types.QueryServer = Keeper{}
