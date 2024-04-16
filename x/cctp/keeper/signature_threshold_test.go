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

func TestSignatureThreshold(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()
	_, found := keeper.GetSignatureThreshold(ctx)
	require.False(t, found)

	SignatureThreshold := types.SignatureThreshold{Amount: 2}
	keeper.SetSignatureThreshold(ctx, SignatureThreshold)

	threshold, found := keeper.GetSignatureThreshold(ctx)
	require.True(t, found)
	require.Equal(t,
		SignatureThreshold,
		utils.Fill(&threshold),
	)

	newSignatureThreshold := types.SignatureThreshold{Amount: 3}

	keeper.SetSignatureThreshold(ctx, newSignatureThreshold)

	threshold, found = keeper.GetSignatureThreshold(ctx)
	require.True(t, found)
	require.Equal(t,
		newSignatureThreshold,
		utils.Fill(&threshold),
	)
}
