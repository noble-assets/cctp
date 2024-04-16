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
	"github.com/stretchr/testify/require"
)

func TestOwner(t *testing.T) {
	keeper, ctx := mocks.CctpKeeper()

	owner := utils.AccAddress()
	keeper.SetOwner(ctx, owner)

	require.Equal(t, owner, keeper.GetOwner(ctx))

	newOwner := utils.AccAddress()
	keeper.SetOwner(ctx, newOwner)

	require.Equal(t, newOwner, keeper.GetOwner(ctx))
}
