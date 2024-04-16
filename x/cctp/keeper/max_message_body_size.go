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

package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetMaxMessageBodySize returns the MaxMessageBodySize
func (k Keeper) GetMaxMessageBodySize(ctx context.Context) (val types.MaxMessageBodySize, found bool) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.MaxMessageBodySizeKey))

	b := store.Get(types.KeyPrefix(types.MaxMessageBodySizeKey))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetMaxMessageBodySize sets MaxMessageBodySize in the store
func (k Keeper) SetMaxMessageBodySize(ctx context.Context, amount types.MaxMessageBodySize) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.MaxMessageBodySizeKey))
	b := k.cdc.MustMarshal(&amount)
	store.Set(types.KeyPrefix(types.MaxMessageBodySizeKey), b)
}
