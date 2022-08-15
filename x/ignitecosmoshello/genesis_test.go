package ignitecosmoshello_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ignite-cosmos-hello/testutil/keeper"
	"ignite-cosmos-hello/testutil/nullify"
	"ignite-cosmos-hello/x/ignitecosmoshello"
	"ignite-cosmos-hello/x/ignitecosmoshello/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitecosmoshelloKeeper(t)
	ignitecosmoshello.InitGenesis(ctx, *k, genesisState)
	got := ignitecosmoshello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
