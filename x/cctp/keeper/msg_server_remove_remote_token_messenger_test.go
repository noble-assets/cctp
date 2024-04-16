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
* Remote token messenger not found
 */

func TestRemoveRemoteTokenMessengerHappyPath(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	addMessage := types.MsgAddRemoteTokenMessenger{
		From:     owner,
		DomainId: 0,
		Address:  tokenMessenger,
	}

	_, err := server.AddRemoteTokenMessenger(ctx, &addMessage)
	require.Nil(t, err)

	removeMessage := types.MsgRemoveRemoteTokenMessenger{
		From:     owner,
		DomainId: addMessage.DomainId,
	}

	_, err = server.RemoveRemoteTokenMessenger(ctx, &removeMessage)
	require.Nil(t, err)

	_, found := testkeeper.GetRemoteTokenMessenger(ctx, removeMessage.DomainId)
	require.False(t, found)
}

func TestRemoveRemoteTokenMessengerAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgRemoveRemoteTokenMessenger{
		From:     utils.AccAddress(),
		DomainId: 0,
	}

	require.PanicsWithValue(t, "cctp owner not found in state", func() {
		_, _ = server.RemoveRemoteTokenMessenger(ctx, &message)
	})
}

func TestRemoveRemoteTokenMessengerInvalidAuthority(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgRemoveRemoteTokenMessenger{
		From:     "not the authority address",
		DomainId: 0,
	}

	_, err := server.RemoveRemoteTokenMessenger(ctx, &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot remove remote token messengers")
}

func TestRemoveRemoteTokenMessengerTokenMessengerNotFound(t *testing.T) {
	testkeeper, ctx := mocks.CctpKeeper()
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := utils.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgRemoveRemoteTokenMessenger{
		From:     owner,
		DomainId: 0,
	}

	_, err := server.RemoveRemoteTokenMessenger(ctx, &message)
	require.ErrorIs(t, types.ErrRemoteTokenMessengerNotFound, err)
	require.Contains(t, err.Error(), "no remote token messenger was found for this domain")
}
