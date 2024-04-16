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

	"github.com/circlefin/noble-cctp/utils/mocks"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/stretchr/testify/require"
)

func TestBurningAndMintingPaused(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	paused := types.BurningAndMintingPaused{Paused: true}
	keeper.SetBurningAndMintingPaused(ctx, paused)

	isPaused, found := keeper.GetBurningAndMintingPaused(ctx)
	require.True(t, found)
	require.True(t, isPaused.Paused)

	newPaused := types.BurningAndMintingPaused{Paused: false}

	keeper.SetBurningAndMintingPaused(ctx, newPaused)

	isPaused, found = keeper.GetBurningAndMintingPaused(ctx)
	require.True(t, found)
	require.False(t, isPaused.Paused)
}
