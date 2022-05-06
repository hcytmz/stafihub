package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/keeper"
	"github.com/stafihub/stafihub/x/bridge/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.DepositCountList = k.GetDepositCountList(ctx)
	genesis.RelayFeeList = k.GetRelayFeeList(ctx)
	genesis.ChainIdList = k.GetChainIdList(ctx)
	relayFeeReceiver, found := k.GetRelayFeeReceiver(ctx)
	if found {
		genesis.RelayFeeReceiver = relayFeeReceiver.String()
	}

	genesis.ResourceIdToDenom = k.GetAllResourceIdToDenom(ctx)
	genesis.ResourceIdToDenomType = k.GetAllResourceIdDenomTypes(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
