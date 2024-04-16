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
 * Owner not set
 * Pending owner not set
 * Invalid Pending owner
 */

func TestAcceptOwnerHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pendingOwner := utils.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: pendingOwner,
	}

	_, err := server.AcceptOwner(ctx, &message)
	require.Nil(t, err)

	newOwner := testkeeper.GetOwner(ctx)
	require.Equal(t, pendingOwner, newOwner)
}

func TestAcceptOwnerOwnerNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	pendingOwner := utils.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: pendingOwner,
	}

	require.Panicsf(t, func() {
		_, _ = server.AcceptOwner(ctx, &message)
	}, "cctp owner not found in state")
}

func TestAcceptOwnerPendingOwnerNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAcceptOwner{
		From: utils.AccAddress(),
	}

	_, err := server.AcceptOwner(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "pending owner is not set")
}

func TestAcceptOwnerInvalidPendingOwner(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pendingOwner := utils.AccAddress()
	testkeeper.SetPendingOwner(ctx, pendingOwner)

	message := types.MsgAcceptOwner{
		From: utils.AccAddress(),
	}

	_, err := server.AcceptOwner(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "you are not the pending owner")
}
