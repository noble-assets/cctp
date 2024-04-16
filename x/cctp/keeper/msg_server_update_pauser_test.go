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

func TestUpdatePauserHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pauser := utils.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	newPauser := utils.AccAddress()

	message := types.MsgUpdatePauser{
		From:      owner,
		NewPauser: newPauser,
	}

	_, err := server.UpdatePauser(ctx, &message)
	require.Nil(t, err)

	actual := testkeeper.GetPauser(ctx)
	require.Equal(t, newPauser, actual)
}

func TestUpdatePauserAuthorityIsNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	pauser := utils.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	message := types.MsgUpdatePauser{
		From:      "not the authority",
		NewPauser: utils.AccAddress(),
	}
	require.Panicsf(t, func() {
		_, _ = server.UpdatePauser(ctx, &message)
	}, "cctp owner not found in state")
}

func TestUpdatePauserInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)
	pauser := utils.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	newPauser := utils.AccAddress()

	message := types.MsgUpdatePauser{
		From:      utils.AccAddress(),
		NewPauser: newPauser,
	}

	_, err := server.UpdatePauser(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot update the pauser")
}
