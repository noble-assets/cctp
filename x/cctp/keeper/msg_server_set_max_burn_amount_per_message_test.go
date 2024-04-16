/*
 * Copyright (c) 2024, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
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

func TestSetMaxBurnAmountPerMessageHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       tokenController,
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	_, err := server.SetMaxBurnAmountPerMessage(ctx, &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetPerMessageBurnLimit(ctx, message.LocalToken)
	require.True(t, found)
	require.Equal(t, message.LocalToken, actual.Denom)
	require.Equal(t, message.Amount, actual.Amount)
}

func TestSetMaxBurnAmountPerMessageAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       utils.AccAddress(),
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	require.PanicsWithValue(t, "cctp token controller not found in state", func() {
		_, _ = server.SetMaxBurnAmountPerMessage(ctx, &message)
	})
}

func TestSetMaxBurnAmountPerMessageInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgSetMaxBurnAmountPerMessage{
		From:       "not authority",
		LocalToken: "uusdc",
		Amount:     math.NewInt(123),
	}

	_, err := server.SetMaxBurnAmountPerMessage(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot set the max burn amount per message")
}
