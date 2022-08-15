package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ignite-cosmos-hello/testutil/keeper"
	"ignite-cosmos-hello/x/ignitecosmoshello/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgnitecosmoshelloKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
