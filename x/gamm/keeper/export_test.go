package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

<<<<<<< HEAD
	"github.com/osmosis-labs/osmosis/v10/x/gamm/types"
=======
	"github.com/osmosis-labs/osmosis/v11/x/gamm/types"
>>>>>>> ccc5afe (Add function)
)

// SetParams sets the total set of params.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.setParams(ctx, params)
}

// SetParams sets the total set of params.
func (k Keeper) SetParams2(ctx sdk.Context, params types.Params) {
	k.setParams(ctx, params)
}
