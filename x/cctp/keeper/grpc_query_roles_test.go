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

func TestRolesQuery(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	owner := "test-owner"
	keeper.SetOwner(ctx, owner)
	attesterManager := "test-attester-manager"
	keeper.SetAttesterManager(ctx, attesterManager)
	pauser := "test-pauser"
	keeper.SetPauser(ctx, pauser)
	tokenController := "test-token-controller"
	keeper.SetTokenController(ctx, tokenController)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryRolesRequest
		response *types.QueryRolesResponse
		err      error
	}{
		{
			desc:    "First",
			request: &types.QueryRolesRequest{},
			response: &types.QueryRolesResponse{
				Owner:           owner,
				AttesterManager: attesterManager,
				Pauser:          pauser,
				TokenController: tokenController,
			},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Roles(ctx, tc.request)
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
