package keeper_test

import (
	"testing"

	testkeeper "github.com/athlete-labs/core/testutil/keeper"
	"github.com/athlete-labs/core/x/core/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CoreKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
