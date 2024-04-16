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

func TestMaxMessageBodySize(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	_, found := keeper.GetMaxMessageBodySize(ctx)
	require.False(t, found)

	MaxMessageBodySize := types.MaxMessageBodySize{Amount: 21}
	keeper.SetMaxMessageBodySize(ctx, MaxMessageBodySize)

	maxMessageBodySize, found := keeper.GetMaxMessageBodySize(ctx)
	require.True(t, found)
	require.Equal(t,
		MaxMessageBodySize,
		utils.Fill(&maxMessageBodySize),
	)

	newMaxMessageBodySize := types.MaxMessageBodySize{Amount: 22}

	keeper.SetMaxMessageBodySize(ctx, newMaxMessageBodySize)

	maxMessageBodySize, found = keeper.GetMaxMessageBodySize(ctx)
	require.True(t, found)
	require.Equal(t,
		newMaxMessageBodySize,
		utils.Fill(&maxMessageBodySize),
	)
}
