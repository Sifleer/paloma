package treasury

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/palomachain/paloma/x/treasury/keeper"
	"github.com/palomachain/paloma/x/treasury/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
