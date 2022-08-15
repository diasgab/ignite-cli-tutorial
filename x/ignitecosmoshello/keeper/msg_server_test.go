package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "ignite-cosmos-hello/testutil/keeper"
	"ignite-cosmos-hello/x/ignitecosmoshello/keeper"
	"ignite-cosmos-hello/x/ignitecosmoshello/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgnitecosmoshelloKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
