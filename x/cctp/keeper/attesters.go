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

// GetAttester returns an attester
func (k Keeper) GetAttester(ctx context.Context, key string) (val types.Attester, found bool) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.AttesterKeyPrefix))

	b := store.Get(types.KeyPrefix(string(types.AttesterKey([]byte(key)))))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetAttester sets an attester in the store
func (k Keeper) SetAttester(ctx context.Context, key types.Attester) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.AttesterKeyPrefix))
	b := k.cdc.MustMarshal(&key)
	store.Set(types.KeyPrefix(string(types.AttesterKey([]byte(key.Attester)))), b)
}

// DeleteAttester removes an attester
func (k Keeper) DeleteAttester(ctx context.Context, key string) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.AttesterKeyPrefix))
	store.Delete(types.AttesterKey([]byte(key)))
}

// GetAllAttesters returns all attesters
func (k Keeper) GetAllAttesters(ctx context.Context) (list []types.Attester) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(adapter, types.KeyPrefix(types.AttesterKeyPrefix))
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Attester
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
