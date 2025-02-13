package argon

import (
	cctptypes "github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	routertypes "github.com/strangelove-ventures/noble-router/x/router/types"
)

func CreateStoreLoader(upgradeHeight int64) baseapp.StoreLoader {
	storeUpgrades := storeTypes.StoreUpgrades{
		Added: []string{
			cctptypes.ModuleName, routertypes.ModuleName,
		},
	}

	return upgradeTypes.UpgradeStoreLoader(upgradeHeight, &storeUpgrades)
}
