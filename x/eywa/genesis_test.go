package eywa_test

import (
	"testing"

	keepertest "eywa/testutil/keeper"
	"eywa/testutil/nullify"
	"eywa/x/eywa"
	"eywa/x/eywa/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EywaKeeper(t)
	eywa.InitGenesis(ctx, *k, genesisState)
	got := eywa.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
