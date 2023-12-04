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
package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveRemoteTokenMessenger = "remove_remote_token_messenger"

var _ sdk.Msg = &MsgRemoveRemoteTokenMessenger{}

func NewMsgRemoveRemoteTokenMessenger(from string, domainId uint32) *MsgRemoveRemoteTokenMessenger {
	return &MsgRemoveRemoteTokenMessenger{
		From:     from,
		DomainId: domainId,
	}
}

func (msg *MsgRemoveRemoteTokenMessenger) Route() string {
	return RouterKey
}

func (msg *MsgRemoveRemoteTokenMessenger) Type() string {
	return TypeMsgRemoveRemoteTokenMessenger
}

func (msg *MsgRemoveRemoteTokenMessenger) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgRemoveRemoteTokenMessenger) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveRemoteTokenMessenger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}
