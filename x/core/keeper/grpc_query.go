package keeper

import (
	"github.com/athlete-labs/core/x/core/types"
)

var _ types.QueryServer = Keeper{}
