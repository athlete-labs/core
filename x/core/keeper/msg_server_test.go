package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/athlete-labs/core/testutil/keeper"
	"github.com/athlete-labs/core/x/core/keeper"
	"github.com/athlete-labs/core/x/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CoreKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
