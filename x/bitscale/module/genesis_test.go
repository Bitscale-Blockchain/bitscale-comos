package bitscale_test

import (
	"testing"

	keepertest "bitscale/testutil/keeper"
	"bitscale/testutil/nullify"
	bitscale "bitscale/x/bitscale/module"
	"bitscale/x/bitscale/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BitscaleKeeper(t)
	bitscale.InitGenesis(ctx, k, genesisState)
	got := bitscale.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
