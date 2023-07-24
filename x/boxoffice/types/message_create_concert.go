package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateConcert = "create_concert"

var _ sdk.Msg = &MsgCreateConcert{}

func NewMsgCreateConcert(creator string, name string, volume string) *MsgCreateConcert {
	return &MsgCreateConcert{
		Creator: creator,
		Name:    name,
		Volume:  volume,
	}
}

func (msg *MsgCreateConcert) Route() string {
	return RouterKey
}

func (msg *MsgCreateConcert) Type() string {
	return TypeMsgCreateConcert
}

func (msg *MsgCreateConcert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateConcert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateConcert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	volume, err := strconv.ParseUint(msg.Volume, 10, 64)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidVolume, "Error parsing concert volume %s", err)
	}
	if volume <= 0 {
		return sdkerrors.Wrapf(ErrInvalidVolume, "Concert volume needs to be > 0")
	}
	//TODO: Add further stateless checks if any
	return nil
}
