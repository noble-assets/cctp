/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
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
	"testing"

	"cosmossdk.io/errors"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/keeper"
	"github.com/circlefin/noble-fiattokenfactory/x/fiattokenfactory/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func FiatTokenfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		MockBankKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}

type MockFiatTokenfactoryKeeper struct{}

func (MockFiatTokenfactoryKeeper) Mint(ctx sdk.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	return &types.MsgMintResponse{}, nil
}

func (MockFiatTokenfactoryKeeper) Burn(ctx sdk.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	return &types.MsgBurnResponse{}, nil
}

func (MockFiatTokenfactoryKeeper) GetMintingDenom(ctx sdk.Context) (val types.MintingDenom) {
	return types.MintingDenom{Denom: "uusdc"}
}

func ErrFiatTokenfactoryKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		MockBankKeeper{},
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

// MockErrFiatTokenfactoryKeeper - all dependencies err
type MockErrFiatTokenfactoryKeeper struct{}

func (k MockErrFiatTokenfactoryKeeper) Mint(ctx sdk.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	return nil, errors.Wrap(types.ErrBurn, "error calling mint")
}

func (k MockErrFiatTokenfactoryKeeper) GetMintingDenom(ctx sdk.Context) (val types.MintingDenom) {
	return types.MintingDenom{Denom: "uusdc"}
}

func (MockErrFiatTokenfactoryKeeper) Burn(ctx sdk.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	return &types.MsgBurnResponse{}, types.ErrBurn
}
