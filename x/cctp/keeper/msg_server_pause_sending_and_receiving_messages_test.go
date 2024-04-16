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
 */
func TestPauseSendingAndReceivingMessagesHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	pauser := utils.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	message := types.MsgPauseSendingAndReceivingMessages{
		From: pauser,
	}
	_, err := server.PauseSendingAndReceivingMessages(ctx, &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetSendingAndReceivingMessagesPaused(ctx)
	require.True(t, found)
	require.Equal(t, true, actual.Paused)
}

func TestPauseSendingAndReceivingMessagesAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgPauseSendingAndReceivingMessages{
		From: "authority",
	}

	require.PanicsWithValue(t, "cctp pauser not found in state", func() {
		_, _ = server.PauseSendingAndReceivingMessages(ctx, &message)
	})
}

func TestPauseSendingAndReceivingMessagesInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	pauser := utils.AccAddress()
	testkeeper.SetPauser(ctx, pauser)

	message := types.MsgPauseSendingAndReceivingMessages{
		From: "not the authority",
	}
	_, err := server.PauseSendingAndReceivingMessages(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot pause sending and receiving")
}
