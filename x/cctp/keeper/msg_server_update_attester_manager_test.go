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

func TestUpdateAttesterManagerHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	attesterManager := utils.AccAddress()
	testkeeper.SetAttesterManager(ctx, attesterManager)

	newAttesterManager := utils.AccAddress()

	message := types.MsgUpdateAttesterManager{
		From:               owner,
		NewAttesterManager: newAttesterManager,
	}

	_, err := server.UpdateAttesterManager(ctx, &message)
	require.Nil(t, err)

	actual := testkeeper.GetAttesterManager(ctx)
	require.Equal(t, newAttesterManager, actual)
}

func TestUpdateAttesterManagerAuthorityIsNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUpdateAttesterManager{
		From:               "not the authority",
		NewAttesterManager: utils.AccAddress(),
	}
	require.Panicsf(t, func() {
		_, _ = server.UpdateAttesterManager(ctx, &message)
	}, "cctp owner not found in state")
}

func TestUpdateAttesterManagerInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	newAttesterManager := utils.AccAddress()

	message := types.MsgUpdateAttesterManager{
		From:               utils.AccAddress(),
		NewAttesterManager: newAttesterManager,
	}

	_, err := server.UpdateAttesterManager(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the attester manager")
}
