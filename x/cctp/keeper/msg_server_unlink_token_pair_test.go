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
 * Token pair not found
 */
func TestUnlinkTokenPairHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	tokenPair := types.TokenPair{
		RemoteDomain: 0,
		RemoteToken:  token,
		LocalToken:   "uusdc",
	}
	testkeeper.SetTokenPair(ctx, tokenPair)

	message := types.MsgUnlinkTokenPair{
		From:         tokenController,
		RemoteDomain: tokenPair.RemoteDomain,
		RemoteToken:  tokenPair.RemoteToken,
		LocalToken:   tokenPair.LocalToken,
	}

	_, err := server.UnlinkTokenPair(ctx, &message)
	require.Nil(t, err)

	_, found := testkeeper.GetTokenPair(ctx, message.RemoteDomain, message.RemoteToken)
	require.False(t, found)
}

func TestUnlinkTokenPairAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgUnlinkTokenPair{
		From:         utils.AccAddress(),
		RemoteDomain: 0,
		RemoteToken:  token,
		LocalToken:   "uusdc",
	}

	require.PanicsWithValue(t, "cctp token controller not found in state", func() {
		_, _ = server.UnlinkTokenPair(ctx, &message)
	})
}

func TestUnlinkTokenPairInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgUnlinkTokenPair{
		From:         "not the authority",
		RemoteDomain: 0,
		RemoteToken:  token,
		LocalToken:   "uusdc",
	}

	_, err := server.UnlinkTokenPair(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot unlink token pairs")
}

func TestUnlinkTokenPairTokenPairNotFound(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := utils.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgUnlinkTokenPair{
		From:         tokenController,
		RemoteDomain: 0,
		RemoteToken:  token,
		LocalToken:   "uusdc",
	}

	_, err := server.UnlinkTokenPair(ctx, &message)
	require.ErrorIs(t, types.ErrTokenPairNotFound, err)
	require.Contains(t, err.Error(), "token pair doesn't exist in store")
}
