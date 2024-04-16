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
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestSendingAndReceivingMessagesPausedQuery(t *testing.T) {
	SendingAndReceivingMessagesPaused := types.SendingAndReceivingMessagesPaused{Paused: true}

	for _, tc := range []struct {
		desc     string
		set      bool
		request  *types.QueryGetSendingAndReceivingMessagesPausedRequest
		response *types.QueryGetSendingAndReceivingMessagesPausedResponse
		err      error
	}{
		{
			desc:     "HappyPath",
			set:      true,
			request:  &types.QueryGetSendingAndReceivingMessagesPausedRequest{},
			response: &types.QueryGetSendingAndReceivingMessagesPausedResponse{Paused: SendingAndReceivingMessagesPaused},
		},
		{
			desc:    "NotFound",
			set:     false,
			request: &types.QueryGetSendingAndReceivingMessagesPausedRequest{},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			keeper, ctx := mocks.CctpKeeper()

			if tc.set {
				keeper.SetSendingAndReceivingMessagesPaused(ctx, SendingAndReceivingMessagesPaused)
			}

			response, err := keeper.SendingAndReceivingMessagesPaused(ctx, tc.request)

			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					utils.Fill(tc.response),
					utils.Fill(response),
				)
			}
		})
	}
}
