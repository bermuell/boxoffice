package types

import (
	"testing"

	"github.com/bermuell/boxoffice/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateConcert_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateConcert
		err  error
	}{
		{
			name: "invalid creater address",
			msg: MsgCreateConcert{
				Creator: "invalid_address",
				Name:    "My concert",
				Volume:  "10",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid creator address",
			msg: MsgCreateConcert{
				Creator: sample.AccAddress(),
				Name:    "My concert",
				Volume:  "10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
