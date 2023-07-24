package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBuyTicket = "buy_ticket"

var _ sdk.Msg = &MsgBuyTicket{}

func NewMsgBuyTicket(creator string, concertIndex string) *MsgBuyTicket {
	return &MsgBuyTicket{
		Creator:      creator,
		ConcertIndex: concertIndex,
	}
}

func (msg *MsgBuyTicket) Route() string {
	return RouterKey
}

func (msg *MsgBuyTicket) Type() string {
	return TypeMsgBuyTicket
}

func (msg *MsgBuyTicket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyTicket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
