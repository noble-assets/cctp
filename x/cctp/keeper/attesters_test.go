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

func createNAttesters(keeper *keeper.Keeper, ctx context.Context, n int) []types.Attester {
	items := make([]types.Attester, n)
	for i := range items {
		items[i].Attester = "Attester" + strconv.Itoa(i)
		keeper.SetAttester(ctx, items[i])
	}
	return items
}

func TestAttestersGet(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNAttesters(cctpKeeper, ctx, 10)
	for _, item := range items {
		attester, found := cctpKeeper.GetAttester(ctx,
			item.Attester,
		)
		require.True(t, found)
		require.Equal(t,
			utils.Fill(&item),
			utils.Fill(&attester),
		)
	}
}

func TestAttestersRemove(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNAttesters(cctpKeeper, ctx, 10)
	for _, item := range items {
		cctpKeeper.DeleteAttester(ctx, item.Attester)
		_, found := cctpKeeper.GetAttester(ctx, item.Attester)
		require.False(t, found)
	}
}

func TestAttestersGetAll(t *testing.T) {
	cctpKeeper, ctx := mocks.CctpKeeper()
	items := createNAttesters(cctpKeeper, ctx, 10)
	denom := make([]types.Attester, len(items))
	copy(denom, items)
	require.ElementsMatch(t,
		utils.Fill(denom),
		utils.Fill(cctpKeeper.GetAllAttesters(ctx)),
	)
}
