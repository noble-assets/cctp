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

	"github.com/circlefin/noble-cctp/utils"
	"github.com/circlefin/noble-cctp/utils/mocks"
	"github.com/circlefin/noble-cctp/x/cctp/keeper"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func createNTokenPairs(keeper *keeper.Keeper, ctx context.Context, n int) []types.TokenPair {
	items := make([]types.TokenPair, n)
	for i := range items {
		items[i].RemoteDomain = uint32(i)
		items[i].RemoteToken = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, byte(i)}
		items[i].LocalToken = "token" + strconv.Itoa(i)

		keeper.SetTokenPair(ctx, items[i])
	}
	return items
}

func TestTokenPairsGet(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	for _, item := range items {
		tokenPair, found := cctpKeeper.GetTokenPair(ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		require.True(t, found)
		require.Equal(t,
			utils.Fill(&item),
			utils.Fill(&tokenPair),
		)
	}
}

func TestTokenPairsRemove(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	for _, item := range items {
		cctpKeeper.DeleteTokenPair(
			ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		_, found := cctpKeeper.GetTokenPair(
			ctx,
			item.RemoteDomain,
			item.RemoteToken,
		)
		require.False(t, found)
	}
}

func TestTokenPairsGetAll(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNTokenPairs(cctpKeeper, ctx, 10)
	denom := make([]types.TokenPair, len(items))
	copy(denom, items)

	require.ElementsMatch(t,
		utils.Fill(denom),
		utils.Fill(cctpKeeper.GetAllTokenPairs(ctx)),
	)
}
