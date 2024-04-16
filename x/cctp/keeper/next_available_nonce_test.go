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
)

func TestNextAvailableNonce(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	_, found := keeper.GetNextAvailableNonce(ctx)
	require.False(t, found)

	savedNonce := types.Nonce{Nonce: 21}
	keeper.SetNextAvailableNonce(ctx, savedNonce)

	next, found := keeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t,
		savedNonce,
		utils.Fill(&next),
	)

	newSavedNonce := types.Nonce{Nonce: 22}

	keeper.SetNextAvailableNonce(ctx, newSavedNonce)

	next, found = keeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t,
		newSavedNonce,
		utils.Fill(&next),
	)
}

func TestNextAvailableNonceReserveAndIncrement(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	savedNonce := types.Nonce{Nonce: 21}
	keeper.SetNextAvailableNonce(ctx, savedNonce)

	prev, found := keeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t,
		savedNonce,
		utils.Fill(&prev),
	)

	// method returns the nonce being reserved
	nextFromMethod := keeper.ReserveAndIncrementNonce(ctx)
	require.Equal(t,
		types.Nonce{
			Nonce: prev.Nonce,
		},
		utils.Fill(&nextFromMethod),
	)

	// retrieving the nonce should get reserved nonce + 1
	next, found := keeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t,
		types.Nonce{
			Nonce: prev.Nonce + 1,
		},
		utils.Fill(&next),
	)
}
