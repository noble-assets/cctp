package keeper_test

import (
	"testing"

	"github.com/circlefin/noble-cctp/utils"
	"github.com/circlefin/noble-cctp/utils/mocks"
	"github.com/circlefin/noble-cctp/x/cctp/keeper"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Authority not set
 * Invalid authority
 */

func TestUpdateTokenControllerHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	newTokenController := utils.AccAddress()

	message := types.MsgUpdateTokenController{
		From:               owner,
		NewTokenController: newTokenController,
	}

	_, err := server.UpdateTokenController(ctx, &message)
	require.Nil(t, err)

	actual := testkeeper.GetTokenController(ctx)
	require.Equal(t, newTokenController, actual)
}

func TestUpdateTokenControllerAuthorityIsNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgUpdateTokenController{
		From:               "not the authority",
		NewTokenController: utils.AccAddress(),
	}
	require.Panicsf(t, func() {
		_, _ = server.UpdateTokenController(ctx, &message)
	}, "cctp owner not found in state")
}

func TestUpdateTokenControllerInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	newTokenController := utils.AccAddress()

	message := types.MsgUpdateTokenController{
		From:               utils.AccAddress(),
		NewTokenController: newTokenController,
	}

	_, err := server.UpdateTokenController(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the authority")
}
