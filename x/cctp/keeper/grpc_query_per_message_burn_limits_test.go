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
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/circlefin/noble-cctp/utils"
	"github.com/circlefin/noble-cctp/utils/mocks"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestPerMessageBurnLimitQuery(t *testing.T) {
	PerMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "uusdc",
		Amount: math.NewInt(int64(42)),
	}

	for _, tc := range []struct {
		desc     string
		set      bool
		request  *types.QueryGetPerMessageBurnLimitRequest
		response *types.QueryGetPerMessageBurnLimitResponse
		err      error
	}{
		{
			desc:     "HappyPath",
			set:      true,
			request:  &types.QueryGetPerMessageBurnLimitRequest{Denom: PerMessageBurnLimit.Denom},
			response: &types.QueryGetPerMessageBurnLimitResponse{BurnLimit: PerMessageBurnLimit},
		},
		{
			desc:    "NotFound",
			set:     false,
			request: &types.QueryGetPerMessageBurnLimitRequest{},
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
				keeper.SetPerMessageBurnLimit(ctx, PerMessageBurnLimit)
			}

			response, err := keeper.PerMessageBurnLimit(ctx, tc.request)

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

func TestPerMessageBurnLimitQueryPaginated(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()
	msgs := createNPerMessageBurnLimits(keeper, ctx, 5)
	perMessageBurnLimits := make([]types.PerMessageBurnLimit, len(msgs))
	copy(perMessageBurnLimits, msgs)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPerMessageBurnLimitsRequest {
		return &types.QueryAllPerMessageBurnLimitsRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(perMessageBurnLimits); i += step {
			resp, err := keeper.PerMessageBurnLimits(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BurnLimits), step)
			require.Subset(t,
				utils.Fill(perMessageBurnLimits),
				utils.Fill(resp.BurnLimits),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(perMessageBurnLimits); i += step {
			resp, err := keeper.PerMessageBurnLimits(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BurnLimits), step)
			require.Subset(t,
				utils.Fill(perMessageBurnLimits),
				utils.Fill(resp.BurnLimits),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.PerMessageBurnLimits(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(perMessageBurnLimits), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			utils.Fill(perMessageBurnLimits),
			utils.Fill(resp.BurnLimits),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.PerMessageBurnLimits(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
	t.Run("PaginateError", func(t *testing.T) {
		_, err := keeper.PerMessageBurnLimits(ctx, request([]byte("key"), 1, 0, true))
		require.Contains(t, err.Error(), "invalid request, either offset or key is expected, got both")
	})
}

func TestPerMessageBurnLimitQueryPaginatedInvalidState(t *testing.T) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	keeper, ctx := mocks.CctpKeeperWithKey(storeKey)

	adapter := runtime.KVStoreAdapter(runtime.NewKVStoreService(storeKey).OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.PerMessageBurnLimitKeyPrefix))
	store.Set(types.KeyPrefix(string(types.PerMessageBurnLimitKey("denom"))), []byte("invalid"))

	_, err := keeper.PerMessageBurnLimits(ctx, &types.QueryAllPerMessageBurnLimitsRequest{})

	parsedErr, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, parsedErr.Code(), codes.Internal)
}
