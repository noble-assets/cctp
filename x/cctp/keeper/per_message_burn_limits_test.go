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
	"context"
	"strconv"
	"testing"

	"cosmossdk.io/math"
	"github.com/circlefin/noble-cctp/utils"
	"github.com/circlefin/noble-cctp/utils/mocks"
	"github.com/circlefin/noble-cctp/x/cctp/keeper"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func createNPerMessageBurnLimits(keeper *keeper.Keeper, ctx context.Context, n int) []types.PerMessageBurnLimit {
	items := make([]types.PerMessageBurnLimit, n)
	for i := range items {
		items[i].Denom = "amount" + strconv.Itoa(i)
		items[i].Amount = math.NewInt(int64(i))
		keeper.SetPerMessageBurnLimit(ctx, items[i])
	}
	return items
}

func TestPerMessageBurnLimit(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	_, found := keeper.GetPerMessageBurnLimit(ctx, "usdc")
	require.False(t, found)

	perMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "usdc",
		Amount: math.NewInt(123),
	}
	keeper.SetPerMessageBurnLimit(ctx, perMessageBurnLimit)

	limit, found := keeper.GetPerMessageBurnLimit(ctx, perMessageBurnLimit.Denom)
	require.True(t, found)
	require.Equal(t,
		perMessageBurnLimit,
		utils.Fill(&limit),
	)

	newPerMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "usdc",
		Amount: math.NewInt(456),
	}

	keeper.SetPerMessageBurnLimit(ctx, newPerMessageBurnLimit)

	limit, found = keeper.GetPerMessageBurnLimit(ctx, newPerMessageBurnLimit.Denom)
	require.True(t, found)
	require.Equal(t,
		newPerMessageBurnLimit,
		utils.Fill(&limit),
	)

	separateCurrencyPerMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  "euroc",
		Amount: math.NewInt(789),
	}
	keeper.SetPerMessageBurnLimit(ctx, separateCurrencyPerMessageBurnLimit)
	limit, found = keeper.GetPerMessageBurnLimit(ctx, separateCurrencyPerMessageBurnLimit.Denom)

	require.True(t, found)
	require.Equal(t,
		separateCurrencyPerMessageBurnLimit,
		utils.Fill(&limit),
	)
}

func TestPerMessageBurnLimitsGetAll(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNPerMessageBurnLimits(cctpKeeper, ctx, 10)
	denom := make([]types.PerMessageBurnLimit, len(items))
	copy(denom, items)

	require.ElementsMatch(t,
		utils.Fill(denom),
		utils.Fill(cctpKeeper.GetAllPerMessageBurnLimits(ctx)),
	)
}
