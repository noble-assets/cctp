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

package mocks

import (
	"context"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/circlefin/noble-cctp/x/cctp/keeper"
	"github.com/circlefin/noble-cctp/x/cctp/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CctpKeeper() (*keeper.Keeper, context.Context) {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	return CctpKeeperWithKey(key)
}

func CctpKeeperWithKey(key *storetypes.KVStoreKey) (*keeper.Keeper, context.Context) {
	logger := log.NewNopLogger()

	state := store.NewCommitMultiStore(db.NewMemDB(), logger, metrics.NewNoOpMetrics())
	state.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
	_ = state.LoadLatestVersion()

	return keeper.NewKeeper(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		logger,
		runtime.NewKVStoreService(key),
		MockBankKeeper{},
		MockFiatTokenfactoryKeeper{},
	), sdk.NewContext(state, cmtproto.Header{}, false, logger)
}

func CctpKeeperWithErrBank() (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	key := storetypes.NewKVStoreKey(types.StoreKey)
	state := store.NewCommitMultiStore(db.NewMemDB(), logger, metrics.NewNoOpMetrics())
	state.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
	_ = state.LoadLatestVersion()

	return keeper.NewKeeper(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		logger,
		runtime.NewKVStoreService(key),
		MockErrBankKeeper{},
		MockFiatTokenfactoryKeeper{},
	), sdk.NewContext(state, cmtproto.Header{}, false, logger)
}

func CctpKeeperWithErrFTF() (*keeper.Keeper, sdk.Context) {
	logger := log.NewNopLogger()

	key := storetypes.NewKVStoreKey(types.StoreKey)
	state := store.NewCommitMultiStore(db.NewMemDB(), logger, metrics.NewNoOpMetrics())
	state.MountStoreWithDB(key, storetypes.StoreTypeIAVL, nil)
	_ = state.LoadLatestVersion()

	return keeper.NewKeeper(
		codec.NewProtoCodec(codectypes.NewInterfaceRegistry()),
		logger,
		runtime.NewKVStoreService(key),
		MockBankKeeper{},
		MockErrFiatTokenfactoryKeeper{},
	), sdk.NewContext(state, cmtproto.Header{}, false, logger)
}
